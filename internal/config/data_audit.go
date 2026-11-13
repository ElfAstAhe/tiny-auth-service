package config

import (
	"fmt"
	"slices"
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/config"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
)

// DataAuditConfig encapsulates the configuration parameters required to initialize
// and validate outbound infrastructure event streaming publishers for data change tracking.
type DataAuditConfig struct {
	SenderKind    string                    `mapstructure:"sender_kind" json:"sender_kind,omitempty" yaml:"sender_kind,omitempty"`
	NotifyTimeout time.Duration             `mapstructure:"notify_timeout" json:"notify_timeout,omitempty" yaml:"notify_timeout,omitempty"`
	AMQPConfig    *config.AMQPSenderConfig  `mapstructure:"amqp_config" json:"amqp_config,omitempty" yaml:"amqp_config"`
	KafkaConfig   *config.KafkaSenderConfig `mapstructure:"kafka_config" json:"kafka_config,omitempty" yaml:"kafka_config"`
}

// NewDataAuditConfig acts as a factory constructor allocating data audit streaming properties.
func NewDataAuditConfig(
	senderKind string,
	notifyTimeout time.Duration,
	amqpConfig *config.AMQPSenderConfig,
	kafkaConfig *config.KafkaSenderConfig,
) *DataAuditConfig {
	return &DataAuditConfig{
		SenderKind:    senderKind,
		NotifyTimeout: notifyTimeout,
		AMQPConfig:    amqpConfig,
		KafkaConfig:   kafkaConfig,
	}
}

// NewDefaultDataAuditConfig returns a default initialized configuration schema blueprint.
func NewDefaultDataAuditConfig() *DataAuditConfig {
	return NewDataAuditConfig(
		defaultDataAuditSenderKind,
		defaultDataAuditSenderNotifyTimeout,
		config.NewDefaultAMQPSenderConfig(),
		config.NewDefaultKafkaSenderConfig(),
	)
}

// Validate executes strict fail-fast conditional routing validation logic across target broker properties before boot constraints.
func (dac *DataAuditConfig) Validate() error {
	if !slices.Contains([]string{"amqp", "kafka"}, dac.SenderKind) {
		return errs.NewConfigValidateError("data_audit_sender", "SenderKind", fmt.Sprintf("unknown sender kind: %s", dac.SenderKind), nil)
	}
	if dac.NotifyTimeout <= 0 {
		return errs.NewConfigValidateError("data_audit_sender", "NotifyTimeout", "must be more than zero", nil)
	}

	switch dac.SenderKind {
	case "amqp":
		if utils.IsNil(dac.AMQPConfig) {
			return errs.NewConfigValidateError("data_audit_sender", "AMQPConfig", "AMQPConfig is required", nil)
		}
		if err := dac.AMQPConfig.Validate(); err != nil {
			return errs.NewConfigValidateError("data_audit_sender", "AMQPConfig", "validate failed", err)
		}
	case "kafka":
		if utils.IsNil(dac.KafkaConfig) {
			return errs.NewConfigValidateError("data_audit_sender", "KafkaConfig", "KafkaConfig is required", nil)
		}
		if err := dac.KafkaConfig.Validate(); err != nil {
			return errs.NewConfigValidateError("data_audit_sender", "KafkaConfig", "validate failed", err)
		}
	}

	return nil
}
