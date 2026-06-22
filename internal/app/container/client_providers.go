package container

import (
	"fmt"

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

	/*
	   tlsConfInst, err := ....
	*/

	clientSender := azure.NewClientSender(
		confInst.LoginAttemptsSender.URL,
		logInst,
		azure.WithHostName(confInst.App.NodeName),
		azure.WithSASLPlain(confInst.LoginAttemptsSender.Username, confInst.LoginAttemptsSender.Password),
		// tls conf пропускаем.. потом добавим..
		azure.WithInsecureSkipVerify(confInst.LoginAttemptsSender.InsecureSkipVerify),
		azure.WithConnectTimeout(confInst.LoginAttemptsSender.ConnectTimeout),
		azure.WithWriteTimeout(confInst.LoginAttemptsSender.WriteTimeout),
	)

	return clientSender, nil
}
