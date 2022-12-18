package grpc

import (
	"context"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/tracing"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/utilities"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/config"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/events"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/metrics"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/models"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/queries"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/services"
	queryService "github.com/JECSand/go-cqrs-service-boilerplate/query_service/protos/user_query"
	"github.com/go-playground/validator"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type grpcService struct {
	log     logging.Logger
	cfg     *config.Config
	v       *validator.Validate
	ps      *services.UserService
	metrics *metrics.QueryServiceMetrics
}

func NewQueryGrpcService(log logging.Logger, cfg *config.Config, v *validator.Validate, ps *services.UserService, metrics *metrics.QueryServiceMetrics) *grpcService {
	return &grpcService{
		log:     log,
		cfg:     cfg,
		v:       v,
		ps:      ps,
		metrics: metrics,
	}
}

func (s *grpcService) CreateUser(ctx context.Context, req *queryService.CreateUserReq) (*queryService.CreateUserRes, error) {
	s.metrics.CreateUserGrpcRequests.Inc()
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "grpcService.CreateUser")
	defer span.Finish()
	// TODO - ADD LOGIC FOR ROOT AND ACTIVE BELOW
	event := events.NewCreateUserEvent(req.GetID(), req.GetEmail(), req.GetUsername(), req.GetPassword(), false, false, time.Now(), time.Now())
	if err := s.v.StructCtx(ctx, event); err != nil {
		s.log.WarnMsg("validate", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}
	if err := s.ps.Events.CreateUser.Handle(ctx, event); err != nil {
		s.log.WarnMsg("CreateUser.Handle", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}
	s.metrics.SuccessGrpcRequests.Inc()
	return &queryService.CreateUserRes{ID: req.GetID()}, nil
}

func (s *grpcService) UpdateUser(ctx context.Context, req *queryService.UpdateUserReq) (*queryService.UpdateUserRes, error) {
	s.metrics.UpdateUserGrpcRequests.Inc()
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "grpcService.UpdateUser")
	defer span.Finish()
	command := events.NewUpdateUserEvent(req.GetID(), req.GetEmail(), req.GetUsername(), time.Now())
	if err := s.v.StructCtx(ctx, command); err != nil {
		s.log.WarnMsg("validate", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}
	if err := s.ps.Events.UpdateUser.Handle(ctx, command); err != nil {
		s.log.WarnMsg("UpdateUser.Handle", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}
	s.metrics.SuccessGrpcRequests.Inc()
	return &queryService.UpdateUserRes{ID: req.GetID()}, nil
}

func (s *grpcService) GetUserById(ctx context.Context, req *queryService.GetUserByIdReq) (*queryService.GetUserByIdRes, error) {
	s.metrics.GetUserByIdGrpcRequests.Inc()
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "grpcService.GetUserById")
	defer span.Finish()
	id, err := uuid.FromString(req.GetID())
	if err != nil {
		s.log.WarnMsg("uuid.FromString", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}
	query := queries.NewGetUserByIdQuery(id)
	if err = s.v.StructCtx(ctx, query); err != nil {
		s.log.WarnMsg("validate", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}
	user, err := s.ps.Queries.GetUserById.Handle(ctx, query)
	if err != nil {
		s.log.WarnMsg("GetUserById.Handle", err)
		return nil, s.errResponse(codes.Internal, err)
	}
	s.metrics.SuccessGrpcRequests.Inc()
	return &queryService.GetUserByIdRes{User: models.UserToGrpcMessage(user)}, nil
}

func (s *grpcService) SearchUser(ctx context.Context, req *queryService.SearchReq) (*queryService.SearchRes, error) {
	s.metrics.SearchUserGrpcRequests.Inc()
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "grpcService.SearchUser")
	defer span.Finish()
	pq := utilities.NewPaginationQuery(int(req.GetSize()), int(req.GetPage()))
	query := queries.NewSearchUserQuery(req.GetSearch(), pq)
	productsList, err := s.ps.Queries.SearchUser.Handle(ctx, query)
	if err != nil {
		s.log.WarnMsg("SearchUser.Handle", err)
		return nil, s.errResponse(codes.Internal, err)
	}
	s.metrics.SuccessGrpcRequests.Inc()
	return models.UserListToGrpc(productsList), nil
}

func (s *grpcService) DeleteUserByID(ctx context.Context, req *queryService.DeleteUserByIdReq) (*queryService.DeleteUserByIdRes, error) {
	s.metrics.DeleteUserGrpcRequests.Inc()
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "grpcService.DeleteUserByID")
	defer span.Finish()

	id, err := uuid.FromString(req.GetID())
	if err != nil {
		s.log.WarnMsg("uuid.FromString", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}
	if err = s.ps.Events.DeleteUser.Handle(ctx, events.NewDeleteUserEvent(id)); err != nil {
		s.log.WarnMsg("DeleteUser.Handle", err)
		return nil, s.errResponse(codes.Internal, err)
	}
	s.metrics.SuccessGrpcRequests.Inc()
	return &queryService.DeleteUserByIdRes{}, nil
}

func (s *grpcService) errResponse(c codes.Code, err error) error {
	s.metrics.ErrorGrpcRequests.Inc()
	return status.Error(c, err.Error())
}
