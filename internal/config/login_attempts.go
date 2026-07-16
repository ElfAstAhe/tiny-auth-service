package config

import (
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/config"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
)

type LoginAttemptsConfig struct {
	*config.AMQPSenderConfig `mapstructure:",squash"`
	NotifyTimeout            time.Duration `mapstructure:"notify_timeout" json:"notify_timeout,omitempty" yaml:"notify_timeout,omitempty"`
}

func NewLoginAttemptsConfig(
	senderConfig *config.AMQPSenderConfig,
	notifyTimeout time.Duration,
) *LoginAttemptsConfig {
	return &LoginAttemptsConfig{
		AMQPSenderConfig: senderConfig,
		NotifyTimeout:    notifyTimeout,
	}
}

func NewDefaultLoginAttemptsConfig() *LoginAttemptsConfig {
	return NewLoginAttemptsConfig(config.NewDefaultAMQPSenderConfig(), defaultLoginAttemptsNotifyTimeout)
}

func (lac *LoginAttemptsConfig) Validate() error {
	if err := lac.AMQPSenderConfig.Validate(); err != nil {
		return errs.NewConfigValidateError("login_attempts_sender", "AMQPSender", "validate failed", err)
	}
	if !(lac.NotifyTimeout > 0) {
		return errs.NewConfigValidateError("login_attempts_sender", "NotifyTimeout", "must be more than zero", nil)
	}

	return nil
}
