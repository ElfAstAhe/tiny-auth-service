package container

import (
	"context"
	"errors"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/infra/pubsub"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/tiny-auth-service/internal/config"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
)

const (
	InstanceLoginAttemptsPublisher       string = "login-attempts-publisher"
	InstanceLoginAttemptsAMQPSubscriber  string = "login-attempts-amqp-subscriber"
	InstanceLoginAttemptsKafkaSubscriber string = "login-attempts-kafka-subscriber"
)

type InfraContainer struct {
	*container.BaseLazyContainer
}

var _ container.Container = (*InfraContainer)(nil)
var _ container.LazyContainer = (*InfraContainer)(nil)

func NewInfraContainer(
	orchestrator container.Orchestrator,
	log logger.Logger,
) *InfraContainer {
	return &InfraContainer{
		BaseLazyContainer: container.NewBaseLazyContainer(
			container.WithLazyName(InfraContainerName),
			container.WithLazyOrchestrator(orchestrator),
			container.WithLazyLogger(log),
		),
	}
}

//goland:noinspection GoUnusedParameter
func (ic *InfraContainer) Init(ctx context.Context) error {
	err := errors.Join(
		ic.RegisterProvider(InstanceLoginAttemptsPublisher, ic.providerLoginAttemptsEventDispatcher),
		ic.RegisterProvider(InstanceLoginAttemptsAMQPSubscriber, ic.providerLoginAttemptsAMQPObserver),
		ic.RegisterProvider(InstanceLoginAttemptsKafkaSubscriber, ic.providerLoginAttemptsKafkaObserver),
	)
	if err != nil {
		return errs.NewContainerError(ic.GetName(), "container init: register providers failed", err)
	}

	// setup publisher
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return errs.NewContainerError(ic.GetName(), "container init: get app config", err)
	}
	subscriber, err := ic.getLoginAttemptsSubscriber(confInst.LoginAttemptsSender.SenderKind)
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

func (ic *InfraContainer) getLoginAttemptsSubscriber(senderKind string) (pubsub.Observer[*dto.LoginAttemptEventDTO], error) {
	switch senderKind {
	case "amqp":
		return container.GetInstance[pubsub.Observer[*dto.LoginAttemptEventDTO]](InstanceLoginAttemptsAMQPSubscriber)
	case "kafka":
		return container.GetInstance[pubsub.Observer[*dto.LoginAttemptEventDTO]](InstanceLoginAttemptsKafkaSubscriber)
	default:
		return nil, errs.NewContainerError(ic.GetName(), fmt.Sprintf("unknown sender kind: %s", senderKind), nil)
	}
}
