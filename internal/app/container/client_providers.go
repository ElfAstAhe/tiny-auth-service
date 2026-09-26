package container

import (
	"fmt"

	"github.com/Azure/go-amqp"
	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/go-service-template/pkg/transport/broker"
	libamqp "github.com/ElfAstAhe/go-service-template/pkg/transport/broker/amqp"
	libkafka "github.com/ElfAstAhe/go-service-template/pkg/transport/broker/kafka"
	libworker "github.com/ElfAstAhe/go-service-template/pkg/transport/worker"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/client/rest"
	"github.com/ElfAstAhe/tiny-auth-service/internal/config"
	"github.com/ElfAstAhe/tiny-auth-service/internal/transport/worker"
	"github.com/segmentio/kafka-go"
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
	if utils.IsNil(confInst.LoginAttemptsSender.AMQPConfig) {
		return nil, errs.NewContainerError(cc.GetName(), "provider: AMQP config absent", nil)
	}
	logInst, err := container.GetInstance[logger.Logger](InstanceLogger)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}
	connectorInst, err := container.GetInstance[broker.Connector[*amqp.Session]](InstanceAMQPConnector)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}
	senderOptsInst, err := container.GetInstance[*amqp.SenderOptions](InstanceAMQPLoginAttemptSenderSenderOpts)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}

	sender, err := libamqp.NewSender(
		libamqp.WithSenderConnector(connectorInst),
		libamqp.WithSenderTargetName(confInst.LoginAttemptsSender.AMQPConfig.TargetName),
		libamqp.WithSenderLogger(logInst),
		libamqp.WithSenderOpts(senderOptsInst),
		libamqp.WithSenderConnectTimeout(confInst.LoginAttemptsSender.AMQPConfig.ConnectTimeout),
		libamqp.WithSenderShutdownTimeout(confInst.LoginAttemptsSender.AMQPConfig.ShutdownTimeout),
		libamqp.WithSenderPublishMaxTryAttempts(confInst.LoginAttemptsSender.AMQPConfig.PublishMaxTryAttempts),
		libamqp.WithSenderPublishBaseRetryDelay(confInst.LoginAttemptsSender.AMQPConfig.PublishBaseRetryDelay),
		libamqp.WithSenderPublishMaxRetryDelay(confInst.LoginAttemptsSender.AMQPConfig.PublishMaxRetryDelay),
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

	connectorInst, err := libamqp.NewConnector(
		libamqp.WithConnectorURL(confInst.AMQPConnector.URL),
		libamqp.WithConnectorConnectTimeout(confInst.AMQPConnector.ConnectTimeout),
		libamqp.WithConnectorShutdownTimeout(confInst.AMQPConnector.ShutdownTimeout),
		libamqp.WithConnectorConnOpts(connOpts),
		libamqp.WithConnectorSessionOpts(sessOpts),
		libamqp.WithConnectorLogger(logInst),
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

func (cc *ClientContainer) providerKafkaLoginAttemptSender() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}
	if utils.IsNil(confInst.LoginAttemptsSender.KafkaConfig) {
		return nil, errs.NewContainerError(cc.GetName(), "provider: kafka config absent", nil)
	}
	logInst, err := container.GetInstance[logger.Logger](InstanceLogger)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}

	return libkafka.NewSender(
		libkafka.WithSenderClientID(confInst.App.NodeName),
		libkafka.WithSenderBrokers(confInst.LoginAttemptsSender.KafkaConfig.Brokers),
		libkafka.WithSenderTargetName(confInst.LoginAttemptsSender.KafkaConfig.TargetName),
		libkafka.WithSenderConnectTimeout(confInst.LoginAttemptsSender.KafkaConfig.ConnectTimeout),
		libkafka.WithSenderShutdownTimeout(confInst.LoginAttemptsSender.KafkaConfig.ShutdownTimeout),
		libkafka.WithSenderPublishRetry(
			confInst.LoginAttemptsSender.KafkaConfig.PublishMaxTryAttempts,
			confInst.LoginAttemptsSender.KafkaConfig.PublishBaseRetryDelay,
			confInst.LoginAttemptsSender.KafkaConfig.PublishMaxRetryDelay,
		),
		libkafka.WithSenderSecurity(
			confInst.LoginAttemptsSender.KafkaConfig.Username,
			confInst.LoginAttemptsSender.KafkaConfig.Password,
		),
		libkafka.WithSenderBatchOptions(
			confInst.LoginAttemptsSender.KafkaConfig.BatchSize,
			int64(confInst.LoginAttemptsSender.KafkaConfig.BatchBytes),
			confInst.LoginAttemptsSender.KafkaConfig.BatchTimeout,
			confInst.LoginAttemptsSender.KafkaConfig.WriteTimeout,
			kafka.RequiredAcks(confInst.LoginAttemptsSender.KafkaConfig.RequiredAcks),
		),
		libkafka.WithSenderLogger(logInst),
	)
}
