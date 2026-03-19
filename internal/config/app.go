package config

import (
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
)

// AppConfig — метаданные сервиса
type AppConfig struct {
	Env          AppEnv `mapstructure:"env"` // dev, prod, test
	MaxListLimit int    `mapstructure:"max_list_limit"`
	TokenIssuer  string `mapstructure:"token_issuer"`
	CipherKey    string `mapstructure:"cipher_key"`
}

func NewAppConfig(env AppEnv, maxListLimit int, tokenIssuer string, cipherKey string) *AppConfig {
	return &AppConfig{
		Env:          env,
		MaxListLimit: maxListLimit,
		TokenIssuer:  tokenIssuer,
		CipherKey:    cipherKey,
	}
}

func NewDefaultAppConfig() *AppConfig {
	return NewAppConfig(defaultAppEnv, defaultMaxListLimit, defaultTokenIssuer, "")
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

	return nil
}
