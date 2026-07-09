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
	InstanceAuthAuditClient          string = "audit-client"
	InstanceDataAuditClient          string = "data-audit-client"
	InstanceAMQPClientSender         string = "amqp-client-sender"
	InstanceAMQPClientSenderConnOpts string = "amqp-client-sender-conn-opts"
	InstanceAMQPClientSenderSessOpts string = "amqp-client-sender-sess-opts"
	InstanceAMQPClientSenderOpts     string = "amqp-client-sender-opts"
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
		cc.RegisterProvider(InstanceAMQPClientSender, cc.providerAMQPClientSender),
		cc.RegisterProvider(InstanceAMQPClientSenderConnOpts, cc.providerAMQPClientSenderConnOpts),
		cc.RegisterProvider(InstanceAMQPClientSenderSessOpts, cc.providerAMQPClientSenderSessOpts),
		cc.RegisterProvider(InstanceAMQPClientSenderOpts, cc.providerAMQPClientSenderOpts),
	)
	if err != nil {
		return errs.NewContainerError(cc.GetName(), "container init: register providers failed", err)
	}

	return nil
}

func (cc *ClientContainer) Close(ctx context.Context) error {
	var closeErrs []error
	// retrieve all instances to close
	loginAttemptsSenderInst, err := container.GetInstance[libamqp.ClientSingleSender[*amqp.SendOptions, *amqp.MessageHeader]](InstanceAMQPClientSender)
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
