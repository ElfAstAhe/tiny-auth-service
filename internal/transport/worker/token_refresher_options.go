package worker

import (
	"strings"
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/helper"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/pkg/transport/worker"
)

type TokenRefresherOption func(*TokenRefresherOptions)

type TokenRefresherOptions struct {
	*worker.BaseTokenRefresherOptions
	JWTHelper *helper.JWTHelper
	LoginUC   TokenRefresherLoginUseCase
	Username  string
	Password  string
}

func NewTokenRefreshOptions() *TokenRefresherOptions {
	return &TokenRefresherOptions{
		BaseTokenRefresherOptions: worker.NewBaseTokenRefresherOptions(),
	}
}

func (tro *TokenRefresherOptions) Validate() error {
	if err := tro.BaseTokenRefresherOptions.Validate(); err != nil {
		return errs.NewTlCommonError("Validate", "base options validation failed", err)
	}
	if utils.IsNil(tro.LoginUC) {
		return errs.NewTlCommonError("Validate", "login use case is required", nil)
	}
	if strings.TrimSpace(tro.Username) == "" {
		return errs.NewTlCommonError("Validate", "username is required", nil)
	}
	//if strings.TrimSpace(tro.Password) == "" {
	//    return errs.NewTlCommonError("Validate", "password is required", nil)
	//}
	if utils.IsNil(tro.JWTHelper) {
		return errs.NewTlCommonError("Validate", "jwtHelper is required", nil)
	}

	return nil
}

func (tro *TokenRefresherOptions) BaseOptions() []worker.BaseTokenRefresherOption {
	return []worker.BaseTokenRefresherOption{
		worker.WithTokenRefresherName(tro.Name),
		worker.WithTokenRefresherStartInterval(tro.StartInterval),
		worker.WithTokenRefresherStopTimeout(tro.StopTimeout),
		worker.WithTokenRefresherScheduleInterval(tro.ScheduleInterval),
		worker.WithTokenRefresherErrorScheduleInterval(tro.ErrorScheduleInterval),
		worker.WithTokenRefresherLogger(tro.Logger),
	}
}

// ====================================================================
// Fluent API методы для сборки опций
// ====================================================================

func WithTokenRefresherName(name string) TokenRefresherOption {
	return func(options *TokenRefresherOptions) {
		if utils.IsNil(options.BaseTokenRefresherOptions) {
			options.BaseTokenRefresherOptions = worker.NewBaseTokenRefresherOptions()
		}
		options.Name = name
	}
}

func WithTokenRefresherStartInterval(interval time.Duration) TokenRefresherOption {
	return func(options *TokenRefresherOptions) {
		if utils.IsNil(options.BaseTokenRefresherOptions) {
			options.BaseTokenRefresherOptions = worker.NewBaseTokenRefresherOptions()
		}
		options.StartInterval = interval
	}
}

func WithTokenRefresherStopTimeout(timeout time.Duration) TokenRefresherOption {
	return func(options *TokenRefresherOptions) {
		if utils.IsNil(options.BaseTokenRefresherOptions) {
			options.BaseTokenRefresherOptions = worker.NewBaseTokenRefresherOptions()
		}
		options.StopTimeout = timeout
	}
}

func WithTokenRefresherScheduleInterval(interval time.Duration) TokenRefresherOption {
	return func(options *TokenRefresherOptions) {
		if utils.IsNil(options.BaseTokenRefresherOptions) {
			options.BaseTokenRefresherOptions = worker.NewBaseTokenRefresherOptions()
		}
		options.ScheduleInterval = interval
	}
}

func WithTokenRefresherErrorScheduleInterval(interval time.Duration) TokenRefresherOption {
	return func(options *TokenRefresherOptions) {
		if utils.IsNil(options.BaseTokenRefresherOptions) {
			options.BaseTokenRefresherOptions = worker.NewBaseTokenRefresherOptions()
		}
		options.ErrorScheduleInterval = interval
	}
}

func WithTokenRefresherLogger(logger logger.Logger) TokenRefresherOption {
	return func(options *TokenRefresherOptions) {
		if utils.IsNil(options.BaseTokenRefresherOptions) {
			options.BaseTokenRefresherOptions = worker.NewBaseTokenRefresherOptions()
		}
		options.Logger = logger
	}
}

func WithTokenRefresherJWTHelper(jwtHelper *helper.JWTHelper) TokenRefresherOption {
	return func(options *TokenRefresherOptions) {
		options.JWTHelper = jwtHelper
	}
}

func WithTokenRefresherCredentials(username, password string) TokenRefresherOption {
	return func(options *TokenRefresherOptions) {
		options.Username = username
		options.Password = password
	}
}

func WithTokenRefresherLoginUseCase(loginUC TokenRefresherLoginUseCase) TokenRefresherOption {
	return func(options *TokenRefresherOptions) {
		options.LoginUC = loginUC
	}
}
