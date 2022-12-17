package kafka

import (
	"context"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/config"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/metrics"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/services"
	"github.com/go-playground/validator"
	"github.com/segmentio/kafka-go"
	"sync"
)

const (
	PoolSize = 30
)

// TODO NEXT 12/17/22 START HERE FINISHING QUERY SERVICE KAFKA PACKAGE
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
		case s.cfg.KafkaTopics.ProductCreated.TopicName:
			s.processProductCreated(ctx, r, m)
		case s.cfg.KafkaTopics.ProductUpdated.TopicName:
			s.processProductUpdated(ctx, r, m)
		case s.cfg.KafkaTopics.ProductDeleted.TopicName:
			s.processProductDeleted(ctx, r, m)
		}
	}
}
