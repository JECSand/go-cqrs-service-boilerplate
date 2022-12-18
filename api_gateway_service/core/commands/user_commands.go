package commands

import (
	"context"
	"github.com/JECSand/go-cqrs-service-boilerplate/api_gateway_service/config"
	kafkaClient "github.com/JECSand/go-cqrs-service-boilerplate/pkg/kafka"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/tracing"
	kafkaMessages "github.com/JECSand/go-cqrs-service-boilerplate/protos/kafka"
	"github.com/opentracing/opentracing-go"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"time"
)

type CreateUserCmdHandler interface {
	Handle(ctx context.Context, command *CreateUserCommand) error
}

type createUserHandler struct {
	log           logging.Logger
	cfg           *config.Config
	kafkaProducer kafkaClient.Producer
}

func NewCreateUserHandler(log logging.Logger, cfg *config.Config, kafkaProducer kafkaClient.Producer) *createUserHandler {
	return &createUserHandler{
		log:           log,
		cfg:           cfg,
		kafkaProducer: kafkaProducer,
	}
}

func (c *createUserHandler) Handle(ctx context.Context, command *CreateUserCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "createUserHandler.Handle")
	defer span.Finish()
	createDTO := &kafkaMessages.UserCreate{
		Email:    command.CreateDto.Email,
		Username: command.CreateDto.Username,
		Password: command.CreateDto.Password,
	}
	dtoBytes, err := proto.Marshal(createDTO)
	if err != nil {
		return err
	}
	return c.kafkaProducer.PublishMessage(ctx, kafka.Message{
		Topic:   c.cfg.KafkaTopics.UserCreate.TopicName,
		Value:   dtoBytes,
		Time:    time.Now().UTC(),
		Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	})
}

// UpdateUserCmdHandler ...
type UpdateUserCmdHandler interface {
	Handle(ctx context.Context, command *UpdateUserCommand) error
}

type updateUserCmdHandler struct {
	log           logging.Logger
	cfg           *config.Config
	kafkaProducer kafkaClient.Producer
}

func NewUpdateUserHandler(log logging.Logger, cfg *config.Config, kafkaProducer kafkaClient.Producer) *updateUserCmdHandler {
	return &updateUserCmdHandler{
		log:           log,
		cfg:           cfg,
		kafkaProducer: kafkaProducer,
	}
}

func (c *updateUserCmdHandler) Handle(ctx context.Context, command *UpdateUserCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "updateUserCmdHandler.Handle")
	defer span.Finish()
	updateDTO := &kafkaMessages.UserUpdate{
		ID:       command.UpdateDto.ID.String(),
		Username: command.UpdateDto.Username,
		Email:    command.UpdateDto.Email,
	}
	dtoBytes, err := proto.Marshal(updateDTO)
	if err != nil {
		return err
	}
	return c.kafkaProducer.PublishMessage(ctx, kafka.Message{
		Topic:   c.cfg.KafkaTopics.UserUpdate.TopicName,
		Value:   dtoBytes,
		Time:    time.Now().UTC(),
		Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	})
}

// DeleteUserCmdHandler ...
type DeleteUserCmdHandler interface {
	Handle(ctx context.Context, command *DeleteUserCommand) error
}

type deleteUserHandler struct {
	log           logging.Logger
	cfg           *config.Config
	kafkaProducer kafkaClient.Producer
}

func NewDeleteProductHandler(log logging.Logger, cfg *config.Config, kafkaProducer kafkaClient.Producer) *deleteUserHandler {
	return &deleteUserHandler{log: log, cfg: cfg, kafkaProducer: kafkaProducer}
}

func (c *deleteUserHandler) Handle(ctx context.Context, command *DeleteUserCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "deleteUserHandler.Handle")
	defer span.Finish()
	deleteDTO := &kafkaMessages.UserDelete{ID: command.ID.String()}
	dtoBytes, err := proto.Marshal(deleteDTO)
	if err != nil {
		return err
	}
	return c.kafkaProducer.PublishMessage(ctx, kafka.Message{
		Topic:   c.cfg.KafkaTopics.UserDelete.TopicName,
		Value:   dtoBytes,
		Time:    time.Now().UTC(),
		Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	})
}
