package queries

import (
	"context"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/config"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/core/entities"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/core/repositories"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
)

/*
GET USER BY ID
*/

// GetUserByIdHandler ...
type GetUserByIdHandler interface {
	Handle(ctx context.Context, query *GetUserByIdQuery) (*entities.User, error)
}

type getUserByIdHandler struct {
	log    logging.Logger
	cfg    *config.Config
	pgRepo repositories.Repository
}

// NewGetUserByIdHandler ...
func NewGetUserByIdHandler(log logging.Logger, cfg *config.Config, pgRepo repositories.Repository) *getUserByIdHandler {
	return &getUserByIdHandler{
		log:    log,
		cfg:    cfg,
		pgRepo: pgRepo,
	}
}

// Handle ...
func (q *getUserByIdHandler) Handle(ctx context.Context, query *GetUserByIdQuery) (*entities.User, error) {
	return q.pgRepo.GetUserById(ctx, query.ID)
}

/*
COUNT USERS
*/

// CountUsersHandler ...
type CountUsersHandler interface {
	Handle(ctx context.Context) (int, error)
}

type countUsersHandler struct {
	log    logging.Logger
	cfg    *config.Config
	pgRepo repositories.Repository
}

// NewCountUsersHandler ...
func NewCountUsersHandler(log logging.Logger, cfg *config.Config, pgRepo repositories.Repository) *countUsersHandler {
	return &countUsersHandler{
		log:    log,
		cfg:    cfg,
		pgRepo: pgRepo,
	}
}

// Handle ...
func (q *countUsersHandler) Handle(ctx context.Context) (int, error) {
	return q.pgRepo.CountUsers(ctx)
}
