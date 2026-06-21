package container

import (
	"context"
	"errors"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/transport/amqp"
)

const (
	InstanceAuthAuditClient  string = "audit-client"
	InstanceDataAuditClient  string = "data-audit-client"
	InstanceAMQPClientSender string = "amqp-client-sender"
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
		cc.RegisterProvider(InstanceAMQPClientSender, cc.providerAMQPClientSender),
	)
	if err != nil {
		return errs.NewContainerError(cc.GetName(), "container init: register providers failed", err)
	}

	return nil
}

func (cc *ClientContainer) Close(ctx context.Context) error {
	var closeErrs []error
	// retrieve all instances to close
	amqpClientSenderInst, err := container.GetInstance[amqp.ClientSender](InstanceAMQPClientSender)
	if err != nil {
		return errs.NewContainerError(cc.GetName(), "container close: get amqp client sender failed", err)
	}
	// close all instances
	err = amqpClientSenderInst.Close(ctx)
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
