package config

import (
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
)

// AppConfig — метаданные сервиса
type AppConfig struct {
	Env                AppEnv        `mapstructure:"env" json:"env,omitempty" yaml:"env,omitempty"` // dev, prod, test
	MaxListLimit       int           `mapstructure:"max_list_limit" json:"max_list_limit,omitempty" yaml:"max_list_limit,omitempty"`
	TokenIssuer        string        `mapstructure:"token_issuer" json:"token_issuer,omitempty" yaml:"token_issuer,omitempty"`
	CipherKey          string        `mapstructure:"cipher_key" json:"cipher_key,omitempty" yaml:"cipher_key,omitempty"`
	DefShutdownTimeout time.Duration `mapstructure:"def_shutdown_timeout" json:"def_shutdown_timeout,omitempty" yaml:"def_shutdown_timeout,omitempty"`
}

func NewAppConfig(
	env AppEnv,
	maxListLimit int,
	tokenIssuer string,
	cipherKey string,
	defShutdownTimeout time.Duration,
) *AppConfig {
	return &AppConfig{
		Env:                env,
		MaxListLimit:       maxListLimit,
		TokenIssuer:        tokenIssuer,
		CipherKey:          cipherKey,
		DefShutdownTimeout: defShutdownTimeout,
	}
}

func NewDefaultAppConfig() *AppConfig {
	return NewAppConfig(
		defaultAppEnv,
		defaultMaxListLimit,
		defaultTokenIssuer,
		"",
		defaultDefShutdownTimeout,
	)
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

	if ac.CipherKey == "" {
		return errs.NewConfigValidateError("app", "cipher-key", "must not be empty", nil)
	}

	if !(ac.DefShutdownTimeout > 0) {
		return errs.NewConfigValidateError("app", "def_shutdown_timeout", "must be positive", nil)
	}

	return nil
}
