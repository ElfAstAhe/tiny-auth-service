package container

import (
	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/infra/pubsub"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/go-service-template/pkg/transport/broker"
	libamqp "github.com/ElfAstAhe/go-service-template/pkg/transport/broker/amqp"
	"github.com/ElfAstAhe/tiny-auth-service/internal/config"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	appamqp "github.com/ElfAstAhe/tiny-auth-service/internal/transport/amqp"
	appkafka "github.com/ElfAstAhe/tiny-auth-service/internal/transport/kafka"
)

func (ic *InfraContainer) providerLoginAttemptsAMQPObserver() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(ic.GetName(), "provider: retrieve instance failed", err)
	}
	clientSender, err := container.GetInstance[libamqp.AMQPSender](InstanceAMQPLoginAttemptSender)
	if err != nil {
		return nil, errs.NewContainerError(ic.GetName(), "container init: retrieve clientSender failed", err)
	}
	observer := appamqp.NewLoginAttemptObserver("login-attempts-amqp-observer", clientSender, confInst.LoginAttemptsSender.SenderKind)

	return observer, nil
}

func (ic *InfraContainer) providerLoginAttemptsKafkaObserver() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(ic.GetName(), "provider: retrieve instance failed", err)
	}
	clientSender, err := container.GetInstance[broker.Sender](InstanceKafkaLoginAttemptsSender)
	if err != nil {
		return nil, errs.NewContainerError(ic.GetName(), "container init: retrieve clientSender failed", err)
	}
	observer := appkafka.NewLoginAttemptObserver("login-attempts-kafka-observer", clientSender, confInst.LoginAttemptsSender.SenderKind)

	return observer, nil
}

func (ic *InfraContainer) providerLoginAttemptsEventDispatcher() (any, error) {
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
