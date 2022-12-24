package cache

import (
	"context"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/models"
)

type Cache interface {
	PutUser(ctx context.Context, key string, user *models.User)
	GetUser(ctx context.Context, key string) (*models.User, error)
	DeleteUser(ctx context.Context, key string)
	DeleteAllUsers(ctx context.Context)
}
