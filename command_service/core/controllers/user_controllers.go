package controllers

import (
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/config"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/core/commands"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/core/queries"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/core/repositories"
	kafkaClient "github.com/JECSand/go-cqrs-service-boilerplate/pkg/kafka"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
)

// UserService ...
type UserService struct {
	Commands *commands.UserCommands
	Queries  *queries.UserQueries
}

// NewUserService ...
func NewUserService(log logging.Logger, cfg *config.Config, pgRepo repositories.Repository, kafkaProducer kafkaClient.Producer) *UserService {
	updateUserHandler := commands.NewUpdateUserHandler(log, cfg, pgRepo, kafkaProducer)
	createUserHandler := commands.NewCreateUserHandler(log, cfg, pgRepo, kafkaProducer)
	deleteUserHandler := commands.NewDeleteUserHandler(log, cfg, pgRepo, kafkaProducer)
	getUserHByIdHandler := queries.NewGetUserByIdHandler(log, cfg, pgRepo)
	countUsersHandler := queries.NewCountUsersHandler(log, cfg, pgRepo)
	userCommands := commands.NewUserCommands(createUserHandler, updateUserHandler, deleteUserHandler)
	userQueries := queries.NewUserQueries(getUserHByIdHandler, countUsersHandler)
	return &UserService{
		Commands: userCommands,
		Queries:  userQueries,
	}
}
