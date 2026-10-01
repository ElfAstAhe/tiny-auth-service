package container

import (
	"context"
	"errors"

	_ "expvar"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/tiny-auth-service/internal/config"
)

// Orchestrator manages the root lifecycle and initialization sequence of the application's dependency graphs,
// explicitly routing parent configuration footprints and logging boundaries to individual sub-containers.
type Orchestrator struct {
	*container.BaseOrchestrator                // Generic framework-level root orchestration handle
	conf                        *config.Config // Structured baseline microservice configuration model
	logger                      logger.Logger  // Central infrastructure diagnostic logging engine instance
}

// Compile-time interface compliance verification
var _ container.Orchestrator = (*Orchestrator)(nil)

// NewOrchestrator acts as a factory constructor allocating a root orchestrator handle bound to core environment dependencies.
func NewOrchestrator(conf *config.Config, log logger.Logger) *Orchestrator {
	return &Orchestrator{
		BaseOrchestrator: container.NewBaseOrchestrator(log),
		conf:             conf,
		logger:           log,
	}
}

// Init extracts the base application container instance and explicitly mounts root configuration and logger blueprints into the dependency graph.
func (o *Orchestrator) Init(ctx context.Context) error {
	appCnt, err := o.GetContainer(AppContainerName)
	if err != nil {
		return errs.NewContainerError(OrchestratorName, "init failed", err)
	}
	err = errors.Join(
		appCnt.RegisterInstance(InstanceConfig, o.conf),
		appCnt.RegisterInstance(InstanceLogger, o.logger),
	)
	if err != nil {
		return errs.NewContainerError(OrchestratorName, "init failed", err)
	}

	return o.BaseOrchestrator.Init(ctx)
}
