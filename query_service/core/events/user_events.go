package events

import (
	"context"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/config"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/cache"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/data"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/models"
	"github.com/opentracing/opentracing-go"
)

// CreateUserEventHandler ...
type CreateUserEventHandler interface {
	Handle(ctx context.Context, event *CreateUserEvent) error
}

type createUserEventHandler struct {
	log        logging.Logger
	cfg        *config.Config
	mongoDB    data.Database
	redisCache cache.Cache
}

func NewCreateUserEventHandler(log logging.Logger, cfg *config.Config, mongoDB data.Database, redisCache cache.Cache) *createUserEventHandler {
	return &createUserEventHandler{
		log:        log,
		cfg:        cfg,
		mongoDB:    mongoDB,
		redisCache: redisCache,
	}
}

func (c *createUserEventHandler) Handle(ctx context.Context, event *CreateUserEvent) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "createUserEventHandler.Handle")
	defer span.Finish()
	user := &models.User{
		ID:        event.ID,
		Email:     event.Email,
		Username:  event.Username,
		Password:  event.Password,
		Root:      event.Root,
		Active:    event.Active,
		CreatedAt: event.CreatedAt,
		UpdatedAt: event.UpdatedAt,
	}
	created, err := c.mongoDB.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	c.redisCache.PutUser(ctx, created.ID, created)
	return nil
}

// UpdateUserEventHandler ...
type UpdateUserEventHandler interface {
	Handle(ctx context.Context, event *UpdateUserEvent) error
}

type updateUserEventHandler struct {
	log        logging.Logger
	cfg        *config.Config
	mongoDB    data.Database
	redisCache cache.Cache
}

func NewUpdateUserEventHandler(log logging.Logger, cfg *config.Config, mongoDB data.Database, redisCache cache.Cache) *updateUserEventHandler {
	return &updateUserEventHandler{
		log:        log,
		cfg:        cfg,
		mongoDB:    mongoDB,
		redisCache: redisCache,
	}
}

func (c *updateUserEventHandler) Handle(ctx context.Context, event *UpdateUserEvent) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "updateUserEventHandler.Handle")
	defer span.Finish()
	user := &models.User{
		ID:        event.ID,
		Email:     event.Email,
		Username:  event.Username,
		UpdatedAt: event.UpdatedAt,
	}
	updated, err := c.mongoDB.UpdateUser(ctx, user)
	if err != nil {
		return err
	}
	c.redisCache.PutUser(ctx, updated.ID, updated)
	return nil
}

// DeleteUserEventHandler ...
type DeleteUserEventHandler interface {
	Handle(ctx context.Context, event *DeleteUserEvent) error
}

type deleteUserEventHandler struct {
	log        logging.Logger
	cfg        *config.Config
	mongoDB    data.Database
	redisCache cache.Cache
}

func NewDeleteUserEventHandler(log logging.Logger, cfg *config.Config, mongoDB data.Database, redisCache cache.Cache) *deleteUserEventHandler {
	return &deleteUserEventHandler{
		log:        log,
		cfg:        cfg,
		mongoDB:    mongoDB,
		redisCache: redisCache,
	}
}

func (c *deleteUserEventHandler) Handle(ctx context.Context, event *DeleteUserEvent) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "deleteUserEventHandler.Handle")
	defer span.Finish()
	if err := c.mongoDB.DeleteUser(ctx, event.ID); err != nil {
		return err
	}
	c.redisCache.DeleteUser(ctx, event.ID.String())
	return nil
}
