package queries

import (
	"context"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/config"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/core/entities"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/core/repositories"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
)

// GetUserByIdHandler ...
type GetUserByIdHandler interface {
	Handle(ctx context.Context, query *GetUserByIdQuery) (*entities.User, error)
}

type getUserByIdHandler struct {
	log    logging.Logger
	cfg    *config.Config
	pgRepo repositories.Repository
}

func NewGetUserByIdHandler(log logging.Logger, cfg *config.Config, pgRepo repositories.Repository) *getUserByIdHandler {
	return &getUserByIdHandler{
		log:    log,
		cfg:    cfg,
		pgRepo: pgRepo,
	}
}

func (q *getUserByIdHandler) Handle(ctx context.Context, query *GetUserByIdQuery) (*entities.User, error) {
	return q.pgRepo.GetUserById(ctx, query.ID)
}
