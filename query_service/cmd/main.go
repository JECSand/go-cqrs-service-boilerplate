package cmd

import (
	"flag"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/config"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/server"
	"log"
)

func main() {
	flag.Parse()
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	appLogger := logging.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()
	appLogger.WithName("QueryService")
	s := server.NewServer(appLogger, cfg)
	appLogger.Fatal(s.Run())
}
