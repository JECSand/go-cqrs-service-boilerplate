package server

import (
	"context"
	"github.com/JECSand/go-cqrs-service-boilerplate/api_gateway_service/config"
	"github.com/JECSand/go-cqrs-service-boilerplate/api_gateway_service/core/client"
	v1 "github.com/JECSand/go-cqrs-service-boilerplate/api_gateway_service/core/controllers/http/v1"
	"github.com/JECSand/go-cqrs-service-boilerplate/api_gateway_service/core/metrics"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/authentication"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/interceptors"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/kafka"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/tracing"
	"github.com/go-playground/validator"
	"github.com/heptiolabs/healthcheck"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/swaggo/swag/example/basic/docs"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	maxHeaderBytes = 1 << 20
	stackSize      = 1 << 10 // 1 KB
	bodyLimit      = "2M"
	readTimeout    = 15 * time.Second
	writeTimeout   = 15 * time.Second
	gzipLevel      = 5
)

type server struct {
	log  logging.Logger
	auth authentication.Authenticator
	cfg  *config.Config
	v    *validator.Validate
	mw   middleware.MiddlewareManager
	im   interceptors.InterceptorManager
	echo *echo.Echo
	ps   *service.ProductService
	m    *metrics.ApiGatewayMetrics
}

func NewServer(log logging.Logger, auth authentication.Authenticator, cfg *config.Config) *server {
	return &server{
		log:  log,
		auth: auth,
		cfg:  cfg,
		echo: echo.New(), v: validator.New(),
	}
}

func (s *server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()
	s.mw = middleware.NewMiddlewareManager(s.log, s.cfg)
	s.im = interceptors.NewInterceptorManager(s.log, s.auth)
	s.m = metrics.NewApiGatewayMetrics(s.cfg)

	queryService, err := client.NewQueryServiceClient(ctx, s.cfg, s.im)
	if err != nil {
		return err
	}
	defer readerServiceConn.Close() // nolint: errcheck
	rsClient := readerService.NewReaderServiceClient(readerServiceConn)
	kafkaProducer := kafka.NewProducer(s.log, s.cfg.Kafka.Brokers)
	defer kafkaProducer.Close() // nolint: errcheck
	s.ps = service.NewProductService(s.log, s.cfg, kafkaProducer, rsClient)
	productHandlers := v1.NewProductsHandlers(s.echo.Group(s.cfg.Http.ProductsPath), s.log, s.mw, s.cfg, s.ps, s.v, s.m)
	productHandlers.MapRoutes()
	go func() {
		if err := s.runHttpServer(); err != nil {
			s.log.Errorf(" s.runHttpServer: %v", err)
			cancel()
		}
	}()
	s.log.Infof("API Gateway is listening on PORT: %s", s.cfg.Http.Port)
	s.runMetrics(cancel)
	s.runHealthCheck(ctx)
	if s.cfg.Jaeger.Enable {
		tracer, closer, err := tracing.NewJaegerTracer(s.cfg.Jaeger)
		if err != nil {
			return err
		}
		defer closer.Close() // nolint: errcheck
		opentracing.SetGlobalTracer(tracer)
	}
	<-ctx.Done()
	if err := s.echo.Server.Shutdown(ctx); err != nil {
		s.log.WarnMsg("echo.Server.Shutdown", err)
	}
	return nil
}

func (s *server) runHealthCheck(ctx context.Context) {
	health := healthcheck.NewHandler()

	health.AddReadinessCheck(s.cfg.ServiceName, healthcheck.AsyncWithContext(ctx, func() error {
		if s.cfg != nil {
			return nil
		}
		return errors.New("Config not loaded")
	}, time.Duration(s.cfg.Probes.CheckIntervalSeconds)*time.Second))

	go func() {
		s.log.Infof("API_Gateway Kubernetes probes listening on port: %s", s.cfg.Probes.Port)
		if err := http.ListenAndServe(s.cfg.Probes.Port, health); err != nil {
			s.log.WarnMsg("ListenAndServe", err)
		}
	}()
}

func (s *server) runMetrics(cancel context.CancelFunc) {
	metricsServer := echo.New()
	metricsServer.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         stackSize,
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	go func() {
		metricsServer.GET(s.cfg.Probes.PrometheusPath, echo.WrapHandler(promhttp.Handler()))
		s.log.Infof("Metrics server is running on port: %s", s.cfg.Probes.PrometheusPort)
		if err := metricsServer.Start(s.cfg.Probes.PrometheusPort); err != nil {
			s.log.Errorf("metricsServer.Start: %v", err)
			cancel()
		}
	}()
}

func (s *server) runHttpServer() error {
	s.mapRoutes()
	s.echo.Server.ReadTimeout = readTimeout
	s.echo.Server.WriteTimeout = writeTimeout
	s.echo.Server.MaxHeaderBytes = maxHeaderBytes
	return s.echo.Start(s.cfg.Http.Port)
}

func (s *server) mapRoutes() {
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "API Gateway"
	docs.SwaggerInfo.Description = "API Gateway CQRS microservices."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"
	s.echo.GET("/swagger/*", echoSwagger.WrapHandler)
	s.echo.Use(s.mw.RequestLoggerMiddleware)
	s.echo.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         stackSize,
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	s.echo.Use(middleware.RequestID())
	s.echo.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: gzipLevel,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))
	s.echo.Use(middleware.BodyLimit(bodyLimit))
}
