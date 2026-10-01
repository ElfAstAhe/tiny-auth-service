package container

import (
	"context"
	"errors"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/db"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/tiny-auth-service/internal/config"
	"github.com/ElfAstAhe/tiny-auth-service/internal/transport/worker"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
	"github.com/hellofresh/health-go/v5"
	healthPg "github.com/hellofresh/health-go/v5/checks/pgx5"
)

const (
	// InstanceHealthStatus defines the lookup token key targeted for registered liveness and readiness monitors.
	InstanceHealthStatus string = "health-status"
)

// ServiceContainer structures a lazy-loaded lifecycle dependency injection container managing cross-cutting diagnostic probes and cross-container runtime setups.
type ServiceContainer struct {
	*container.BaseLazyContainer // Generic framework-level baseline container orchestration handle
}

// Compile-time interface compliance verifications
var _ container.Container = (*ServiceContainer)(nil)
var _ container.LazyContainer = (*ServiceContainer)(nil)

// NewServiceContainer acts as a factory constructor deploying structural configuration, logger, and orchestrator boundaries.
func NewServiceContainer(
	orchestrator container.Orchestrator,
	log logger.Logger,
) *ServiceContainer {
	return &ServiceContainer{
		BaseLazyContainer: container.NewBaseLazyContainer(
			container.WithLazyName(ServiceContainerName),
			container.WithLazyOrchestrator(orchestrator),
			container.WithLazyLogger(log),
		),
	}
}

// Init triggers synchronous registration sequences linking required lazy diagnostic callbacks and links runtime dependencies.
func (sc *ServiceContainer) Init(initCtx context.Context) error {
	// add providers
	err := errors.Join(
		sc.RegisterProvider(InstanceHealthStatus, sc.providerHealthStatus),
	)
	if err != nil {
		return errs.NewContainerError(sc.GetName(), "container init: register providers failed", err)
	}

	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return errs.NewContainerError(sc.GetName(), "container init: retrieve instance failed", err)
	}
	// init health checks
	dbInst, err := container.GetInstance[db.DB](InstanceDB)
	if err != nil {
		return errs.NewContainerError(sc.GetName(), "container init: retrieve instance failed", err)
	}
	healthStatusInst, err := container.GetInstance[*health.Health](InstanceHealthStatus)
	if err != nil {
		return errs.NewContainerError(sc.GetName(), "container init: retrieve instance failed", err)
	}
	// Регистрируем Postgres. Либа сама будет делать Ping
	err = healthStatusInst.Register(health.Config{
		Name:      dbInst.GetDriver(),
		Timeout:   confInst.DB.ConnTimeout,
		SkipOnErr: false,
		Check: healthPg.New(healthPg.Config{
			DSN: confInst.DB.DSN,
		}),
	})
	if err != nil {
		return errs.NewContainerError(sc.GetName(), "container init: register health status check failed", err)
	}
	// setup token refresher
	simpleLoginUCInst, err := container.GetInstance[usecase.LoginSimpleUseCase](InstanceLoginSimpleUC)
	if err != nil {
		return errs.NewContainerError(sc.GetName(), "container init: retrieve instance failed", err)
	}
	tokenRefresher, err := container.GetInstance[*worker.TokenRefresher](InstanceTokenRefresher)
	if err != nil {
		return errs.NewContainerError(sc.GetName(), "container init: retrieve instance failed", err)
	}
	tokenRefresher.SetSimpleLoginUC(simpleLoginUCInst)

	return nil
}
