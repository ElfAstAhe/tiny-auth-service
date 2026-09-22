package container

import (
	"context"
	"errors"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
)

const (
	InstanceDataAuditClient                  string = "data-audit-client"
	InstanceAMQPConnector                    string = "amqp-connector"
	InstanceAMQPConnectorConnOpts            string = "amqp-connector-conn-opts"
	InstanceAMQPConnectorSessOpts            string = "amqp-connector-sess-opts"
	InstanceAMQPLoginAttemptSender           string = "amqp-login-attempt-sender"
	InstanceAMQPLoginAttemptSenderSenderOpts string = "amqp-client-sender-sender-opts"
	InstanceKafkaLoginAttemptsSender         string = "kafka-login-attempts-sender"
)

type ClientContainer struct {
	*container.BaseLazyContainer
}

var _ container.Container = (*ClientContainer)(nil)
var _ container.LazyContainer = (*ClientContainer)(nil)

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
