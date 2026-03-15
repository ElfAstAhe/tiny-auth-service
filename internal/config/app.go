package config

import (
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
)

// AppConfig — метаданные сервиса
type AppConfig struct {
	Env          AppEnv `mapstructure:"env"` // dev, prod, test
	MaxListLimit int    `mapstructure:"max_list_limit"`
}

func NewAppConfig(env AppEnv, maxListLimit int) *AppConfig {
	return &AppConfig{
		Env:          env,
		MaxListLimit: maxListLimit,
	}
}

func NewDefaultAppConfig() *AppConfig {
	return NewAppConfig(defaultAppEnv, usecase.DefaultMaxLimit)
}

func (ac *AppConfig) Validate() error {
	if ac.Env == "" {
		return errs.NewConfigValidateError("app", "env", "must not be empty", nil)
	}

	if !ac.Env.Exists() {
		return errs.NewConfigValidateError("app", "env", "env value not match", nil)
	}

	if ac.MaxListLimit < 0 {
		return errs.NewConfigValidateError("app", "max_list_limit", "must be positive", nil)
	}

	return nil
}
