package repositories

import (
	"context"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/core/entities"
	"github.com/gofrs/uuid"
)

type Repository interface {
	CreateUser(ctx context.Context, product *entities.User) (*entities.User, error)
	UpdateUser(ctx context.Context, product *entities.User) (*entities.User, error)
	DeleteUserByID(ctx context.Context, uuid uuid.UUID) error
	GetUserById(ctx context.Context, uuid uuid.UUID) (*entities.User, error)
}
