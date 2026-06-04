package container

import (
	"context"
	"errors"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
)

const (
	InstanceTokenRefresher string = "InstanceTokenRefresher"
)

type WorkerContainer struct {
	*container.BaseLazyContainer
}

var _ container.Container = (*WorkerContainer)(nil)
var _ container.LazyContainer = (*WorkerContainer)(nil)

func NewWorkerContainer(orchestrator container.Orchestrator) *WorkerContainer {
	return &WorkerContainer{
		BaseLazyContainer: container.NewBaseLazyContainer(WorkerContainerName, orchestrator),
	}
}

func (wc *WorkerContainer) Init(initCtx context.Context) error {
	err := errors.Join(
		wc.RegisterProvider(InstanceTokenRefresher, wc.providerTokenRefresher),
	)
	if err != nil {
		return errs.NewContainerError(wc.GetName(), "container init: register providers failed", err)
	}

	return nil
}
