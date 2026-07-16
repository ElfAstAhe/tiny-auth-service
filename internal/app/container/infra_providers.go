package container

import (
	"github.com/Azure/go-amqp"
	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/infra/pubsub"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	libamqp "github.com/ElfAstAhe/go-service-template/pkg/transport/amqp"
	"github.com/ElfAstAhe/tiny-auth-service/internal/config"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	appamqp "github.com/ElfAstAhe/tiny-auth-service/internal/transport/amqp"
)

func (ic *InfraContainer) providerLoginAttemptsObserver() (any, error) {
	clientSender, err := container.GetInstance[libamqp.Sender[*amqp.SendOptions, *amqp.MessageHeader]](InstanceAMQPLoginAttemptSender)
	if err != nil {
		return nil, errs.NewContainerError(ic.GetName(), "container init: retrieve clientSender failed", err)
	}
	observer := appamqp.NewLoginAttemptObserver("login-attempts-amqp-observer", clientSender)

	return observer, nil
}

func (ic *InfraContainer) providerLoginAttemptsPublisher() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(ic.GetName(), "provider: retrieve instance failed", err)
	}
	logInst, err := container.GetInstance[logger.Logger](InstanceLogger)
	if err != nil {
		return nil, errs.NewContainerError(ic.GetName(), "provider: retrieve instance failed", err)
	}

	publisher := pubsub.NewEventDispatcher[*dto.LoginAttemptEventDTO]("login-attempts-publisher", confInst.LoginAttemptsSender.NotifyTimeout, logInst)

	return publisher, nil
}
