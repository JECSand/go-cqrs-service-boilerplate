package kafka

import (
	"context"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/config"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/core/commands"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/core/controllers"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/core/metrics"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/tracing"
	kafkaMessages "github.com/JECSand/go-cqrs-service-boilerplate/protos/kafka"
	"github.com/avast/retry-go"
	"github.com/go-playground/validator"
	"github.com/gofrs/uuid"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"sync"
	"time"
)

const (
	PoolSize      = 30
	retryAttempts = 3
	retryDelay    = 300 * time.Millisecond
)

var (
	retryOptions = []retry.Option{retry.Attempts(retryAttempts), retry.Delay(retryDelay), retry.DelayType(retry.BackOffDelay)}
)

type userMessageProcessor struct {
	log     logging.Logger
	cfg     *config.Config
	v       *validator.Validate
	ps      *controllers.UserService
	metrics *metrics.CommandServiceMetrics
}

func NewProductMessageProcessor(log logging.Logger, cfg *config.Config, v *validator.Validate, ps *controllers.UserService, metrics *metrics.CommandServiceMetrics) *userMessageProcessor {
	return &userMessageProcessor{
		log:     log,
		cfg:     cfg,
		v:       v,
		ps:      ps,
		metrics: metrics,
	}
}

func (s *userMessageProcessor) commitMessage(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.metrics.SuccessKafkaMessages.Inc()
	s.log.KafkaLogCommittedMessage(m.Topic, m.Partition, m.Offset)
	if err := r.CommitMessages(ctx, m); err != nil {
		s.log.WarnMsg("commitMessage", err)
	}
}

func (s *userMessageProcessor) commitErrMessage(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.metrics.ErrorKafkaMessages.Inc()
	s.log.KafkaLogCommittedMessage(m.Topic, m.Partition, m.Offset)
	if err := r.CommitMessages(ctx, m); err != nil {
		s.log.WarnMsg("commitMessage", err)
	}
}

func (s *userMessageProcessor) logProcessMessage(m kafka.Message, workerID int) {
	s.log.KafkaProcessMessage(m.Topic, m.Partition, string(m.Value), workerID, m.Offset, m.Time)
}

func (s *userMessageProcessor) processCreateUser(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.metrics.CreateUserKafkaMessages.Inc()
	ctx, span := tracing.StartKafkaConsumerTracerSpan(ctx, m.Headers, "userMessageProcessor.processCreateUser")
	defer span.Finish()
	var msg kafkaMessages.UserCreate
	if err := proto.Unmarshal(m.Value, &msg); err != nil {
		s.log.WarnMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}
	id, err := uuid.FromString(msg.GetID())
	if err != nil {
		s.log.WarnMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}
	// TODO: Add logic to manage a new user's root and active fields
	command := commands.NewCreateUserCommand(id, msg.GetEmail(), msg.GetUsername(), msg.GetPassword(), false, false)
	if err = s.v.StructCtx(ctx, command); err != nil {
		s.log.WarnMsg("validate", err)
		s.commitErrMessage(ctx, r, m)
		return
	}
	if err = retry.Do(func() error {
		return s.ps.Commands.CreateUser.Handle(ctx, command)
	}, append(retryOptions, retry.Context(ctx))...); err != nil {
		s.log.WarnMsg("CreateUser.Handle", err)
		s.metrics.ErrorKafkaMessages.Inc()
		return
	}
	s.commitMessage(ctx, r, m)
}

func (s *userMessageProcessor) processUpdateUser(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.metrics.UpdateUserKafkaMessages.Inc()
	ctx, span := tracing.StartKafkaConsumerTracerSpan(ctx, m.Headers, "userMessageProcessor.processUpdateUser")
	defer span.Finish()
	msg := &kafkaMessages.UserUpdate{}
	if err := proto.Unmarshal(m.Value, msg); err != nil {
		s.log.WarnMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}
	id, err := uuid.FromString(msg.GetID())
	if err != nil {
		s.log.WarnMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}
	command := commands.NewUpdateUserCommand(id, msg.GetEmail(), msg.GetUsername())
	if err = s.v.StructCtx(ctx, command); err != nil {
		s.log.WarnMsg("validate", err)
		s.commitErrMessage(ctx, r, m)
		return
	}
	if err = retry.Do(func() error {
		return s.ps.Commands.UpdateUser.Handle(ctx, command)
	}, append(retryOptions, retry.Context(ctx))...); err != nil {
		s.log.WarnMsg("UpdateUser.Handle", err)
		s.metrics.ErrorKafkaMessages.Inc()
		return
	}
	s.commitMessage(ctx, r, m)
}

func (s *userMessageProcessor) processDeleteUser(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.metrics.DeleteUserKafkaMessages.Inc()
	ctx, span := tracing.StartKafkaConsumerTracerSpan(ctx, m.Headers, "userMessageProcessor.processDeleteUser")
	defer span.Finish()
	msg := &kafkaMessages.UserDelete{}
	if err := proto.Unmarshal(m.Value, msg); err != nil {
		s.log.WarnMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}
	id, err := uuid.FromString(msg.GetID())
	if err != nil {
		s.log.WarnMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}
	command := commands.NewDeleteUserCommand(id)
	if err = s.v.StructCtx(ctx, command); err != nil {
		s.log.WarnMsg("validate", err)
		s.commitErrMessage(ctx, r, m)
		return
	}
	if err = retry.Do(func() error {
		return s.ps.Commands.DeleteUser.Handle(ctx, command)
	}, append(retryOptions, retry.Context(ctx))...); err != nil {
		s.log.WarnMsg("DeleteUser.Handle", err)
		s.metrics.ErrorKafkaMessages.Inc()
		return
	}
	s.commitMessage(ctx, r, m)
}

func (s *userMessageProcessor) ProcessMessages(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		m, err := r.FetchMessage(ctx)
		if err != nil {
			s.log.Warnf("workerID: %v, err: %v", workerID, err)
			continue
		}
		s.logProcessMessage(m, workerID)
		switch m.Topic {
		case s.cfg.KafkaTopics.UserCreate.TopicName:
			s.processCreateUser(ctx, r, m)
		case s.cfg.KafkaTopics.UserUpdate.TopicName:
			s.processUpdateUser(ctx, r, m)
		case s.cfg.KafkaTopics.UserDelete.TopicName:
			s.processDeleteUser(ctx, r, m)
		}
	}
}
