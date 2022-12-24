package config

import (
	"flag"
	"fmt"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/constants"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/kafka"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/probes"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/tracing"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "API Gateway service config path")
}

// Config structures the configuration for the api gateway service
type Config struct {
	ServiceName string          `mapstructure:"serviceName"`
	Logger      *logging.Config `mapstructure:"logging"`
	KafkaTopics KafkaTopics     `mapstructure:"kafkaTopics"`
	Http        Http            `mapstructure:"http"`
	Grpc        Grpc            `mapstructure:"grpc"`
	Kafka       *kafka.Config   `mapstructure:"kafka"`
	Probes      probes.Config   `mapstructure:"probes"`
	Jaeger      *tracing.Config `mapstructure:"jaeger"`
}

type Http struct {
	Port                string   `mapstructure:"port"`
	Development         bool     `mapstructure:"development"`
	BasePath            string   `mapstructure:"basePath"`
	usersPath           string   `mapstructure:"usersPath"`
	DebugHeaders        bool     `mapstructure:"debugHeaders"`
	HttpClientDebug     bool     `mapstructure:"httpClientDebug"`
	DebugErrorsResponse bool     `mapstructure:"debugErrorsResponse"`
	IgnoreLogUrls       []string `mapstructure:"ignoreLogUrls"`
}

type Grpc struct {
	QueryServicePort string `mapstructure:"queryServicePort"`
}

type KafkaTopics struct {
	UserCreate kafka.TopicConfig `mapstructure:"userCreate"`
	UserUpdate kafka.TopicConfig `mapstructure:"userUpdate"`
	UserDelete kafka.TopicConfig `mapstructure:"userDelete"`
}

func InitConfig() (*Config, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv(constants.ConfigPath)
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			getwd, err := os.Getwd()
			if err != nil {
				return nil, errors.Wrap(err, "os.Getwd")
			}
			configPath = fmt.Sprintf("%s/api_gateway_service/config/config.yaml", getwd)
		}
	}

	cfg := &Config{}

	viper.SetConfigType(constants.Yaml)
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	httpPort := os.Getenv(constants.HttpPort)
	if httpPort != "" {
		cfg.Http.Port = httpPort
	}
	kafkaBrokers := os.Getenv(constants.KafkaBrokers)
	if kafkaBrokers != "" {
		cfg.Kafka.Brokers = []string{kafkaBrokers}
	}
	jaegerAddr := os.Getenv(constants.JaegerHostPort)
	if jaegerAddr != "" {
		cfg.Jaeger.HostPort = jaegerAddr
	}
	queryServicePort := os.Getenv(constants.QueryServicePort)
	if queryServicePort != "" {
		cfg.Grpc.QueryServicePort = queryServicePort
	}
	return cfg, nil
}
