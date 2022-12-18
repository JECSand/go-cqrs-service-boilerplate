package config

import (
	"flag"
	"fmt"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/constants"
	kafkaClient "github.com/JECSand/go-cqrs-service-boilerplate/pkg/kafka"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/postgres"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/probes"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/tracing"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "Command service config path")
}

type Config struct {
	ServiceName string              `mapstructure:"serviceName"`
	Logger      *logging.Config     `mapstructure:"logger"`
	KafkaTopics KafkaTopics         `mapstructure:"kafkaTopics"`
	GRPC        GRPC                `mapstructure:"grpc"`
	Postgresql  *postgres.Config    `mapstructure:"postgres"`
	Kafka       *kafkaClient.Config `mapstructure:"kafka"`
	Probes      probes.Config       `mapstructure:"probes"`
	Jaeger      *tracing.Config     `mapstructure:"jaeger"`
}

type GRPC struct {
	Port        string `mapstructure:"port"`
	Development bool   `mapstructure:"development"`
}

type KafkaTopics struct {
	UserCreate  kafkaClient.TopicConfig `mapstructure:"userCreate"`
	UserCreated kafkaClient.TopicConfig `mapstructure:"userCreated"`
	UserUpdate  kafkaClient.TopicConfig `mapstructure:"userUpdate"`
	UserUpdated kafkaClient.TopicConfig `mapstructure:"userUpdated"`
	UserDelete  kafkaClient.TopicConfig `mapstructure:"userDelete"`
	UserDeleted kafkaClient.TopicConfig `mapstructure:"userDeleted"`
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
			configPath = fmt.Sprintf("%s/command_service/config/config.yaml", getwd)
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
	grpcPort := os.Getenv(constants.GrpcPort)
	if grpcPort != "" {
		cfg.GRPC.Port = grpcPort
	}
	postgresHost := os.Getenv(constants.PostgresqlHost)
	if postgresHost != "" {
		cfg.Postgresql.Host = postgresHost
	}
	postgresPort := os.Getenv(constants.PostgresqlPort)
	if postgresPort != "" {
		cfg.Postgresql.Port = postgresPort
	}
	jaegerAddr := os.Getenv(constants.JaegerHostPort)
	if jaegerAddr != "" {
		cfg.Jaeger.HostPort = jaegerAddr
	}
	kafkaBrokers := os.Getenv(constants.KafkaBrokers)
	if kafkaBrokers != "" {
		cfg.Kafka.Brokers = []string{kafkaBrokers}
	}
	return cfg, nil
}
