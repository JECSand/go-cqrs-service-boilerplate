package queries

import (
	"context"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/config"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/cache"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/data"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/models"
	"github.com/opentracing/opentracing-go"
)

// GetUserByIdHandler ...
type GetUserByIdHandler interface {
	Handle(ctx context.Context, query *GetUserByIdQuery) (*models.User, error)
}

type getUserByIdHandler struct {
	log        logging.Logger
	cfg        *config.Config
	mongoDB    data.Database
	redisCache cache.Cache
}

func NewGetUserByIdHandler(log logging.Logger, cfg *config.Config, mongoDB data.Database, redisCache cache.Cache) *getUserByIdHandler {
	return &getUserByIdHandler{
		log:        log,
		cfg:        cfg,
		mongoDB:    mongoDB,
		redisCache: redisCache,
	}
}

func (q *getUserByIdHandler) Handle(ctx context.Context, query *GetUserByIdQuery) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "getUserByIdHandler.Handle")
	defer span.Finish()
	if product, err := q.redisCache.GetUser(ctx, query.ID.String()); err == nil && product != nil {
		return product, nil
	}
	product, err := q.mongoDB.GetUserById(ctx, query.ID)
	if err != nil {
		return nil, err
	}
	q.redisCache.PutUser(ctx, product.ID, product)
	return product, nil
}

// SearchUserHandler ...
type SearchUserHandler interface {
	Handle(ctx context.Context, query *SearchUserQuery) (*models.UsersList, error)
}

type searchUserHandler struct {
	log        logging.Logger
	cfg        *config.Config
	mongoDB    data.Database
	redisCache cache.Cache
}

func NewSearchUserHandler(log logging.Logger, cfg *config.Config, mongoDB data.Database, redisCache cache.Cache) *searchUserHandler {
	return &searchUserHandler{
		log:        log,
		cfg:        cfg,
		mongoDB:    mongoDB,
		redisCache: redisCache,
	}
}

func (s *searchUserHandler) Handle(ctx context.Context, query *SearchUserQuery) (*models.UsersList, error) {
	return s.mongoDB.Search(ctx, query.Text, query.Pagination)
}
