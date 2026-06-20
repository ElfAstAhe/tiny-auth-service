package container

import (
	"context"
	"errors"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/infra/pubsub"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
)

const (
	InstanceLoginAttemptsPublisher  string = "login-attempts-publisher"
	InstanceLoginAttemptsSubscriber string = "login-attempts-subscriber"
)

type InfraContainer struct {
	*container.BaseLazyContainer
}

var _ container.Container = (*InfraContainer)(nil)
var _ container.LazyContainer = (*InfraContainer)(nil)

func NewInfraContainer(orchestrator container.Orchestrator) *InfraContainer {
	return &InfraContainer{
		BaseLazyContainer: container.NewBaseLazyContainer(InfraContainerName, orchestrator),
	}
}

func (ic *InfraContainer) Init(ctx context.Context) error {
	err := errors.Join(
		ic.RegisterProvider(InstanceAuthAuditClient, ic.providerLoginAttemptsPublisher),
	)
	if err != nil {
		return errs.NewContainerError(ic.GetName(), "container init: register providers failed", err)
	}

	// setup publisher
	subscriber, err := container.GetInstance[pubsub.Observer[*dto.LoginAttemptEventDTO]](InstanceLoginAttemptsSubscriber)
	if err != nil {
		return errs.NewContainerError(ic.GetName(), "container init: get subscriber failed", err)
	}
	publisher, err := container.GetInstance[pubsub.Publisher[*dto.LoginAttemptEventDTO]](InstanceLoginAttemptsPublisher)
	if err != nil {
		return errs.NewContainerError(ic.GetName(), "container init: retrieve publisher failed", err)
	}

	publisher.Register(subscriber)

	return nil
}
