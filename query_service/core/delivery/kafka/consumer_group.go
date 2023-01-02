package kafka

import (
	"context"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/tracing"
	kafkaMessages "github.com/JECSand/go-cqrs-service-boilerplate/protos/kafka"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/config"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/events"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/metrics"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/services"
	"github.com/avast/retry-go"
	"github.com/go-playground/validator"
	"github.com/gofrs/uuid"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"sync"
	"time"
)

const (
	PoolSize = 30
)

const (
	retryAttempts = 3
	retryDelay    = 300 * time.Millisecond
)

var (
	retryOptions = []retry.Option{retry.Attempts(retryAttempts), retry.Delay(retryDelay), retry.DelayType(retry.BackOffDelay)}
)

type queryMessageProcessor struct {
	log     logging.Logger
	cfg     *config.Config
	v       *validator.Validate
	ps      *services.UserService
	metrics *metrics.QueryServiceMetrics
}

func NewQueryMessageProcessor(log logging.Logger, cfg *config.Config, v *validator.Validate, ps *services.UserService, metrics *metrics.QueryServiceMetrics) *queryMessageProcessor {
	return &queryMessageProcessor{
		log:     log,
		cfg:     cfg,
		v:       v,
		ps:      ps,
		metrics: metrics,
	}
}

func (s *queryMessageProcessor) ProcessMessages(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int) {
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
		case s.cfg.KafkaTopics.UserCreated.TopicName:
			s.processUserCreated(ctx, r, m)
		case s.cfg.KafkaTopics.UserUpdated.TopicName:
			s.processUserUpdated(ctx, r, m)
		case s.cfg.KafkaTopics.UserDeleted.TopicName:
			s.processUserDeleted(ctx, r, m)
		}
	}
}

func (s *queryMessageProcessor) processUserCreated(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.metrics.CreateUserKafkaMessages.Inc()
	ctx, span := tracing.StartKafkaConsumerTracerSpan(ctx, m.Headers, "queryMessageProcessor.processUserCreated")
	defer span.Finish()
	msg := &kafkaMessages.UserCreated{}
	if err := proto.Unmarshal(m.Value, msg); err != nil {
		s.log.WarnMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}
	p := msg.GetUser()
	// TODO: Write logic for Root and Active User fields below
	event := events.NewCreateUserEvent(p.GetID(), p.GetEmail(), p.GetUsername(), p.GetPassword(), p.GetRoot(), p.GetActive(), p.GetCreatedAt().AsTime(), p.GetUpdatedAt().AsTime())
	if err := s.v.StructCtx(ctx, event); err != nil {
		s.log.WarnMsg("validate", err)
		s.commitErrMessage(ctx, r, m)
		return
	}
	if err := retry.Do(func() error {
		return s.ps.Events.CreateUser.Handle(ctx, event)
	}, append(retryOptions, retry.Context(ctx))...); err != nil {
		s.log.WarnMsg("CreateUser.Handle", err)
		s.metrics.ErrorKafkaMessages.Inc()
		return
	}
	s.commitMessage(ctx, r, m)
}

func (s *queryMessageProcessor) processUserUpdated(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.metrics.UpdateUserKafkaMessages.Inc()
	ctx, span := tracing.StartKafkaConsumerTracerSpan(ctx, m.Headers, "queryMessageProcessor.processUserUpdated")
	defer span.Finish()
	msg := &kafkaMessages.UserUpdated{}
	if err := proto.Unmarshal(m.Value, msg); err != nil {
		s.log.WarnMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}
	p := msg.GetUser()
	event := events.NewUpdateUserEvent(p.GetID(), p.GetEmail(), p.GetUsername(), p.GetUpdatedAt().AsTime())
	if err := s.v.StructCtx(ctx, event); err != nil {
		s.log.WarnMsg("validate", err)
		s.commitErrMessage(ctx, r, m)
		return
	}
	if err := retry.Do(func() error {
		return s.ps.Events.UpdateUser.Handle(ctx, event)
	}, append(retryOptions, retry.Context(ctx))...); err != nil {
		s.log.WarnMsg("UpdateUser.Handle", err)
		s.metrics.ErrorKafkaMessages.Inc()
		return
	}
	s.commitMessage(ctx, r, m)
}

func (s *queryMessageProcessor) processUserDeleted(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.metrics.DeleteUserKafkaMessages.Inc()
	ctx, span := tracing.StartKafkaConsumerTracerSpan(ctx, m.Headers, "queryMessageProcessor.processUserDeleted")
	defer span.Finish()
	msg := &kafkaMessages.UserDeleted{}
	if err := proto.Unmarshal(m.Value, msg); err != nil {
		s.log.WarnMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}
	id, err := uuid.FromString(msg.GetID())
	if err != nil {
		s.log.WarnMsg("uuid.FromString", err)
		s.commitErrMessage(ctx, r, m)
		return
	}
	event := events.NewDeleteUserEvent(id)
	if err = s.v.StructCtx(ctx, event); err != nil {
		s.log.WarnMsg("validate", err)
		s.commitErrMessage(ctx, r, m)
		return
	}
	if err = retry.Do(func() error {
		return s.ps.Events.DeleteUser.Handle(ctx, event)
	}, append(retryOptions, retry.Context(ctx))...); err != nil {
		s.log.WarnMsg("DeleteUser.Handle", err)
		s.metrics.ErrorKafkaMessages.Inc()
		return
	}
	s.commitMessage(ctx, r, m)
}

func (s *queryMessageProcessor) commitMessage(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.metrics.SuccessKafkaMessages.Inc()
	s.log.KafkaLogCommittedMessage(m.Topic, m.Partition, m.Offset)
	if err := r.CommitMessages(ctx, m); err != nil {
		s.log.WarnMsg("commitMessage", err)
	}
}

func (s *queryMessageProcessor) logProcessMessage(m kafka.Message, workerID int) {
	s.log.KafkaProcessMessage(m.Topic, m.Partition, string(m.Value), workerID, m.Offset, m.Time)
}

func (s *queryMessageProcessor) commitErrMessage(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.metrics.ErrorKafkaMessages.Inc()
	s.log.KafkaLogCommittedMessage(m.Topic, m.Partition, m.Offset)
	if err := r.CommitMessages(ctx, m); err != nil {
		s.log.WarnMsg("commitMessage", err)
	}
}
