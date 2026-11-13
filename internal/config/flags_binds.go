package config

import (
	"errors"

	"github.com/ElfAstAhe/go-service-template/pkg/config"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func bindFlags(flags *pflag.FlagSet, v *viper.Viper) error {
	err := errors.Join(
		// App
		v.BindPFlag(config.KeyAppEnv, flags.Lookup(config.FlagAppEnv)),
		v.BindPFlag(config.KeyAppInitTimeout, flags.Lookup(config.FlagAppInitTimeout)),
		v.BindPFlag(config.KeyAppStopTimeout, flags.Lookup(config.FlagAppStopTimeout)),
		v.BindPFlag(config.KeyAppCloseTimeout, flags.Lookup(config.FlagAppCloseTimeout)),
		v.BindPFlag(keyAppNodeName, flags.Lookup(FlagAppNodeName)),
		v.BindPFlag(keyAppMaxListLimit, flags.Lookup(FlagAppMaxListLimit)),
		v.BindPFlag(keyAppTokenIssuer, flags.Lookup(FlagAppTokenIssuer)),
		v.BindPFlag(keyAppCipherKey, flags.Lookup(FlagAppCipherKey)),
		v.BindPFlag(keyAppDefShutdownTimeout, flags.Lookup(FlagAppDefShutdownTimeout)),
		// Auth
		v.BindPFlag(config.KeyAuthJWTSecret, flags.Lookup(FlagAuthJWTSecret)),
		v.BindPFlag(config.KeyAuthJWTSigningMethod, flags.Lookup(FlagAuthJWTSigningMethod)),
		v.BindPFlag(config.KeyAuthAccessTokenTTL, flags.Lookup(FlagAuthAccessTokenTTL)),
		v.BindPFlag(config.KeyAuthRefreshTokenTTL, flags.Lookup(FlagAuthRefreshTokenTTL)),
		v.BindPFlag(config.KeyAuthRSAPrivateKeyPath, flags.Lookup(FlagAuthRSAPrivateKeyPath)),
		v.BindPFlag(config.KeyAuthMasterPasswordSalt, flags.Lookup(FlagAuthMasterPasswordSalt)),
		// HTTP
		v.BindPFlag(config.KeyHTTPAddress, flags.Lookup(FlagHTTPAddress)),
		v.BindPFlag(config.KeyHTTPReadTimeout, flags.Lookup(FlagHTTPReadTimeout)),
		v.BindPFlag(config.KeyHTTPWriteTimeout, flags.Lookup(FlagHTTPWriteTimeout)),
		v.BindPFlag(config.KeyHTTPIdleTimeout, flags.Lookup(FlagHTTPIdleTimeout)),
		v.BindPFlag(config.KeyHTTPShutdownTimeout, flags.Lookup(FlagHTTPShutdownTimeout)),
		v.BindPFlag(config.KeyHTTPPrivateKeyPath, flags.Lookup(FlagHTTPPrivateKeyPath)),
		v.BindPFlag(config.KeyHTTPCertificatePath, flags.Lookup(FlagHTTPCertificatePath)),
		v.BindPFlag(config.KeyHTTPSecure, flags.Lookup(FlagHTTPSecure)),
		v.BindPFlag(config.KeyHTTPMaxRequestBodySize, flags.Lookup(FlagHTTPMaxRequestBodySize)),
		// gRPC
		v.BindPFlag(config.KeyGRPCAddress, flags.Lookup(FlagGRPCAddress)),
		v.BindPFlag(config.KeyGRPCMaxConnIdle, flags.Lookup(FlagGRPCMaxConnIdle)),
		v.BindPFlag(config.KeyGRPCMaxConnAge, flags.Lookup(FlagGRPCMaxConnAge)),
		v.BindPFlag(config.KeyGRPCMaxConnAgeGrace, flags.Lookup(FlagGRPCMaxConnAgeGrace)),
		v.BindPFlag(config.KeyGRPCTimeout, flags.Lookup(FlagGRPCTimeout)),
		v.BindPFlag(config.KeyGRPCKeepAliveTime, flags.Lookup(FlagGRPCKeepAliveTime)),
		v.BindPFlag(config.KeyGRPCKeepAliveTimeout, flags.Lookup(FlagGRPCKeepAliveTimeout)),
		v.BindPFlag(config.KeyGRPCShutdownTimeout, flags.Lookup(FlagGRPCShutdownTimeout)),
		// Log
		v.BindPFlag(config.KeyLogLevel, flags.Lookup(FlagLogLevel)),
		v.BindPFlag(config.KeyLogFormat, flags.Lookup(FlagLogFormat)),
		// DB
		v.BindPFlag(config.KeyDBDriver, flags.Lookup(FlagDBDriver)),
		v.BindPFlag(config.KeyDBDSN, flags.Lookup(FlagDBDSN)),
		v.BindPFlag(config.KeyDBMaxOpenConns, flags.Lookup(FlagDBMaxOpenConns)),
		v.BindPFlag(config.KeyDBMaxIdleConns, flags.Lookup(FlagDBMaxIdleConns)),
		v.BindPFlag(config.KeyDBConnMaxIdleLifetime, flags.Lookup(FlagDBMaxIdleLifetime)),
		v.BindPFlag(config.KeyDBConnTimeout, flags.Lookup(FlagDBConnTimeout)),
		// Telemetry
		v.BindPFlag(config.KeyTelemetryEnabled, flags.Lookup(FlagTelemetryEnabled)),
		v.BindPFlag(config.KeyTelemetryExporterEndpoint, flags.Lookup(FlagTelemetryExporterEndpoint)),
		v.BindPFlag(config.KeyTelemetryServiceName, flags.Lookup(FlagTelemetryServiceName)),
		v.BindPFlag(config.KeyTelemetrySampleRate, flags.Lookup(FlagTelemetrySampleRate)),
		v.BindPFlag(config.KeyTelemetryTimeout, flags.Lookup(FlagTelemetryTimeout)),
		// login attempts sender amqp connector
		v.BindPFlag(keyLoginAttemptsSenderAMQPConnectorURL, flags.Lookup(FlagLoginAttemptsSenderAMQPConnectorURL)),
		v.BindPFlag(keyLoginAttemptsSenderAMQPConnectorUsername, flags.Lookup(FlagLoginAttemptsSenderAMQPConnectorUsername)),
		v.BindPFlag(keyLoginAttemptsSenderAMQPConnectorPassword, flags.Lookup(FlagLoginAttemptsSenderAMQPConnectorPassword)),
		v.BindPFlag(keyLoginAttemptsSenderAMQPConnectorConnectTimeout, flags.Lookup(FlagLoginAttemptsSenderAMQPConnectorConnectTimeout)),
		v.BindPFlag(keyLoginAttemptsSenderAMQPConnectorWriteTimeout, flags.Lookup(FlagLoginAttemptsSenderAMQPConnectorWriteTimeout)),
		v.BindPFlag(keyLoginAttemptsSenderAMQPConnectorIdleTimeout, flags.Lookup(FlagLoginAttemptsSenderAMQPConnectorIdleTimeout)),
		v.BindPFlag(keyLoginAttemptsSenderAMQPConnectorShutdownTimeout, flags.Lookup(FlagLoginAttemptsSenderAMQPConnectorShutdownTimeout)),
		// login attempts sender
		v.BindPFlag(keyLoginAttemptsSenderKind, flags.Lookup(FlagLoginAttemptsSenderKind)),
		v.BindPFlag(keyLoginAttemptsSenderNotifyTimeout, flags.Lookup(FlagLoginAttemptsSenderNotifyTimeout)),
		// login attempts sender amqp
		v.BindPFlag(keyLoginAttemptsSenderAMQPTargetName, flags.Lookup(FlagLoginAttemptsSenderAMQPTargetName)),
		v.BindPFlag(keyLoginAttemptsSenderAMQPConnectTimeout, flags.Lookup(FlagLoginAttemptsSenderAMQPConnectTimeout)),
		v.BindPFlag(keyLoginAttemptsSenderAMQPShutdownTimeout, flags.Lookup(FlagLoginAttemptsSenderAMQPShutdownTimeout)),
		v.BindPFlag(keyLoginAttemptsSenderAMQPPublishMaxTryAttempts, flags.Lookup(FlagLoginAttemptsSenderAMQPPublishMaxTryAttempts)),
		v.BindPFlag(keyLoginAttemptsSenderAMQPPublishBaseRetryDelay, flags.Lookup(FlagLoginAttemptsSenderAMQPPublishBaseRetryDelay)),
		v.BindPFlag(keyLoginAttemptsSenderAMQPPublishMaxRetryDelay, flags.Lookup(FlagLoginAttemptsSenderAMQPPublishMaxRetryDelay)),
		// login attempts sender kafka
		v.BindPFlag(keyLoginAttemptsSenderKafkaBrokers, flags.Lookup(FlagLoginAttemptsSenderKafkaBrokers)),
		v.BindPFlag(keyLoginAttemptsSenderKafkaTargetName, flags.Lookup(FlagLoginAttemptsSenderKafkaTargetName)),
		v.BindPFlag(keyLoginAttemptsSenderKafkaConnectTimeout, flags.Lookup(FlagLoginAttemptsSenderKafkaConnectTimeout)),
		v.BindPFlag(keyLoginAttemptsSenderKafkaIdleTimeout, flags.Lookup(FlagLoginAttemptsSenderKafkaIdleTimeout)),
		v.BindPFlag(keyLoginAttemptsSenderKafkaShutdownTimeout, flags.Lookup(FlagLoginAttemptsSenderKafkaShutdownTimeout)),
		v.BindPFlag(keyLoginAttemptsSenderKafkaPublishMaxTryAttempts, flags.Lookup(FlagLoginAttemptsSenderKafkaPublishMaxTryAttempts)),
		v.BindPFlag(keyLoginAttemptsSenderKafkaPublishMaxRetryDelay, flags.Lookup(FlagLoginAttemptsSenderKafkaPublishMaxRetryDelay)),
		v.BindPFlag(keyLoginAttemptsSenderKafkaPublishBaseRetryDelay, flags.Lookup(FlagLoginAttemptsSenderKafkaPublishBaseRetryDelay)),
		v.BindPFlag(keyLoginAttemptsSenderKafkaUsername, flags.Lookup(FlagLoginAttemptsSenderKafkaUsername)),
		v.BindPFlag(keyLoginAttemptsSenderKafkaPassword, flags.Lookup(FlagLoginAttemptsSenderKafkaPassword)),
		v.BindPFlag(keyLoginAttemptsSenderKafkaBatchSize, flags.Lookup(FlagLoginAttemptsSenderKafkaBatchSize)),
		v.BindPFlag(keyLoginAttemptsSenderKafkaBatchBytes, flags.Lookup(FlagLoginAttemptsSenderKafkaBatchBytes)),
		v.BindPFlag(keyLoginAttemptsSenderKafkaBatchTimeout, flags.Lookup(FlagLoginAttemptsSenderKafkaBatchTimeout)),
		v.BindPFlag(keyLoginAttemptsSenderKafkaWriteTimeout, flags.Lookup(FlagLoginAttemptsSenderKafkaWriteTimeout)),
		v.BindPFlag(keyLoginAttemptsSenderKafkaRequiredAcks, flags.Lookup(FlagLoginAttemptsSenderKafkaRequiredAcks)),
		// data audit sender amqp connector
		v.BindPFlag(keyDataAuditSenderAMQPConnectorURL, flags.Lookup(FlagDataAuditSenderAMQPConnectorURL)),
		v.BindPFlag(keyDataAuditSenderAMQPConnectorUsername, flags.Lookup(FlagDataAuditSenderAMQPConnectorUsername)),
		v.BindPFlag(keyDataAuditSenderAMQPConnectorPassword, flags.Lookup(FlagDataAuditSenderAMQPConnectorPassword)),
		v.BindPFlag(keyDataAuditSenderAMQPConnectorConnectTimeout, flags.Lookup(FlagDataAuditSenderAMQPConnectorConnectTimeout)),
		v.BindPFlag(keyDataAuditSenderAMQPConnectorWriteTimeout, flags.Lookup(FlagDataAuditSenderAMQPConnectorWriteTimeout)),
		v.BindPFlag(keyDataAuditSenderAMQPConnectorIdleTimeout, flags.Lookup(FlagDataAuditSenderAMQPConnectorIdleTimeout)),
		v.BindPFlag(keyDataAuditSenderAMQPConnectorShutdownTimeout, flags.Lookup(FlagDataAuditSenderAMQPConnectorShutdownTimeout)),
		// data audit sender
		v.BindPFlag(keyDataAuditSenderKind, flags.Lookup(FlagDataAuditSenderKind)),
		v.BindPFlag(keyDataAuditSenderNotifyTimeout, flags.Lookup(FlagDataAuditSenderNotifyTimeout)),
		// data audit sender amqp
		v.BindPFlag(keyDataAuditSenderAMQPTargetName, flags.Lookup(FlagDataAuditSenderAMQPTargetName)),
		v.BindPFlag(keyDataAuditSenderAMQPConnectTimeout, flags.Lookup(FlagDataAuditSenderAMQPConnectTimeout)),
		v.BindPFlag(keyDataAuditSenderAMQPShutdownTimeout, flags.Lookup(FlagDataAuditSenderAMQPShutdownTimeout)),
		v.BindPFlag(keyDataAuditSenderAMQPPublishMaxTryAttempts, flags.Lookup(FlagDataAuditSenderAMQPPublishMaxTryAttempts)),
		v.BindPFlag(keyDataAuditSenderAMQPPublishBaseRetryDelay, flags.Lookup(FlagDataAuditSenderAMQPPublishBaseRetryDelay)),
		v.BindPFlag(keyDataAuditSenderAMQPPublishMaxRetryDelay, flags.Lookup(FlagDataAuditSenderAMQPPublishMaxRetryDelay)),
		// data audit sender kafka
		v.BindPFlag(keyDataAuditSenderKafkaBrokers, flags.Lookup(FlagDataAuditSenderKafkaBrokers)),
		v.BindPFlag(keyDataAuditSenderKafkaTargetName, flags.Lookup(FlagDataAuditSenderKafkaTargetName)),
		v.BindPFlag(keyDataAuditSenderKafkaConnectTimeout, flags.Lookup(FlagDataAuditSenderKafkaConnectTimeout)),
		v.BindPFlag(keyDataAuditSenderKafkaIdleTimeout, flags.Lookup(FlagDataAuditSenderKafkaIdleTimeout)),
		v.BindPFlag(keyDataAuditSenderKafkaShutdownTimeout, flags.Lookup(FlagDataAuditSenderKafkaShutdownTimeout)),
		v.BindPFlag(keyDataAuditSenderKafkaPublishMaxTryAttempts, flags.Lookup(FlagDataAuditSenderKafkaPublishMaxTryAttempts)),
		v.BindPFlag(keyDataAuditSenderKafkaPublishMaxRetryDelay, flags.Lookup(FlagDataAuditSenderKafkaPublishMaxRetryDelay)),
		v.BindPFlag(keyDataAuditSenderKafkaPublishBaseRetryDelay, flags.Lookup(FlagDataAuditSenderKafkaPublishBaseRetryDelay)),
		v.BindPFlag(keyDataAuditSenderKafkaUsername, flags.Lookup(FlagDataAuditSenderKafkaUsername)),
		v.BindPFlag(keyDataAuditSenderKafkaPassword, flags.Lookup(FlagDataAuditSenderKafkaPassword)),
		v.BindPFlag(keyDataAuditSenderKafkaBatchSize, flags.Lookup(FlagDataAuditSenderKafkaBatchSize)),
		v.BindPFlag(keyDataAuditSenderKafkaBatchBytes, flags.Lookup(FlagLoginAttemptsSenderKafkaBatchBytes)),
		v.BindPFlag(keyDataAuditSenderKafkaBatchTimeout, flags.Lookup(FlagDataAuditSenderKafkaBatchTimeout)),
		v.BindPFlag(keyDataAuditSenderKafkaWriteTimeout, flags.Lookup(FlagDataAuditSenderKafkaWriteTimeout)),
		v.BindPFlag(keyDataAuditSenderKafkaRequiredAcks, flags.Lookup(FlagDataAuditSenderKafkaRequiredAcks)),
	)
	if err != nil {
		return errs.NewConfigError("bind flags with keys", err)
	}

	return nil
}
