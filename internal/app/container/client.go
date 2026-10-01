package container

import (
	"context"
	"errors"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
)

const (
	// InstanceDataAuditClient defines the lookup token key targeted for primary public REST data audit log client instance.
	InstanceDataAuditClient string = "data-audit-client"
	// InstanceAMQPConnector structures the global framework session gateway managing connection allocations to AMQP clusters.
	InstanceAMQPConnector string = "amqp-connector"
	// InstanceAMQPConnectorConnOpts sets unique configuration metrics formatting low-level network connection parameters.
	InstanceAMQPConnectorConnOpts string = "amqp-connector-conn-opts"
	// InstanceAMQPConnectorSessOpts tracks identifier trajectories handling transactional session-level properties setup.
	InstanceAMQPConnectorSessOpts string = "amqp-connector-sess-opts"
	// InstanceAMQPLoginAttemptSender pins the structural delivery proxy routing authentication logs via AMQP messaging protocols.
	InstanceAMQPLoginAttemptSender string = "amqp-login-attempt-sender"
	// InstanceAMQPLoginAttemptSenderSenderOpts details specific parameters bounding wait thresholds or settlement modes under AMQP producers.
	InstanceAMQPLoginAttemptSenderSenderOpts string = "amqp-client-sender-sender-opts"
	// InstanceKafkaLoginAttemptsSender maps configuration metrics initializing high-throughput async Kafka events tracking publishers.
	InstanceKafkaLoginAttemptsSender string = "kafka-login-attempts-sender"
)

// ClientContainer structures a lazy-loaded lifecycle dependency injection container managing outbound system integrations clients and configurations options.
type ClientContainer struct {
	*container.BaseLazyContainer // Generic framework-level baseline container orchestration handle
}

// Compile-time interface compliance verifications
var _ container.Container = (*ClientContainer)(nil)
var _ container.LazyContainer = (*ClientContainer)(nil)

// NewClientContainer acts as a factory constructor deploying structural configuration, logger, and orchestrator boundaries.
func NewClientContainer(
	orchestrator container.Orchestrator,
	log logger.Logger,
) *ClientContainer {
	return &ClientContainer{
		BaseLazyContainer: container.NewBaseLazyContainer(
			container.WithLazyName(ClientContainerName),
			container.WithLazyOrchestrator(orchestrator),
			container.WithLazyLogger(log),
		),
	}
}

// Init registers integration REST handlers, AMQP connectors setups, and Kafka event producers providers within the framework graph.
//
//goland:noinspection DuplicatedCode,GoUnusedParameter
func (cc *ClientContainer) Init(ctx context.Context) error {
	err := errors.Join(
		cc.RegisterProvider(InstanceDataAuditClient, cc.providerDataAuditRestClient),
		cc.RegisterProvider(InstanceAMQPLoginAttemptSender, cc.providerAMQPLoginAttemptSender),
		cc.RegisterProvider(InstanceAMQPLoginAttemptSenderSenderOpts, cc.providerAMQPLoginAttemptSenderSenderOpts),
		cc.RegisterProvider(InstanceAMQPConnector, cc.providerAMQPConnector),
		cc.RegisterProvider(InstanceAMQPConnectorConnOpts, cc.providerAMQPConnectorConnOpts),
		cc.RegisterProvider(InstanceAMQPConnectorSessOpts, cc.providerAMQPConnectorSessOpts),
		cc.RegisterProvider(InstanceKafkaLoginAttemptsSender, cc.providerKafkaLoginAttemptSender),
	)
	if err != nil {
		return errs.NewContainerError(cc.GetName(), "container init: register providers failed", err)
	}

	return nil
}
