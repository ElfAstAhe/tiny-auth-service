package container

import (
	"context"
	"errors"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/infra/metrics"
	"github.com/ElfAstAhe/go-service-template/pkg/infra/pubsub"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/tiny-auth-service/internal/config"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
)

const (
	// InstanceLoginAttemptsPublisher defines the lookup token key mapping the internal aggregate pubsub publisher dispatch engine.
	InstanceLoginAttemptsPublisher string = "login-attempts-publisher"
	// InstanceLoginAttemptsAMQPSubscriber structures the registration identifier for the outbound AMQP event stream observer.
	InstanceLoginAttemptsAMQPSubscriber string = "login-attempts-amqp-subscriber"
	// InstanceLoginAttemptsKafkaSubscriber structures the registration identifier for the outbound Kafka event stream observer.
	InstanceLoginAttemptsKafkaSubscriber string = "login-attempts-kafka-subscriber"
)

// InfraContainer structures a lazy-loaded lifecycle dependency injection container managing cross-cutting telemetry metrics and broker notification routers.
type InfraContainer struct {
	*container.BaseLazyContainer // Generic framework-level baseline container orchestration handle
}

// Compile-time interface compliance verifications
var _ container.Container = (*InfraContainer)(nil)
var _ container.LazyContainer = (*InfraContainer)(nil)

// NewInfraContainer acts as a factory constructor deploying structural configuration, logger, and orchestrator boundaries.
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

// Init registers event-driven notification providers, initializes cross-cutting Prometheus metrics, and dynamically binds active infrastructure observers.
//
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

	// setup metrics
	metrics.InitHTTPMetrics()
	metrics.InitRepositoryMetrics()
	metrics.InitBrokerSenderMetrics()

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

// getLoginAttemptsSubscriber routes conditional evaluations matching explicit string tokens back to strongly-typed pubsub framework observers instances.
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
