package container

import (
	"context"
	"errors"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
)

const (
	// InstanceTokenRefresher defines the global framework registration lookup token key targeted for active credentials refresh worker.
	InstanceTokenRefresher string = "InstanceTokenRefresher"
)

// WorkerContainer structures a lazy-loaded lifecycle dependency injection container managing background schedule components.
type WorkerContainer struct {
	*container.BaseLazyContainer // Generic framework-level baseline container orchestration handle
}

// Compile-time interface compliance verifications
var _ container.Container = (*WorkerContainer)(nil)
var _ container.LazyContainer = (*WorkerContainer)(nil)

// NewWorkerContainer acts as a factory constructor deploying structural configuration, logger, and orchestrator boundaries.
func NewWorkerContainer(
	orchestrator container.Orchestrator,
	log logger.Logger,
) *WorkerContainer {
	return &WorkerContainer{
		BaseLazyContainer: container.NewBaseLazyContainer(
			container.WithLazyName(WorkerContainerName),
			container.WithLazyOrchestrator(orchestrator),
			container.WithLazyLogger(log),
		),
	}
}

// Init triggers synchronous registration sequences linking required lazy provider callbacks inside the framework instance map graph.
func (wc *WorkerContainer) Init(initCtx context.Context) error {
	err := errors.Join(
		wc.RegisterProvider(InstanceTokenRefresher, wc.providerTokenRefresher),
	)
	if err != nil {
		return errs.NewContainerError(wc.GetName(), "container init: register providers failed", err)
	}

	return nil
}
