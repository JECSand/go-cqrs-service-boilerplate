package cmd

import (
	"flag"
	"github.com/JECSand/go-cqrs-service-boilerplate/api_gateway_service/config"
	"log"
)

func main() {
	flag.Parse()
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()
	appLogger.WithName("ApiGateway")
	s := server.NewServer(appLogger, cfg)
	appLogger.Fatal(s.Run())
}
