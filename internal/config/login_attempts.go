package config

import (
	"fmt"
	"slices"
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/config"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
)

type LoginAttemptsConfig struct {
	SenderKind    string                    `mapstructure:"sender_kind" json:"sender_kind,omitempty" yaml:"sender_kind,omitempty"`
	NotifyTimeout time.Duration             `mapstructure:"notify_timeout" json:"notify_timeout,omitempty" yaml:"notify_timeout,omitempty"`
	AMQPConfig    *config.AMQPSenderConfig  `mapstructure:"amqp_config" json:"amqp_config,omitempty" yaml:"amqp_config"`
	KafkaConfig   *config.KafkaSenderConfig `mapstructure:"kafka_config" json:"kafka_config,omitempty" yaml:"kafka_config"`
}

func NewLoginAttemptsConfig(
	senderKind string,
	notifyTimeout time.Duration,
	amqpConfig *config.AMQPSenderConfig,
	kafkaConfig *config.KafkaSenderConfig,
) *LoginAttemptsConfig {
	return &LoginAttemptsConfig{
		SenderKind:    senderKind,
		NotifyTimeout: notifyTimeout,
		AMQPConfig:    amqpConfig,
		KafkaConfig:   kafkaConfig,
	}
}

func NewDefaultLoginAttemptsConfig() *LoginAttemptsConfig {
	return NewLoginAttemptsConfig(
		defaultLoginAttemptsSenderKind,
		defaultLoginAttemptsSenderNotifyTimeout,
		config.NewDefaultAMQPSenderConfig(),
		config.NewDefaultKafkaSenderConfig(),
	)
}

func (lac *LoginAttemptsConfig) Validate() error {
	if !slices.Contains([]string{"amqp", "kafka"}, lac.SenderKind) {
		return errs.NewConfigValidateError("login_attempts_sender", "SenderKind", fmt.Sprintf("unknown sender kind: %s", lac.SenderKind), nil)
	}
	if !(lac.NotifyTimeout > 0) {
		return errs.NewConfigValidateError("login_attempts_sender", "NotifyTimeout", "must be more than zero", nil)
	}

	switch lac.SenderKind {
	case "amqp":
		if utils.IsNil(lac.AMQPConfig) {
			return errs.NewConfigValidateError("login_attempts_sender", "AMQPConfig", "AMQPConfig is required", nil)
		}
		if err := lac.AMQPConfig.Validate(); err != nil {
			return errs.NewConfigValidateError("login_attempts_sender", "AMQPConfig", "validate failed", err)
		}
	case "kafka":
		if utils.IsNil(lac.KafkaConfig) {
			return errs.NewConfigValidateError("login_attempts_sender", "KafkaConfig", "KafkaConfig is required", nil)
		}
		if err := lac.KafkaConfig.Validate(); err != nil {
			return errs.NewConfigValidateError("login_attempts_sender", "KafkaConfig", "validate failed", err)
		}
	}

	return nil
}
