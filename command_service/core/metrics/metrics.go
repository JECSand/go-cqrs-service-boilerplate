package metrics

import (
	"fmt"
	"github.com/JECSand/go-cqrs-service-boilerplate/command_service/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type CommandServiceMetrics struct {
	SuccessGrpcRequests     prometheus.Counter
	ErrorGrpcRequests       prometheus.Counter
	CreateUserGrpcRequests  prometheus.Counter
	UpdateUserGrpcRequests  prometheus.Counter
	DeleteUserGrpcRequests  prometheus.Counter
	GetUserByIdGrpcRequests prometheus.Counter
	SearchUserGrpcRequests  prometheus.Counter
	SuccessKafkaMessages    prometheus.Counter
	ErrorKafkaMessages      prometheus.Counter
	CreateUserKafkaMessages prometheus.Counter
	UpdateUserKafkaMessages prometheus.Counter
	DeleteUserKafkaMessages prometheus.Counter
}

func NewCommandServiceMetrics(cfg *config.Config) *CommandServiceMetrics {
	return &CommandServiceMetrics{
		SuccessGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of success grpc requests",
		}),
		ErrorGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_error_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of error grpc requests",
		}),
		CreateUserGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_user_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of create user grpc requests",
		}),
		UpdateUserGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_update_user_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of update user grpc requests",
		}),
		DeleteUserGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_delete_user_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of delete user grpc requests",
		}),
		GetUserByIdGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_get_user_by_id_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of get user by id grpc requests",
		}),
		SearchUserGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_search_user_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of search user grpc requests",
		}),
		CreateUserKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_user_kafka_messages_total", cfg.ServiceName),
			Help: "The total number of create user kafka messages",
		}),
		UpdateUserKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_update_user_kafka_messages_total", cfg.ServiceName),
			Help: "The total number of update user kafka messages",
		}),
		DeleteUserKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_delete_user_kafka_messages_total", cfg.ServiceName),
			Help: "The total number of delete user kafka messages",
		}),
		SuccessKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_kafka_processed_messages_total", cfg.ServiceName),
			Help: "The total number of success kafka processed messages",
		}),
		ErrorKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_error_kafka_processed_messages_total", cfg.ServiceName),
			Help: "The total number of error kafka processed messages",
		}),
	}
}
