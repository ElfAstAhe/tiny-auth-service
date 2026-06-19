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
	address string,
	targetName string,
	username string,
	password string,
	secure bool,
	insecureSkipVerify bool,
	connectTimeout time.Duration,
	writeTimeout time.Duration,
	notifyTimeout time.Duration,
	shutdownTimeout time.Duration,
) *AMQPSenderConfig {
	return &AMQPSenderConfig{
		AMQPSenderConfig: config.NewAMQPSenderConfig(
			url,
			address,
			targetName,
			username,
			password,
			secure,
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
