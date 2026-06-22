package config

import (
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/config"
)

type LoginAttemptsSenderConfig struct {
	*config.AMQPSenderConfig `mapstructure:",squash"`
}

func NewLoginAttemptsSenderConfig(
	url string,
	targetName string,
	username string,
	password string,
	insecureSkipVerify bool,
	connectTimeout time.Duration,
	writeTimeout time.Duration,
	notifyTimeout time.Duration,
	shutdownTimeout time.Duration,
) *LoginAttemptsSenderConfig {
	return &LoginAttemptsSenderConfig{
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

func NewDefaultLoginAttemptsSenderConfig() *LoginAttemptsSenderConfig {
	return &LoginAttemptsSenderConfig{
		AMQPSenderConfig: config.NewDefaultAMQPSenderConfig(),
	}
}
