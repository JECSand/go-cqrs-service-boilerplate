package grpc

import (
	"context"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/config"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/core/commands"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/core/controllers"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/core/metrics"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/core/queries"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/mappings"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/protos/user_command"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/tracing"
	"github.com/go-playground/validator"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcService struct {
	log     logging.Logger
	cfg     *config.Config
	v       *validator.Validate
	ps      *controllers.UserService
	metrics *metrics.CommandServiceMetrics
}

func NewCommandGrpcService(log logging.Logger, cfg *config.Config, v *validator.Validate, ps *controllers.UserService, metrics *metrics.CommandServiceMetrics) *grpcService {
	return &grpcService{
		log:     log,
		cfg:     cfg,
		v:       v,
		ps:      ps,
		metrics: metrics,
	}
}

func (s *grpcService) CreateUser(ctx context.Context, req *commandService.CreateUserReq) (*commandService.CreateUserRes, error) {
	s.metrics.CreateUserGrpcRequests.Inc()
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "grpcService.CreateUser")
	defer span.Finish()
	id, err := uuid.FromString(req.GetID())
	if err != nil {
		s.log.WarnMsg("uuid.FromString", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}
	// TODO: Add logic to manage new User Active and Root fields
	command := commands.NewCreateUserCommand(id, req.GetEmail(), req.GetUsername(), req.GetPassword(), false, false)
	if err = s.v.StructCtx(ctx, command); err != nil {
		s.log.WarnMsg("validate", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}
	err = s.ps.Commands.CreateUser.Handle(ctx, command)
	if err != nil {
		s.log.WarnMsg("CreateUser.Handle", err)
		return nil, s.errResponse(codes.Internal, err)
	}
	s.metrics.SuccessGrpcRequests.Inc()
	return &commandService.CreateUserRes{ID: id.String()}, nil
}

func (s *grpcService) UpdateUser(ctx context.Context, req *commandService.UpdateUserReq) (*commandService.UpdateUserRes, error) {
	s.metrics.UpdateUserGrpcRequests.Inc()
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "grpcService.UpdateUser")
	defer span.Finish()
	id, err := uuid.FromString(req.GetID())
	if err != nil {
		s.log.WarnMsg("uuid.FromString", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}
	command := commands.NewUpdateUserCommand(id, req.GetEmail(), req.GetUsername())
	if err = s.v.StructCtx(ctx, command); err != nil {
		s.log.WarnMsg("validate", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}
	err = s.ps.Commands.UpdateUser.Handle(ctx, command)
	if err != nil {
		s.log.WarnMsg("UpdateProduct.Handle", err)
		return nil, s.errResponse(codes.Internal, err)
	}
	s.metrics.SuccessGrpcRequests.Inc()
	return &commandService.UpdateUserRes{}, nil
}

func (s *grpcService) GetUserById(ctx context.Context, req *commandService.GetUserByIdReq) (*commandService.GetUserByIdRes, error) {
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
	found, err := s.ps.Queries.GetUserById.Handle(ctx, query)
	if err != nil {
		s.log.WarnMsg("GetUserById.Handle", err)
		return nil, s.errResponse(codes.Internal, err)
	}
	s.metrics.SuccessGrpcRequests.Inc()
	return &commandService.GetUserByIdRes{User: mappings.CommandUserToGrpc(found)}, nil
}

func (s *grpcService) errResponse(c codes.Code, err error) error {
	s.metrics.ErrorGrpcRequests.Inc()
	return status.Error(c, err.Error())
}