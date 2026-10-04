package worker

import (
	"strings"
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/go-service-template/pkg/transport/worker"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
)

const (
	// DefaultTokenRefresherErrorScheduleInterval время повторного запроса токена при ошибке
	DefaultTokenRefresherErrorScheduleInterval time.Duration = 5 * time.Second
)

// BaseTokenRefresherOption defines a functional option signature for configuring token refresher background scheduler structures.
type BaseTokenRefresherOption func(*BaseTokenRefresherOptions)

// BaseTokenRefresherOptions опции scheduler воркера обновления токена
type BaseTokenRefresherOptions struct {
	*worker.BaseSchedulerOptions
	ErrorScheduleInterval time.Duration
	TokenRefreshAction    TokenRefreshAction
}

// NewBaseTokenRefresherOptions acts as a factory constructor allocating a nested baseline schedule profile configuration blueprints pointer.
func NewBaseTokenRefresherOptions() *BaseTokenRefresherOptions {
	return &BaseTokenRefresherOptions{
		BaseSchedulerOptions:  worker.NewBaseSchedulerOptions(),
		ErrorScheduleInterval: DefaultTokenRefresherErrorScheduleInterval,
	}
}

// Validate выполняет строгую проверку параметров конфигурации перед аллокацией памяти кэша,
// защищая приложение от логических ошибок планировщика и паник из-за nil-зависимостей.
func (tro *BaseTokenRefresherOptions) Validate() error {
	if strings.TrimSpace(tro.Name) == "" {
		return errs.NewTlCommonError("Validate", "name is required", nil)
	}
	if tro.StartInterval <= 0 {
		return errs.NewTlCommonError("Validate", "start interval is required", nil)
	}
	if tro.ScheduleInterval <= 0 {
		return errs.NewTlCommonError("Validate", "schedule interval is required", nil)
	}
	if tro.StopTimeout <= 0 {
		return errs.NewTlCommonError("Validate", "stop timeout is required", nil)
	}
	if utils.IsNil(tro.TimerDispatcher) {
		return errs.NewTlCommonError("Validate", "timer dispatcher is required", nil)
	}
	if utils.IsNil(tro.Logger) {
		return errs.NewTlCommonError("Validate", "logger is required", nil)
	}
	if tro.ErrorScheduleInterval <= 0 {
		return errs.NewTlCommonError("Validate", "error schedule interval is required", nil)
	}
	if utils.IsNil(tro.TokenRefreshAction) {
		return errs.NewTlCommonError("Validate", "token refresh action is required", nil)
	}

	return nil
}

// ====================================================================
// Fluent API методы для сборки опций
// ====================================================================

// WithTokenRefreshedName sets a flat diagnostic tag string pointer identifying the instance inside operational telemetry records blocks.
func WithTokenRefreshedName(name string) BaseTokenRefresherOption {
	return func(options *BaseTokenRefresherOptions) {
		options.Name = name
	}
}

// WithTokenRefresherStartInterval настраивает первичную задержку (холодное смещение) перед самым первым тиком таймера
func WithTokenRefresherStartInterval(interval time.Duration) BaseTokenRefresherOption {
	return func(options *BaseTokenRefresherOptions) {
		options.StartInterval = interval
	}
}

// WithTokenRefresherScheduleInterval настраивает фиксированный интервал периодического повторения задач (период)
func WithTokenRefresherScheduleInterval(interval time.Duration) BaseTokenRefresherOption {
	return func(options *BaseTokenRefresherOptions) {
		options.ScheduleInterval = interval
	}
}

// WithTokenRefresherStopTimeout настраивает временной лимит (таймаут) на мягкое завершение активной итерации обработчика
func WithTokenRefresherStopTimeout(timeout time.Duration) BaseTokenRefresherOption {
	return func(options *BaseTokenRefresherOptions) {
		options.StopTimeout = timeout
	}
}

// WithTokenRefresherLogger настраивает logger
func WithTokenRefresherLogger(logger logger.Logger) BaseTokenRefresherOption {
	return func(options *BaseTokenRefresherOptions) {
		options.Logger = logger
	}
}

// WithTokenRefresherErrorScheduleInterval bounds specific temporal intervals to retry security tokens generation sweeps on downstream dependency errors.
func WithTokenRefresherErrorScheduleInterval(errorScheduleInterval time.Duration) BaseTokenRefresherOption {
	return func(options *BaseTokenRefresherOptions) {
		options.ErrorScheduleInterval = errorScheduleInterval
	}
}

// WithTokenRefresherTokenRefreshAction registers a dedicated executable closure payload responsible for invoking core credential retrieval calls.
func WithTokenRefresherTokenRefreshAction(tokenRefreshAction TokenRefreshAction) BaseTokenRefresherOption {
	return func(options *BaseTokenRefresherOptions) {
		options.TokenRefreshAction = tokenRefreshAction
	}
}
