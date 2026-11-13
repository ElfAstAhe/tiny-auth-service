package container

import (
	"fmt"

	"github.com/Azure/go-amqp"
	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/go-service-template/pkg/transport/broker"
	"github.com/ElfAstAhe/go-service-template/pkg/transport/broker/amqp/azure"
	libkafka "github.com/ElfAstAhe/go-service-template/pkg/transport/broker/kafka"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/config"
	"github.com/segmentio/kafka-go"
)

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
	connectorInst, err := container.GetInstance[broker.Connector[*amqp.Session]](InstanceAMQPLoginAttemptsConnector)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}
	senderOptsInst, err := container.GetInstance[*amqp.SenderOptions](InstanceAMQPLoginAttemptSenderSenderOpts)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}

	sender, err := azure.NewSender(
		azure.WithSenderConnector(connectorInst),
		azure.WithSenderTargetName(confInst.LoginAttemptsSender.AMQPConfig.TargetName),
		azure.WithSenderLogger(logInst),
		azure.WithSenderOpts(senderOptsInst),
		azure.WithSenderConnectTimeout(confInst.LoginAttemptsSender.AMQPConfig.ConnectTimeout),
		azure.WithSenderShutdownTimeout(confInst.LoginAttemptsSender.AMQPConfig.ShutdownTimeout),
		azure.WithSenderPublishMaxTryAttempts(confInst.LoginAttemptsSender.AMQPConfig.PublishMaxTryAttempts),
		azure.WithSenderPublishBaseRetryDelay(confInst.LoginAttemptsSender.AMQPConfig.PublishBaseRetryDelay),
		azure.WithSenderPublishMaxRetryDelay(confInst.LoginAttemptsSender.AMQPConfig.PublishMaxRetryDelay),
	)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), fmt.Sprintf("provider: create %s instance failed", InstanceAMQPLoginAttemptSender), err)
	}

	return sender, nil
}

//goland:noinspection DuplicatedCode
func (cc *ClientContainer) providerAMQPLoginAttemptsConnector() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}
	logInst, err := container.GetInstance[logger.Logger](InstanceLogger)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}
	connOptsInst, err := container.GetInstance[*amqp.ConnOptions](InstanceAMQPLoginAttemptsConnectorConnOpts)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}
	sessOptsInst, err := container.GetInstance[*amqp.SessionOptions](InstanceAMQPLoginAttemptsConnectorSessOpts)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}

	connectorInst, err := azure.NewConnector(
		azure.WithConnectorURL(confInst.LoginAttemptsConnector.URL),
		azure.WithConnectorConnectTimeout(confInst.LoginAttemptsConnector.ConnectTimeout),
		azure.WithConnectorShutdownTimeout(confInst.LoginAttemptsConnector.ShutdownTimeout),
		azure.WithConnectorConnOpts(connOptsInst),
		azure.WithConnectorSessionOpts(sessOptsInst),
		azure.WithConnectorLogger(logInst),
	)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), fmt.Sprintf("provider: create %s instance failed", InstanceAMQPLoginAttemptsConnector), err)
	}

	return connectorInst, nil
}

func (cc *ClientContainer) providerAMQPLoginAttemptsConnectorConnOpts() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}

	return &amqp.ConnOptions{
		HostName:     confInst.App.NodeName,
		IdleTimeout:  confInst.LoginAttemptsConnector.IdleTimeout,
		WriteTimeout: confInst.LoginAttemptsConnector.WriteTimeout,
		SASLType:     amqp.SASLTypePlain(confInst.LoginAttemptsConnector.Username, confInst.LoginAttemptsConnector.Password),
	}, nil
}

func (cc *ClientContainer) providerAMQPLoginAttemptsConnectorSessOpts() (any, error) {
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

func (cc *ClientContainer) providerAMQPDataAuditSender() (any, error) {
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
	connectorInst, err := container.GetInstance[broker.Connector[*amqp.Session]](InstanceAMQPDataAuditConnector)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}
	senderOptsInst, err := container.GetInstance[*amqp.SenderOptions](InstanceAMQPDataAuditSenderSenderOpts)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}

	sender, err := azure.NewSender(
		azure.WithSenderConnector(connectorInst),
		azure.WithSenderTargetName(confInst.DataAuditSender.AMQPConfig.TargetName),
		azure.WithSenderLogger(logInst),
		azure.WithSenderOpts(senderOptsInst),
		azure.WithSenderConnectTimeout(confInst.DataAuditSender.AMQPConfig.ConnectTimeout),
		azure.WithSenderShutdownTimeout(confInst.DataAuditSender.AMQPConfig.ShutdownTimeout),
		azure.WithSenderPublishMaxTryAttempts(confInst.DataAuditSender.AMQPConfig.PublishMaxTryAttempts),
		azure.WithSenderPublishBaseRetryDelay(confInst.DataAuditSender.AMQPConfig.PublishBaseRetryDelay),
		azure.WithSenderPublishMaxRetryDelay(confInst.DataAuditSender.AMQPConfig.PublishMaxRetryDelay),
	)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), fmt.Sprintf("provider: create %s instance failed", InstanceAMQPDataAuditSender), err)
	}

	return sender, nil
}

func (cc *ClientContainer) providerAMQPDataAuditSenderSenderOpts() (any, error) {
	return &amqp.SessionOptions{
		MaxLinks: 4,
	}, nil
}

func (cc *ClientContainer) providerAMQPDataAuditConnector() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}
	logInst, err := container.GetInstance[logger.Logger](InstanceLogger)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}
	connOptsInst, err := container.GetInstance[*amqp.ConnOptions](InstanceAMQPDataAuditConnectorConnOpts)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}
	sessOptsInst, err := container.GetInstance[*amqp.SessionOptions](InstanceAMQPDataAuditConnectorSessOpts)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}

	connectorInst, err := azure.NewConnector(
		azure.WithConnectorURL(confInst.DataAuditConnector.URL),
		azure.WithConnectorConnectTimeout(confInst.DataAuditConnector.ConnectTimeout),
		azure.WithConnectorShutdownTimeout(confInst.DataAuditConnector.ShutdownTimeout),
		azure.WithConnectorConnOpts(connOptsInst),
		azure.WithConnectorSessionOpts(sessOptsInst),
		azure.WithConnectorLogger(logInst),
	)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), fmt.Sprintf("provider: create %s instance failed", InstanceAMQPDataAuditConnector), err)
	}

	return connectorInst, nil
}

func (cc *ClientContainer) providerAMQPDataAuditConnectorConnOpts() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}

	return &amqp.ConnOptions{
		HostName:     confInst.App.NodeName,
		IdleTimeout:  confInst.DataAuditConnector.IdleTimeout,
		WriteTimeout: confInst.DataAuditConnector.WriteTimeout,
		SASLType:     amqp.SASLTypePlain(confInst.DataAuditConnector.Username, confInst.DataAuditConnector.Password),
	}, nil
}

func (cc *ClientContainer) providerAMQPDataAuditConnectorSessOpts() (any, error) {
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

func (cc *ClientContainer) providerKafkaDataAuditSender() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}
	if utils.IsNil(confInst.DataAuditSender.KafkaConfig) {
		return nil, errs.NewContainerError(cc.GetName(), "provider: kafka config absent", nil)
	}
	logInst, err := container.GetInstance[logger.Logger](InstanceLogger)
	if err != nil {
		return nil, errs.NewContainerError(cc.GetName(), "provider: retrieve instance failed", err)
	}

	return libkafka.NewSender(
		libkafka.WithSenderClientID(confInst.App.NodeName),
		libkafka.WithSenderBrokers(confInst.DataAuditSender.KafkaConfig.Brokers),
		libkafka.WithSenderTargetName(confInst.DataAuditSender.KafkaConfig.TargetName),
		libkafka.WithSenderConnectTimeout(confInst.DataAuditSender.KafkaConfig.ConnectTimeout),
		libkafka.WithSenderShutdownTimeout(confInst.DataAuditSender.KafkaConfig.ShutdownTimeout),
		libkafka.WithSenderPublishRetry(
			confInst.DataAuditSender.KafkaConfig.PublishMaxTryAttempts,
			confInst.DataAuditSender.KafkaConfig.PublishBaseRetryDelay,
			confInst.DataAuditSender.KafkaConfig.PublishMaxRetryDelay,
		),
		libkafka.WithSenderSecurity(
			confInst.DataAuditSender.KafkaConfig.Username,
			confInst.DataAuditSender.KafkaConfig.Password,
		),
		libkafka.WithSenderBatchOptions(
			confInst.DataAuditSender.KafkaConfig.BatchSize,
			int64(confInst.DataAuditSender.KafkaConfig.BatchBytes),
			confInst.DataAuditSender.KafkaConfig.BatchTimeout,
			confInst.DataAuditSender.KafkaConfig.WriteTimeout,
			kafka.RequiredAcks(confInst.DataAuditSender.KafkaConfig.RequiredAcks),
		),
		libkafka.WithSenderLogger(logInst),
	)
}
