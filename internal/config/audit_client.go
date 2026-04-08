package config

import (
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
)

type AuditClientConfig struct {
	BaseURL            string        `mapstructure:"base_url" json:"base_url,omitempty" yaml:"base_url,omitempty"`
	Timeout            time.Duration `mapstructure:"timeout" json:"timeout,omitempty" yaml:"timeout,omitempty"`
	WorkerCount        int           `mapstructure:"worker_count" json:"worker_count,omitempty" yaml:"worker_count,omitempty"`
	DataCapacity       int           `mapstructure:"data_capacity" json:"data_capacity,omitempty" yaml:"data_capacity,omitempty"`
	CompleteProcessing bool          `mapstructure:"complete_processing" json:"complete_processing,omitempty" yaml:"complete_processing,omitempty"`
	ShutdownTimeout    time.Duration `mapstructure:"shutdown_timeout" json:"shutdown_timeout,omitempty" yaml:"shutdown_timeout,omitempty"`
}

func NewAuditClientConfig(
	baseURL string,
	timeout time.Duration,
	workerCount int,
	dataCapacity int,
	completeProcessing bool,
	shutdownTimeout time.Duration,
) *AuditClientConfig {
	return &AuditClientConfig{
		BaseURL:            baseURL,
		Timeout:            timeout,
		WorkerCount:        workerCount,
		DataCapacity:       dataCapacity,
		CompleteProcessing: completeProcessing,
		ShutdownTimeout:    shutdownTimeout,
	}
}

func NewDefaultAuditClientConfig() *AuditClientConfig {
	return NewAuditClientConfig(
		"",
		defaultAuditClientTimeout,
		defaultAuditClientWorkerCount,
		defaultAuditClientDataCapacity,
		defaultAuditClientCompleteProcessing,
		defaultAuditClientShutdownTimeout,
	)
}

func (acc *AuditClientConfig) Validate() error {
	if acc.BaseURL == "" {
		return errs.NewConfigValidateError("audit_client", "base_url", "must not be empty", nil)
	}
	if acc.Timeout <= 0 {
		return errs.NewConfigValidateError("audit_client", "timeout", "must be greater than zero", nil)
	}
	if acc.WorkerCount < 1 {
		return errs.NewConfigValidateError("audit_client", "worker_count", "must be greater than zero", nil)
	}
	if acc.DataCapacity < 1 {
		return errs.NewConfigValidateError("audit_client", "data_capacity", "must be greater than zero", nil)
	}
	if acc.ShutdownTimeout <= 0 {
		return errs.NewConfigValidateError("audit_client", "shutdown_timeout", "must be greater than zero", nil)
	}

	return nil
}
