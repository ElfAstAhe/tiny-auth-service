package container

import (
	"context"
	"errors"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
)

const (
	InstanceAuthAuditClient string = "audit-client"
	InstanceDataAuditClient string = "data-audit-client"
)

type ClientContainer struct {
	*container.BaseLazyContainer
}

var _ container.Container = (*ClientContainer)(nil)
var _ container.LazyContainer = (*ClientContainer)(nil)

func NewClientContainer(orchestrator container.Orchestrator) *ClientContainer {
	return &ClientContainer{
		BaseLazyContainer: container.NewBaseLazyContainer(ClientContainerName, orchestrator),
	}
}

func (cc *ClientContainer) Init(ctx context.Context) error {
	err := errors.Join(
		cc.RegisterProvider(InstanceAuthAuditClient, cc.providerAuthAuditRestClient),
		cc.RegisterProvider(InstanceDataAuditClient, cc.providerDataAuditRestClient),
	)
	if err != nil {
		return errs.NewContainerError(cc.GetName(), "container init: register providers failed", err)
	}

	return nil
}
