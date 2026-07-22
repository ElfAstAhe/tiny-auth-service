package container

import (
	"fmt"

	"github.com/Azure/go-amqp"
	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	libamqp "github.com/ElfAstAhe/go-service-template/pkg/transport/amqp"
	"github.com/ElfAstAhe/go-service-template/pkg/transport/amqp/azure"
	libworker "github.com/ElfAstAhe/go-service-template/pkg/transport/worker"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/client/rest"
	"github.com/ElfAstAhe/tiny-auth-service/internal/config"
	"github.com/ElfAstAhe/tiny-auth-service/internal/transport/worker"
)

//goland:noinspection DuplicatedCode
func (cc *ClientContainer) providerDataAuditRestClient() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}
	logInst, err := container.GetInstance[logger.Logger](InstanceLogger)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}
	tokenRefresherInst, err := container.GetInstance[*worker.TokenRefresher](InstanceTokenRefresher)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}
	dataAuditConf, err := rest.NewAuditClientConfig(
		confInst.DataAuditClient.BaseURL,
		confInst.DataAuditClient.Timeout,
		libworker.NewBasePoolConfig(
			confInst.DataAuditClient.WorkerCount,
			confInst.DataAuditClient.DataCapacity,
			confInst.DataAuditClient.CompleteProcessing,
			confInst.DataAuditClient.ShutdownTimeout,
		),
	)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), fmt.Sprintf("provider: create %s instance config failed", InstanceDataAuditClient), err)
	}

	return rest.NewDataAuditClient("data-audit-client", dataAuditConf, tokenRefresherInst, logInst), nil
}

//goland:noinspection DuplicatedCode
func (cc *ClientContainer) providerAMQPLoginAttemptSender() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}
	logInst, err := container.GetInstance[logger.Logger](InstanceLogger)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}
	connectorInst, err := container.GetInstance[libamqp.Connector[*amqp.Session]](InstanceAMQPConnector)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}
	senderOptsInst, err := container.GetInstance[*amqp.SenderOptions](InstanceAMQPLoginAttemptSenderSenderOpts)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}

	sender, err := azure.NewSender(
		azure.WithSenderConnector(connectorInst),
		azure.WithSenderTargetName(confInst.LoginAttemptsSender.TargetName),
		azure.WithSenderLogger(logInst),
		azure.WithSenderOpts(senderOptsInst),
		azure.WithSenderConnectTimeout(confInst.LoginAttemptsSender.ConnectTimeout),
		azure.WithSenderShutdownTimeout(confInst.LoginAttemptsSender.ShutdownTimeout),
		azure.WithSenderPublishMaxTryAttempts(confInst.LoginAttemptsSender.PublishMaxTryAttempts),
		azure.WithSenderPublishBaseRetryDelay(confInst.LoginAttemptsSender.PublishBaseRetryDelay),
		azure.WithSenderPublishMaxRetryDelay(confInst.LoginAttemptsSender.PublishMaxRetryDelay),
	)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), fmt.Sprintf("provider: create %s instance failed", InstanceAMQPLoginAttemptSender), err)
	}

	return sender, nil
}

//goland:noinspection DuplicatedCode
func (cc *ClientContainer) providerAMQPConnector() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}
	logInst, err := container.GetInstance[logger.Logger](InstanceLogger)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}
	connOpts, err := container.GetInstance[*amqp.ConnOptions](InstanceAMQPConnectorConnOpts)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}
	sessOpts, err := container.GetInstance[*amqp.SessionOptions](InstanceAMQPConnectorSessOpts)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}

	connectorInst, err := azure.NewConnector(
		azure.WithConnectorURL(confInst.AMQPConnector.URL),
		azure.WithConnectorConnectTimeout(confInst.AMQPConnector.ConnectTimeout),
		azure.WithConnectorShutdownTimeout(confInst.AMQPConnector.ShutdownTimeout),
		azure.WithConnectorConnOpts(connOpts),
		azure.WithConnectorSessionOpts(sessOpts),
		azure.WithConnectorLogger(logInst),
	)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), fmt.Sprintf("provider: create %s instance failed", InstanceAMQPConnector), err)
	}

	return connectorInst, nil
}

func (cc *ClientContainer) providerAMQPConnectorConnOpts() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}

	return &amqp.ConnOptions{
		HostName:     confInst.App.NodeName,
		IdleTimeout:  confInst.AMQPConnector.IdleTimeout,
		WriteTimeout: confInst.AMQPConnector.WriteTimeout,
		SASLType:     amqp.SASLTypePlain(confInst.AMQPConnector.Username, confInst.AMQPConnector.Password),
	}, nil
}

func (cc *ClientContainer) providerAMQPConnectorSessOpts() (any, error) {
	return &amqp.SessionOptions{
		MaxLinks: 4,
	}, nil
}

func (cc *ClientContainer) providerAMQPLoginAttemptSenderSenderOpts() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}

	return &amqp.SenderOptions{
		Name:         confInst.App.NodeName,
		ExpiryPolicy: amqp.ExpiryPolicyNever,
		Durability:   amqp.DurabilityUnsettledState,
	}, nil
}
