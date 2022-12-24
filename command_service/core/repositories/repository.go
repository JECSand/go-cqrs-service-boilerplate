package repositories

import (
	"context"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/core/entities"
	"github.com/gofrs/uuid"
)

type Repository interface {
	CreateUser(ctx context.Context, user *entities.User) (*entities.User, error)
	UpdateUser(ctx context.Context, user *entities.User) (*entities.User, error)
	DeleteUserByID(ctx context.Context, id uuid.UUID) error
	GetUserById(ctx context.Context, id uuid.UUID) (*entities.User, error)
}
