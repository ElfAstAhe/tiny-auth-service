package container

import (
	"fmt"

	"github.com/Azure/go-amqp"
	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/go-service-template/pkg/transport/amqp/azure"
	libworker "github.com/ElfAstAhe/go-service-template/pkg/transport/worker"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/client/rest"
	"github.com/ElfAstAhe/tiny-auth-service/internal/config"
	"github.com/ElfAstAhe/tiny-auth-service/internal/transport/worker"
)

//goland:noinspection DuplicatedCode
func (cc *ClientContainer) providerAuthAuditRestClient() (any, error) {
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

	authAuditConf, err := rest.NewAuditClientConfig(
		confInst.AuthAuditClient.BaseURL,
		confInst.AuthAuditClient.Timeout,
		libworker.NewBasePoolConfig(
			confInst.AuthAuditClient.WorkerCount,
			confInst.AuthAuditClient.DataCapacity,
			confInst.AuthAuditClient.CompleteProcessing,
			confInst.AuthAuditClient.ShutdownTimeout,
		),
	)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), fmt.Sprintf("provider: create %s instance config failed", InstanceAuthAuditClient), err)
	}

	return rest.NewAuthAuditClient("auth-audit-client", authAuditConf, tokenRefresherInst, logInst), nil
}

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
func (cc *ClientContainer) providerAMQPClientSender() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}
	logInst, err := container.GetInstance[logger.Logger](InstanceLogger)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}
	connOptsInst, err := container.GetInstance[*amqp.ConnOptions](InstanceAMQPClientSenderConnOpts)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}
	sessOptsInst, err := container.GetInstance[*amqp.SessionOptions](InstanceAMQPClientSenderSessOpts)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}
	senderOptsInst, err := container.GetInstance[*amqp.SenderOptions](InstanceAMQPClientSenderOpts)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}

	clientSender, err := azure.NewClientSingleSender(
		azure.WithSenderURL(confInst.LoginAttemptsSender.URL),
		azure.WithSenderTargetName(confInst.LoginAttemptsSender.TargetName),
		azure.WithSenderLogger(logInst),
		azure.WithSenderConnOpts(connOptsInst),
		azure.WithSenderSessionOpts(sessOptsInst),
		azure.WithSenderOpts(senderOptsInst),
		azure.WithSenderConnectTimeout(confInst.LoginAttemptsSender.ConnectTimeout),
		azure.WithSenderShutdownTimeout(confInst.LoginAttemptsSender.ShutdownTimeout),
		azure.WithSenderPublishMaxTryAttempts(confInst.LoginAttemptsSender.PublishMaxTryAttempts),
		azure.WithSenderPublishBaseRetryDelay(confInst.LoginAttemptsSender.PublishBaseRetryDelay),
		azure.WithSenderPublishMaxRetryDelay(confInst.LoginAttemptsSender.PublishMaxRetryDelay),
	)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), fmt.Sprintf("provider: create %s instance failed", InstanceAMQPClientSender), err)
	}

	return clientSender, nil
}

func (cc *ClientContainer) providerAMQPClientSenderConnOpts() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}

	return &amqp.ConnOptions{
		HostName:     confInst.App.NodeName,
		WriteTimeout: confInst.LoginAttemptsSender.WriteTimeout,
		SASLType:     amqp.SASLTypePlain(confInst.LoginAttemptsSender.Username, confInst.LoginAttemptsSender.Password),
	}, nil
}

func (cc *ClientContainer) providerAMQPClientSenderSessOpts() (any, error) {
	return &amqp.SessionOptions{
		MaxLinks: 4,
	}, nil
}

func (cc *ClientContainer) providerAMQPClientSenderOpts() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}

	return &amqp.SenderOptions{
		Name:         confInst.App.NodeName,
		ExpiryPolicy: amqp.ExpiryPolicyNever,
	}, nil
}
