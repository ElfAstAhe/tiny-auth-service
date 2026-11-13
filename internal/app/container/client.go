package container

import (
	"context"
	"errors"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
)

const (
	// InstanceAMQPLoginAttemptsConnector structures the global framework session gateway managing connection allocations to AMQP clusters.
	InstanceAMQPLoginAttemptsConnector string = "login-attempts-amqp-connector"
	// InstanceAMQPLoginAttemptsConnectorConnOpts sets unique configuration metrics formatting low-level network connection parameters.
	InstanceAMQPLoginAttemptsConnectorConnOpts string = "login-attempts-amqp-connector-conn-opts"
	// InstanceAMQPLoginAttemptsConnectorSessOpts tracks identifier trajectories handling transactional session-level properties setup.
	InstanceAMQPLoginAttemptsConnectorSessOpts string = "login-attempts-amqp-connector-sess-opts"
	// InstanceAMQPLoginAttemptSender pins the structural delivery proxy routing authentication logs via AMQP messaging protocols.
	InstanceAMQPLoginAttemptSender string = "login-attempts-amqp-sender"
	// InstanceAMQPLoginAttemptSenderSenderOpts details specific parameters bounding wait thresholds or settlement modes under AMQP producers.
	InstanceAMQPLoginAttemptSenderSenderOpts string = "login-attempts-amqp-sender-opts"
	// InstanceKafkaLoginAttemptsSender maps configuration metrics initializing high-throughput async Kafka events tracking publishers.
	InstanceKafkaLoginAttemptsSender string = "login-attempts-kafka-sender"

	// InstanceAMQPDataAuditConnector amqp connector for data-audit
	InstanceAMQPDataAuditConnector string = "data-audit-amqp-connector"
	// InstanceAMQPDataAuditConnectorConnOpts amqp connector connection options for data audit
	InstanceAMQPDataAuditConnectorConnOpts string = "data-audit-amqp-connector-conn-opts"
	// InstanceAMQPDataAuditConnectorSessOpts amqp connector session options for data audit
	InstanceAMQPDataAuditConnectorSessOpts string = "data-audit-amqp-connector-sess-opts"
	// InstanceAMQPDataAuditSender amqp sender for data audit
	InstanceAMQPDataAuditSender string = "data-audit-amqp-sender"
	// InstanceAMQPDataAuditSenderSenderOpts amqp sender options for data audit
	InstanceAMQPDataAuditSenderSenderOpts string = "data-audit-amqp-sender-opts"
	// InstanceKafkaDataAuditSender kafka sender for data audit
	InstanceKafkaDataAuditSender string = "data-audit-kafka-sender"
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
		cc.RegisterProvider(InstanceAMQPLoginAttemptSender, cc.providerAMQPLoginAttemptSender),
		cc.RegisterProvider(InstanceAMQPLoginAttemptSenderSenderOpts, cc.providerAMQPLoginAttemptSenderSenderOpts),
		cc.RegisterProvider(InstanceAMQPLoginAttemptsConnector, cc.providerAMQPLoginAttemptsConnector),
		cc.RegisterProvider(InstanceAMQPLoginAttemptsConnectorConnOpts, cc.providerAMQPLoginAttemptsConnectorConnOpts),
		cc.RegisterProvider(InstanceAMQPLoginAttemptsConnectorSessOpts, cc.providerAMQPLoginAttemptsConnectorSessOpts),

		cc.RegisterProvider(InstanceKafkaLoginAttemptsSender, cc.providerKafkaLoginAttemptSender),

		cc.RegisterInstance(InstanceAMQPDataAuditSender, cc.providerAMQPDataAuditSender),
		cc.RegisterInstance(InstanceAMQPDataAuditSenderSenderOpts, cc.providerAMQPDataAuditSenderSenderOpts),
		cc.RegisterInstance(InstanceAMQPDataAuditConnector, cc.providerAMQPDataAuditConnector),
		cc.RegisterInstance(InstanceAMQPDataAuditConnectorConnOpts, cc.providerAMQPDataAuditConnectorConnOpts),
		cc.RegisterInstance(InstanceAMQPDataAuditConnectorSessOpts, cc.providerAMQPDataAuditConnectorSessOpts),

		cc.RegisterInstance(InstanceKafkaDataAuditSender, cc.providerKafkaDataAuditSender),
	)
	if err != nil {
		return errs.NewContainerError(cc.GetName(), "container init: register providers failed", err)
	}

	return nil
}
