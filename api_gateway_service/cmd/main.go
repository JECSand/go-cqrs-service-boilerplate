package cmd

import (
	"flag"
	"github.com/JECSand/go-cqrs-service-boilerplate/api_gateway_service/config"
	"github.com/JECSand/go-cqrs-service-boilerplate/api_gateway_service/core/controllers/access"
	"github.com/JECSand/go-cqrs-service-boilerplate/api_gateway_service/core/server"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/authentication"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
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
	logger.WithName("ApiGateway")
	auth := authentication.NewAuthenticator(logger, access.DefaultAccessRules())
	s := server.NewServer(logger, auth, cfg)
	logger.Fatal(s.Run())
}
