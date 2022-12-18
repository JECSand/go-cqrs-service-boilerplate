package server

import (
	"context"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/authentication"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/constants"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/interceptors"
	kafkaClient "github.com/JECSand/go-cqrs-service-boilerplate/pkg/kafka"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/mongodb"
	redisClient "github.com/JECSand/go-cqrs-service-boilerplate/pkg/redis"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/tracing"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/config"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/cache"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/data"
	queryGrpc "github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/delivery/grpc"
	queryKafka "github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/delivery/kafka"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/metrics"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/services"
	queryService "github.com/JECSand/go-cqrs-service-boilerplate/query_service/protos/user_query"
	"github.com/go-playground/validator"
	"github.com/go-redis/redis/v8"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/heptiolabs/healthcheck"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	maxConnectionIdle = 5
	gRPCTimeout       = 15
	maxConnectionAge  = 5
	gRPCTime          = 10
	stackSize         = 1 << 10 // 1 KB
)

type server struct {
	log         logging.Logger
	auth        authentication.Authenticator
	cfg         *config.Config
	v           *validator.Validate
	kafkaConn   *kafka.Conn
	im          interceptors.InterceptorManager
	mongoClient *mongo.Client
	redisClient redis.UniversalClient
	ps          *services.UserService
	metrics     *metrics.QueryServiceMetrics
}

func NewServer(log logging.Logger, auth authentication.Authenticator, cfg *config.Config) *server {
	return &server{
		log:  log,
		auth: auth,
		cfg:  cfg,
		v:    validator.New(),
	}
}

func (s *server) newReaderGrpcServer() (func() error, *grpc.Server, error) {
	l, err := net.Listen("tcp", s.cfg.GRPC.Port)
	if err != nil {
		return nil, nil, errors.Wrap(err, "net.Listen")
	}
	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: maxConnectionIdle * time.Minute,
			Timeout:           gRPCTimeout * time.Second,
			MaxConnectionAge:  maxConnectionAge * time.Minute,
			Time:              gRPCTime * time.Minute,
		}),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_recovery.UnaryServerInterceptor(),
			s.im.Logger,
		),
		),
	)
	queryGrpcService := queryGrpc.NewQueryGrpcService(s.log, s.cfg, s.v, s.ps, s.metrics)
	queryService.RegisterQueryServiceServer(grpcServer, queryGrpcService)
	grpc_prometheus.Register(grpcServer)
	if s.cfg.GRPC.Development {
		reflection.Register(grpcServer)
	}
	go func() {
		s.log.Infof("Query gRPC server is listening on port: %s", s.cfg.GRPC.Port)
		s.log.Fatal(grpcServer.Serve(l))
	}()
	return l.Close, grpcServer, nil
}

func (s *server) connectKafkaBrokers(ctx context.Context) error {
	kafkaConn, err := kafkaClient.NewKafkaConn(ctx, s.cfg.Kafka)
	if err != nil {
		return errors.Wrap(err, "kafka.NewKafkaCon")
	}

	s.kafkaConn = kafkaConn

	brokers, err := kafkaConn.Brokers()
	if err != nil {
		return errors.Wrap(err, "kafkaConn.Brokers")
	}

	s.log.Infof("kafka connected to brokers: %+v", brokers)

	return nil
}

func (s *server) getConsumerGroupTopics() []string {
	return []string{
		s.cfg.KafkaTopics.UserCreated.TopicName,
		s.cfg.KafkaTopics.UserUpdated.TopicName,
		s.cfg.KafkaTopics.UserDeleted.TopicName,
	}
}

func (s *server) runHealthCheck(ctx context.Context) {
	health := healthcheck.NewHandler()
	health.AddLivenessCheck(s.cfg.ServiceName, healthcheck.AsyncWithContext(ctx, func() error {
		return nil
	}, time.Duration(s.cfg.Probes.CheckIntervalSeconds)*time.Second))
	health.AddReadinessCheck(constants.Redis, healthcheck.AsyncWithContext(ctx, func() error {
		return s.redisClient.Ping(ctx).Err()
	}, time.Duration(s.cfg.Probes.CheckIntervalSeconds)*time.Second))
	health.AddReadinessCheck(constants.MongoDB, healthcheck.AsyncWithContext(ctx, func() error {
		return s.mongoClient.Ping(ctx, nil)
	}, time.Duration(s.cfg.Probes.CheckIntervalSeconds)*time.Second))
	health.AddReadinessCheck(constants.Kafka, healthcheck.AsyncWithContext(ctx, func() error {
		_, err := s.kafkaConn.Brokers()
		if err != nil {
			return err
		}
		return nil
	}, time.Duration(s.cfg.Probes.CheckIntervalSeconds)*time.Second))
	go func() {
		s.log.Infof("Query microservice Kubernetes probes listening on port: %s", s.cfg.Probes.Port)
		if err := http.ListenAndServe(s.cfg.Probes.Port, health); err != nil {
			s.log.WarnMsg("ListenAndServe", err)
		}
	}()
}

func (s *server) runMetrics(cancel context.CancelFunc) {
	metricsServer := echo.New()
	go func() {
		metricsServer.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
			StackSize:         stackSize,
			DisablePrintStack: true,
			DisableStackAll:   true,
		}))
		metricsServer.GET(s.cfg.Probes.PrometheusPath, echo.WrapHandler(promhttp.Handler()))
		s.log.Infof("Metrics server is running on port: %s", s.cfg.Probes.PrometheusPort)
		if err := metricsServer.Start(s.cfg.Probes.PrometheusPort); err != nil {
			s.log.Errorf("metricsServer.Start: %v", err)
			cancel()
		}
	}()
}

func (s *server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()
	s.im = interceptors.NewInterceptorManager(s.log, s.auth)
	s.metrics = metrics.NewQueryServiceMetrics(s.cfg)
	mongoDBConn, err := mongodb.NewMongoDBConn(ctx, s.cfg.Mongo)
	if err != nil {
		return errors.Wrap(err, "NewMongoDBConn")
	}
	s.mongoClient = mongoDBConn
	defer mongoDBConn.Disconnect(ctx) // nolint: errcheck
	s.log.Infof("Mongo connected: %v", mongoDBConn.NumberSessionsInProgress())
	s.redisClient = redisClient.NewRedisClient(s.cfg.Redis)
	defer s.redisClient.Close() // nolint: errcheck
	s.log.Infof("Redis connected: %+v", s.redisClient.PoolStats())
	mongoRepo := data.NewMongoRepository(s.log, s.cfg, s.mongoClient)
	redisRepo := cache.NewRedisCache(s.log, s.cfg, s.redisClient)
	s.ps = services.NewUserService(s.log, s.cfg, mongoRepo, redisRepo)
	readerMessageProcessor := queryKafka.NewQueryMessageProcessor(s.log, s.cfg, s.v, s.ps, s.metrics)
	s.log.Info("Starting Reader Kafka consumers")
	cg := kafkaClient.NewConsumerGroup(s.cfg.Kafka.Brokers, s.cfg.Kafka.GroupID, s.log)
	go cg.ConsumeTopic(ctx, s.getConsumerGroupTopics(), queryKafka.PoolSize, readerMessageProcessor.ProcessMessages)
	if err = s.connectKafkaBrokers(ctx); err != nil {
		return errors.Wrap(err, "s.connectKafkaBrokers")
	}
	defer s.kafkaConn.Close() // nolint: errCheck
	s.runHealthCheck(ctx)
	s.runMetrics(cancel)
	if s.cfg.Jaeger.Enable {
		tracer, closer, err := tracing.NewJaegerTracer(s.cfg.Jaeger)
		if err != nil {
			return err
		}
		defer closer.Close() // nolint: errCheck
		opentracing.SetGlobalTracer(tracer)
	}
	closeGrpcServer, grpcServer, err := s.newReaderGrpcServer()
	if err != nil {
		return errors.Wrap(err, "NewScmGrpcServer")
	}
	defer closeGrpcServer() // nolint: errCheck
	<-ctx.Done()
	grpcServer.GracefulStop()
	return nil
}
