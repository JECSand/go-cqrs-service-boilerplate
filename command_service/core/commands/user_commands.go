package commands

import (
	"context"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/config"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/core/entities"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/core/repositories"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/mappings"
	kafkaClient "github.com/JECSand/go-cqrs-service-boilerplate/pkg/kafka"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/tracing"
	kafkaMessages "github.com/JECSand/go-cqrs-service-boilerplate/protos/kafka"
	"github.com/opentracing/opentracing-go"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"time"
)

// CreateUserCmdHandler ...
type CreateUserCmdHandler interface {
	Handle(ctx context.Context, command *CreateUserCommand) error
}

type createUserHandler struct {
	log           logging.Logger
	cfg           *config.Config
	pgRepo        repositories.Repository
	kafkaProducer kafkaClient.Producer
}

func NewCreateUserHandler(log logging.Logger, cfg *config.Config, pgRepo repositories.Repository, kafkaProducer kafkaClient.Producer) *createUserHandler {
	return &createUserHandler{
		log:           log,
		cfg:           cfg,
		pgRepo:        pgRepo,
		kafkaProducer: kafkaProducer,
	}
}

func (c *createUserHandler) Handle(ctx context.Context, command *CreateUserCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "createUserHandler.Handle")
	defer span.Finish()
	userDTO := &entities.User{
		ID:       command.ID,
		Email:    command.Email,
		Username: command.Username,
		Password: command.Password,
		Root:     command.Root,
		Active:   command.Active,
	}
	user, err := c.pgRepo.CreateUser(ctx, userDTO)
	if err != nil {
		return err
	}
	msg := &kafkaMessages.UserCreated{User: mappings.UserToGrpcMessage(user)}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	message := kafka.Message{
		Topic:   c.cfg.KafkaTopics.UserCreated.TopicName,
		Value:   msgBytes,
		Time:    time.Now().UTC(),
		Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	}
	return c.kafkaProducer.PublishMessage(ctx, message)
}

// UpdateUserCmdHandler ...
type UpdateUserCmdHandler interface {
	Handle(ctx context.Context, command *UpdateUserCommand) error
}

type updateUserHandler struct {
	log           logging.Logger
	cfg           *config.Config
	pgRepo        repositories.Repository
	kafkaProducer kafkaClient.Producer
}

func NewUpdateUserHandler(log logging.Logger, cfg *config.Config, pgRepo repositories.Repository, kafkaProducer kafkaClient.Producer) *updateUserHandler {
	return &updateUserHandler{log: log,
		cfg:           cfg,
		pgRepo:        pgRepo,
		kafkaProducer: kafkaProducer,
	}
}

func (c *updateUserHandler) Handle(ctx context.Context, command *UpdateUserCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "updateUserHandler.Handle")
	defer span.Finish()
	userDTO := &entities.User{
		ID:       command.ID,
		Email:    command.Email,
		Username: command.Username,
	}
	user, err := c.pgRepo.UpdateUser(ctx, userDTO)
	if err != nil {
		return err
	}
	msg := &kafkaMessages.UserUpdated{User: mappings.UserToGrpcMessage(user)}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	message := kafka.Message{
		Topic:   c.cfg.KafkaTopics.UserUpdated.TopicName,
		Value:   msgBytes,
		Time:    time.Now().UTC(),
		Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	}
	return c.kafkaProducer.PublishMessage(ctx, message)
}

// DeleteUserCmdHandler ...
type DeleteUserCmdHandler interface {
	Handle(ctx context.Context, command *DeleteUserCommand) error
}

type deleteUserHandler struct {
	log           logging.Logger
	cfg           *config.Config
	pgRepo        repositories.Repository
	kafkaProducer kafkaClient.Producer
}

func NewDeleteUserHandler(log logging.Logger, cfg *config.Config, pgRepo repositories.Repository, kafkaProducer kafkaClient.Producer) *deleteUserHandler {
	return &deleteUserHandler{
		log:           log,
		cfg:           cfg,
		pgRepo:        pgRepo,
		kafkaProducer: kafkaProducer,
	}
}

func (c *deleteUserHandler) Handle(ctx context.Context, command *DeleteUserCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "deleteUserHandler.Handle")
	defer span.Finish()
	if err := c.pgRepo.DeleteUserByID(ctx, command.ID); err != nil {
		return err
	}
	msg := &kafkaMessages.UserDeleted{ID: command.ID.String()}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	message := kafka.Message{
		Topic:   c.cfg.KafkaTopics.UserDeleted.TopicName,
		Value:   msgBytes,
		Time:    time.Now().UTC(),
		Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	}
	return c.kafkaProducer.PublishMessage(ctx, message)
}
