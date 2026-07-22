package container

import (
	"context"
	"errors"

	"github.com/Azure/go-amqp"
	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	libamqp "github.com/ElfAstAhe/go-service-template/pkg/transport/amqp"
)

const (
	InstanceDataAuditClient                  string = "data-audit-client"
	InstanceAMQPConnector                    string = "amqp-connector"
	InstanceAMQPConnectorConnOpts            string = "amqp-connector-conn-opts"
	InstanceAMQPConnectorSessOpts            string = "amqp-connector-sess-opts"
	InstanceAMQPLoginAttemptSender           string = "amqp-login-attempt-sender"
	InstanceAMQPLoginAttemptSenderSenderOpts string = "amqp-client-sender-sender-opts"
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

//goland:noinspection DuplicatedCode
func (cc *ClientContainer) Init(ctx context.Context) error {
	err := errors.Join(
		//		cc.RegisterProvider(InstanceAuthAuditClient, cc.providerAuthAuditRestClient),
		cc.RegisterProvider(InstanceDataAuditClient, cc.providerDataAuditRestClient),
		cc.RegisterProvider(InstanceAMQPLoginAttemptSender, cc.providerAMQPLoginAttemptSender),
		cc.RegisterProvider(InstanceAMQPLoginAttemptSenderSenderOpts, cc.providerAMQPLoginAttemptSenderSenderOpts),
		cc.RegisterProvider(InstanceAMQPConnector, cc.providerAMQPConnector),
		cc.RegisterProvider(InstanceAMQPConnectorConnOpts, cc.providerAMQPConnectorConnOpts),
		cc.RegisterProvider(InstanceAMQPConnectorSessOpts, cc.providerAMQPConnectorSessOpts),
	)
	if err != nil {
		return errs.NewContainerError(cc.GetName(), "container init: register providers failed", err)
	}

	return nil
}

func (cc *ClientContainer) Close2(ctx context.Context) error {
	var closeErrs []error
	// retrieve all instances to close
	loginAttemptsSenderInst, err := container.GetInstance[libamqp.Sender[*amqp.SendOptions, *amqp.MessageHeader]](InstanceAMQPLoginAttemptSender)
	if err != nil {
		return errs.NewContainerError(cc.GetName(), "container close: retrieve instance failed", err)
	}
	// close all instances
	err = loginAttemptsSenderInst.Close(ctx)
	if err != nil {
		closeErrs = append(closeErrs, err)
	}

	// checks
	err = errors.Join(closeErrs...)
	if err != nil {
		return errs.NewContainerError(cc.GetName(), "container close: close fails", err)
	}

	return nil
}
