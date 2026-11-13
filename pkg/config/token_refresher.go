package config

import (
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/config"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
)

// TokenRefresherConfig содержит параметры конфигурации фонового воркера обновления JWT-токенов,
// включая настройки расписания, интервал повторов при ошибках и учетные данные для аутентификации.
type TokenRefresherConfig struct {
	Periodic              *config.SchedulerConfig       `mapstructure:"periodic" json:"periodic,omitempty" yaml:"periodic,omitempty"`
	ErrorScheduleInterval time.Duration                 `mapstructure:"error_schedule_interval" json:"error_schedule_interval,omitempty" yaml:"error_schedule_interval,omitempty"`
	Credentials           *config.BasicCredentialConfig `mapstructure:"credentials" json:"credentials,omitempty" yaml:"credentials,omitempty"`
}

// NewTokenRefresherConfig создает и инициализирует новый экземпляр конфигурации TokenRefresherConfig.
func NewTokenRefresherConfig(
	periodic *config.SchedulerConfig,
	credentials *config.BasicCredentialConfig,
	errorScheduleInterval time.Duration,
) *TokenRefresherConfig {
	return &TokenRefresherConfig{
		Periodic:              periodic,
		Credentials:           credentials,
		ErrorScheduleInterval: errorScheduleInterval,
	}
}

// NewDefaultTokenRefresherConfig создает экземпляр конфигурации TokenRefresherConfig со значениями по умолчанию.
func NewDefaultTokenRefresherConfig() *TokenRefresherConfig {
	return NewTokenRefresherConfig(
		config.NewDefaultSchedulerConfig(),
		config.NewDefaultBasicCredentialConfig(),
		DefaultTokenRefresherErrorScheduleInterval,
	)
}

// Validate проверяет наличие и корректность параметров расписания, учетных данных и интервала повторов при сбоях.
func (trc *TokenRefresherConfig) Validate() error {
	if utils.IsNil(trc.Periodic) {
		return errs.NewConfigValidateError("token refresher", "periodic", "is required", nil)
	}
	if err := trc.Periodic.Validate(); err != nil {
		return errs.NewConfigValidateError("token refresher", "periodic", "validation failed", err)
	}
	if utils.IsNil(trc.Credentials) {
		return errs.NewConfigValidateError("token refresher", "credentials", "is required", nil)
	}
	if err := trc.Credentials.Validate(); err != nil {
		return errs.NewConfigValidateError("token refresher", "credentials", "validation failed", err)
	}
	if trc.ErrorScheduleInterval <= 0 {
		return errs.NewConfigValidateError("token refresher", "errorScheduleInterval", "must be greater then zero", nil)
	}

	return nil
}
