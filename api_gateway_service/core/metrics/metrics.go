package metrics

import (
	"fmt"
	"github.com/JECSand/go-cqrs-service-boilerplate/api_gateway_service/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type ApiGatewayMetrics struct {
	SuccessHttpRequests     prometheus.Counter
	ErrorHttpRequests       prometheus.Counter
	CreateUserHttpRequests  prometheus.Counter
	UpdateUserHttpRequests  prometheus.Counter
	DeleteUserHttpRequests  prometheus.Counter
	GetUserByIdHttpRequests prometheus.Counter
	SearchUserHttpRequests  prometheus.Counter
}

func NewApiGatewayMetrics(cfg *config.Config) *ApiGatewayMetrics {
	return &ApiGatewayMetrics{
		SuccessHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_http_requests_total", cfg.ServiceName),
			Help: "The total number of success http requests",
		}),
		ErrorHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_error_http_requests_total", cfg.ServiceName),
			Help: "The total number of error http requests",
		}),
		CreateUserHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_user_http_requests_total", cfg.ServiceName),
			Help: "The total number of create user http requests",
		}),
		UpdateUserHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_update_user_http_requests_total", cfg.ServiceName),
			Help: "The total number of update user http requests",
		}),
		DeleteUserHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_delete_user_http_requests_total", cfg.ServiceName),
			Help: "The total number of delete user http requests",
		}),
		GetUserByIdHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_get_user_by_id_http_requests_total", cfg.ServiceName),
			Help: "The total number of get user by id http requests",
		}),
		SearchUserHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_search_user_http_requests_total", cfg.ServiceName),
			Help: "The total number of search user http requests",
		}),
	}
}
