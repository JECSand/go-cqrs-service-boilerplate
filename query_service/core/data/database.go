package data

import (
	"context"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/utilities"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/models"
	"github.com/gofrs/uuid"
)

type Database interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, uuid uuid.UUID) error
	GetUserById(ctx context.Context, uuid uuid.UUID) (*models.User, error)
	Search(ctx context.Context, search string, pagination *utilities.Pagination) (*models.UsersList, error)
}
