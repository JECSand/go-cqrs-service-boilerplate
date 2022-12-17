package queries

import (
	"context"
	"github.com/JECSand/go-cqrs-service-boilerplate/api_gateway_service/config"
	"github.com/JECSand/go-cqrs-service-boilerplate/api_gateway_service/core/dto"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/tracing"
	readerService "github.com/JECSand/go-cqrs-service-boilerplate/reader_service/proto/product_reader"
	"github.com/opentracing/opentracing-go"
)

type GetUserByIdHandler interface {
	Handle(ctx context.Context, query *GetUserByIdQuery) (*dto.UserResponse, error)
}

type getUserByIdHandler struct {
	log      logging.Logger
	cfg      *config.Config
	rsClient readerService.ReaderServiceClient
}

func NewGetProductByIdHandler(log logger.Logger, cfg *config.Config, rsClient readerService.ReaderServiceClient) *getUserByIdHandler {
	return &getUserByIdHandler{log: log, cfg: cfg, rsClient: rsClient}
}

func (q *getUserByIdHandler) Handle(ctx context.Context, query *GetUserByIdQuery) (*dto.UserResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "getUserByIdHandler.Handle")
	defer span.Finish()

	ctx = tracing.InjectTextMapCarrierToGrpcMetaData(ctx, span.Context())
	res, err := q.rsClient.GetUserById(ctx, &readerService.GetUserByIdReq{ID: query.ID.String()})
	if err != nil {
		return nil, err
	}

	return dto.UserResponseFromGrpc(res.GetUser()), nil
}
