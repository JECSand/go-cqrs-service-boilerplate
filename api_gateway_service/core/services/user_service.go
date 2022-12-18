package services

import (
	"github.com/JECSand/go-cqrs-service-boilerplate/api_gateway_service/config"
	"github.com/JECSand/go-cqrs-service-boilerplate/api_gateway_service/core/commands"
	"github.com/JECSand/go-cqrs-service-boilerplate/api_gateway_service/core/queries"
	kafkaClient "github.com/JECSand/go-cqrs-service-boilerplate/pkg/kafka"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
	queryService "github.com/JECSand/go-cqrs-service-boilerplate/query_service/protos/user_query"
)

type UserService struct {
	Commands *commands.UserCommands
	Queries  *queries.UserQueries
}

func NewUserService(log logging.Logger, cfg *config.Config, kafkaProducer kafkaClient.Producer, rsClient queryService.QueryServiceClient) *UserService {
	createUserHandler := commands.NewCreateUserHandler(log, cfg, kafkaProducer)
	updateUserHandler := commands.NewUpdateUserHandler(log, cfg, kafkaProducer)
	deleteUserHandler := commands.NewDeleteProductHandler(log, cfg, kafkaProducer)
	getUserByIdHandler := queries.NewGetProductByIdHandler(log, cfg, rsClient)
	searchUserHandler := queries.NewSearchProductHandler(log, cfg, rsClient)
	UserCommands := commands.NewUserCommands(createUserHandler, updateUserHandler, deleteUserHandler)
	UserQueries := queries.NewUserQueries(getUserByIdHandler, searchUserHandler)
	return &UserService{
		Commands: UserCommands,
		Queries:  UserQueries,
	}
}
