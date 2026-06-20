package config

import (
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/config"
)

type AMQPSenderConfig struct {
	*config.AMQPSenderConfig `mapstructure:",squash"`
}

func NewAMQPSenderConfig(
	url string,
	targetName string,
	username string,
	password string,
	insecureSkipVerify bool,
	connectTimeout time.Duration,
	writeTimeout time.Duration,
	notifyTimeout time.Duration,
	shutdownTimeout time.Duration,
) *AMQPSenderConfig {
	return &AMQPSenderConfig{
		AMQPSenderConfig: config.NewAMQPSenderConfig(
			url,
			targetName,
			username,
			password,
			insecureSkipVerify,
			connectTimeout,
			writeTimeout,
			notifyTimeout,
			shutdownTimeout,
		),
	}
}

func NewDefaultAMQPSenderConfig() *AMQPSenderConfig {
	return &AMQPSenderConfig{
		AMQPSenderConfig: config.NewDefaultAMQPSenderConfig(),
	}
}
