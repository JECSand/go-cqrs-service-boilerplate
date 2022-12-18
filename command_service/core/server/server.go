package server

import (
	"context"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/config"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/core/controllers"
	grpc2 "github.com/JECSand/go-cqrs-service-boilerplate/command_service/core/delivery/grpc"
	kafkaConsumer "github.com/JECSand/go-cqrs-service-boilerplate/command_service/core/delivery/kafka"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/core/metrics"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/core/repositories"
	commandService "github.com/JECSand/go-cqrs-service-boilerplate/command_service/protos/user_command"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/authentication"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/constants"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/interceptors"
	kafkaClient "github.com/JECSand/go-cqrs-service-boilerplate/pkg/kafka"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/postgres"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/tracing"
	"github.com/go-playground/validator"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/heptiolabs/healthcheck"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
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
	log       logging.Logger
	auth      authentication.Authenticator
	cfg       *config.Config
	v         *validator.Validate
	kafkaConn *kafka.Conn
	ps        *controllers.UserService
	im        interceptors.InterceptorManager
	pgConn    *pgxpool.Pool
	metrics   *metrics.CommandServiceMetrics
}

func NewServer(log logging.Logger, auth authentication.Authenticator, cfg *config.Config) *server {
	return &server{
		log:  log,
		auth: auth,
		cfg:  cfg,
		v:    validator.New(),
	}
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

func (s *server) initKafkaTopics(ctx context.Context) {
	controller, err := s.kafkaConn.Controller()
	if err != nil {
		s.log.WarnMsg("kafkaConn.Controller", err)
		return
	}
	controllerURI := net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port))
	s.log.Infof("kafka controller uri: %s", controllerURI)
	conn, err := kafka.DialContext(ctx, "tcp", controllerURI)
	if err != nil {
		s.log.WarnMsg("initKafkaTopics.DialContext", err)
		return
	}
	defer conn.Close() // nolint: errCheck
	s.log.Infof("established new kafka controller connection: %s", controllerURI)
	userCreateTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.UserCreate.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.UserCreate.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.UserCreate.ReplicationFactor,
	}
	userCreatedTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.UserCreated.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.UserCreated.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.UserCreated.ReplicationFactor,
	}
	userUpdateTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.UserUpdate.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.UserUpdate.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.UserUpdate.ReplicationFactor,
	}
	userUpdatedTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.UserUpdated.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.UserUpdated.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.UserUpdated.ReplicationFactor,
	}
	userDeleteTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.UserDelete.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.UserDelete.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.UserDelete.ReplicationFactor,
	}
	userDeletedTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.UserDeleted.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.UserDeleted.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.UserDeleted.ReplicationFactor,
	}
	if err = conn.CreateTopics(
		userCreateTopic,
		userUpdateTopic,
		userCreatedTopic,
		userUpdatedTopic,
		userDeleteTopic,
		userDeletedTopic,
	); err != nil {
		s.log.WarnMsg("kafkaConn.CreateTopics", err)
		return
	}
	s.log.Infof("kafka topics created or already exists: %+v", []kafka.TopicConfig{userCreateTopic, userUpdateTopic, userCreatedTopic, userUpdatedTopic, userDeleteTopic, userDeletedTopic})
}

func (s *server) getConsumerGroupTopics() []string {
	return []string{
		s.cfg.KafkaTopics.UserCreate.TopicName,
		s.cfg.KafkaTopics.UserUpdate.TopicName,
		s.cfg.KafkaTopics.UserDelete.TopicName,
	}
}

func (s *server) runHealthCheck(ctx context.Context) {
	health := healthcheck.NewHandler()
	health.AddLivenessCheck(s.cfg.ServiceName, healthcheck.AsyncWithContext(ctx, func() error {
		return nil
	}, time.Duration(s.cfg.Probes.CheckIntervalSeconds)*time.Second))
	health.AddReadinessCheck(constants.Postgres, healthcheck.AsyncWithContext(ctx, func() error {
		return s.pgConn.Ping(ctx)
	}, time.Duration(s.cfg.Probes.CheckIntervalSeconds)*time.Second))
	health.AddReadinessCheck(constants.Kafka, healthcheck.AsyncWithContext(ctx, func() error {
		_, err := s.kafkaConn.Brokers()
		if err != nil {
			return err
		}
		return nil
	}, time.Duration(s.cfg.Probes.CheckIntervalSeconds)*time.Second))
	go func() {
		s.log.Infof("Command service Kubernetes probes listening on port: %s", s.cfg.Probes.Port)
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

func (s *server) newCommandGrpcServer() (func() error, *grpc.Server, error) {
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
	commandGrpcWriter := grpc2.NewCommandGrpcService(s.log, s.cfg, s.v, s.ps, s.metrics)
	commandService.RegisterCommandServiceServer(grpcServer, commandGrpcWriter)
	grpc_prometheus.Register(grpcServer)
	if s.cfg.GRPC.Development {
		reflection.Register(grpcServer)
	}
	go func() {
		s.log.Infof("Command gRPC server is listening on port: %s", s.cfg.GRPC.Port)
		s.log.Fatal(grpcServer.Serve(l))
	}()
	return l.Close, grpcServer, nil
}

func (s *server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()
	s.im = interceptors.NewInterceptorManager(s.log, s.auth)
	s.metrics = metrics.NewCommandServiceMetrics(s.cfg)
	pgxConn, err := postgres.NewPostgresConn(s.cfg.Postgresql)
	if err != nil {
		return errors.Wrap(err, "postgresql.NewPostgresConn")
	}
	s.pgConn = pgxConn
	s.log.Infof("postgres connected: %v", pgxConn.Stat().TotalConns())
	defer pgxConn.Close()
	kafkaProducer := kafkaClient.NewProducer(s.log, s.cfg.Kafka.Brokers)
	defer kafkaProducer.Close() // nolint: errCheck

	userRepo := repositories.NewUserRepository(s.log, s.cfg, pgxConn)
	s.ps = controllers.NewUserService(s.log, s.cfg, userRepo, kafkaProducer)
	productMessageProcessor := kafkaConsumer.NewProductMessageProcessor(s.log, s.cfg, s.v, s.ps, s.metrics)

	s.log.Info("Starting Writer Kafka consumers")
	cg := kafkaClient.NewConsumerGroup(s.cfg.Kafka.Brokers, s.cfg.Kafka.GroupID, s.log)
	go cg.ConsumeTopic(ctx, s.getConsumerGroupTopics(), kafkaConsumer.PoolSize, productMessageProcessor.ProcessMessages)
	closeGrpcServer, grpcServer, err := s.newCommandGrpcServer()
	if err != nil {
		return errors.Wrap(err, "NewScmGrpcServer")
	}
	defer closeGrpcServer() // nolint: errCheck
	if err = s.connectKafkaBrokers(ctx); err != nil {
		return errors.Wrap(err, "s.connectKafkaBrokers")
	}
	defer s.kafkaConn.Close() // nolint: errCheck
	if s.cfg.Kafka.InitTopics {
		s.initKafkaTopics(ctx)
	}
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
	<-ctx.Done()
	grpcServer.GracefulStop()
	return nil
}
