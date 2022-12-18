package cmd

import (
	"flag"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/authentication"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/config"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/delivery/access"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/server"
	"log"
)

func main() {
	flag.Parse()
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	logger := logging.NewAppLogger(cfg.Logger)
	logger.InitLogger()
	logger.WithName("QueryService")
	auth := authentication.NewAuthenticator(logger, access.DefaultAccessRules())
	s := server.NewServer(logger, auth, cfg)
	logger.Fatal(s.Run())
}
