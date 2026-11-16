package config

import (
	"errors"

	libconf "github.com/ElfAstAhe/go-service-template/pkg/config"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func bindFlags(flags *pflag.FlagSet, v *viper.Viper) error {
	err := errors.Join(
		// App
		v.BindPFlag(libconf.KeyAppEnv, flags.Lookup(libconf.FlagAppEnv)),
		v.BindPFlag(libconf.KeyAppInitTimeout, flags.Lookup(libconf.FlagAppInitTimeout)),
		v.BindPFlag(libconf.KeyAppStopTimeout, flags.Lookup(libconf.FlagAppStopTimeout)),
		v.BindPFlag(libconf.KeyAppCloseTimeout, flags.Lookup(libconf.FlagAppCloseTimeout)),
		v.BindPFlag(keyAppNodeName, flags.Lookup(FlagAppNodeName)),
		v.BindPFlag(keyAppMaxListLimit, flags.Lookup(FlagAppMaxListLimit)),
		v.BindPFlag(keyAppTokenIssuer, flags.Lookup(FlagAppTokenIssuer)),
		v.BindPFlag(keyAppCipherKey, flags.Lookup(FlagAppCipherKey)),
		v.BindPFlag(keyAppDefShutdownTimeout, flags.Lookup(FlagAppDefShutdownTimeout)),
		// Auth
		v.BindPFlag(libconf.KeyAuthJWTSecret, flags.Lookup(FlagAuthJWTSecret)),
		v.BindPFlag(libconf.KeyAuthJWTSigningMethod, flags.Lookup(FlagAuthJWTSigningMethod)),
		v.BindPFlag(libconf.KeyAuthAccessTokenTTL, flags.Lookup(FlagAuthAccessTokenTTL)),
		v.BindPFlag(libconf.KeyAuthRefreshTokenTTL, flags.Lookup(FlagAuthRefreshTokenTTL)),
		v.BindPFlag(libconf.KeyAuthRSAPrivateKeyPath, flags.Lookup(FlagAuthRSAPrivateKeyPath)),
		v.BindPFlag(libconf.KeyAuthMasterPasswordSalt, flags.Lookup(FlagAuthMasterPasswordSalt)),
		// HTTP
		v.BindPFlag(libconf.KeyHTTPAddress, flags.Lookup(FlagHTTPAddress)),
		v.BindPFlag(libconf.KeyHTTPReadTimeout, flags.Lookup(FlagHTTPReadTimeout)),
		v.BindPFlag(libconf.KeyHTTPWriteTimeout, flags.Lookup(FlagHTTPWriteTimeout)),
		v.BindPFlag(libconf.KeyHTTPIdleTimeout, flags.Lookup(FlagHTTPIdleTimeout)),
		v.BindPFlag(libconf.KeyHTTPShutdownTimeout, flags.Lookup(FlagHTTPShutdownTimeout)),
		v.BindPFlag(libconf.KeyHTTPPrivateKeyPath, flags.Lookup(FlagHTTPPrivateKeyPath)),
		v.BindPFlag(libconf.KeyHTTPCertificatePath, flags.Lookup(FlagHTTPCertificatePath)),
		v.BindPFlag(libconf.KeyHTTPSecure, flags.Lookup(FlagHTTPSecure)),
		v.BindPFlag(libconf.KeyHTTPMaxRequestBodySize, flags.Lookup(FlagHTTPMaxRequestBodySize)),
		// gRPC
		v.BindPFlag(libconf.KeyGRPCAddress, flags.Lookup(FlagGRPCAddress)),
		v.BindPFlag(libconf.KeyGRPCMaxConnIdle, flags.Lookup(FlagGRPCMaxConnIdle)),
		v.BindPFlag(libconf.KeyGRPCMaxConnAge, flags.Lookup(FlagGRPCMaxConnAge)),
		v.BindPFlag(libconf.KeyGRPCMaxConnAgeGrace, flags.Lookup(FlagGRPCMaxConnAgeGrace)),
		v.BindPFlag(libconf.KeyGRPCTimeout, flags.Lookup(FlagGRPCTimeout)),
		v.BindPFlag(libconf.KeyGRPCKeepAliveTime, flags.Lookup(FlagGRPCKeepAliveTime)),
		v.BindPFlag(libconf.KeyGRPCKeepAliveTimeout, flags.Lookup(FlagGRPCKeepAliveTimeout)),
		v.BindPFlag(libconf.KeyGRPCShutdownTimeout, flags.Lookup(FlagGRPCShutdownTimeout)),
		// Log
		v.BindPFlag(libconf.KeyLogLevel, flags.Lookup(FlagLogLevel)),
		v.BindPFlag(libconf.KeyLogFormat, flags.Lookup(FlagLogFormat)),
		// DB
		v.BindPFlag(libconf.KeyDBDriver, flags.Lookup(FlagDBDriver)),
		v.BindPFlag(libconf.KeyDBDSN, flags.Lookup(FlagDBDSN)),
		v.BindPFlag(libconf.KeyDBMaxOpenConns, flags.Lookup(FlagDBMaxOpenConns)),
		v.BindPFlag(libconf.KeyDBMaxIdleConns, flags.Lookup(FlagDBMaxIdleConns)),
		v.BindPFlag(libconf.KeyDBConnMaxIdleLifetime, flags.Lookup(FlagDBMaxIdleLifetime)),
		v.BindPFlag(libconf.KeyDBConnTimeout, flags.Lookup(FlagDBConnTimeout)),
		// Telemetry
		v.BindPFlag(libconf.KeyTelemetryEnabled, flags.Lookup(FlagTelemetryEnabled)),
		v.BindPFlag(libconf.KeyTelemetryExporterEndpoint, flags.Lookup(FlagTelemetryExporterEndpoint)),
		v.BindPFlag(libconf.KeyTelemetryServiceName, flags.Lookup(FlagTelemetryServiceName)),
		v.BindPFlag(libconf.KeyTelemetrySampleRate, flags.Lookup(FlagTelemetrySampleRate)),
		v.BindPFlag(libconf.KeyTelemetryTimeout, flags.Lookup(FlagTelemetryTimeout)),
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
		// auth audit outbox tail cutter
		v.BindPFlag(keyAuthAuditOutboxTCTailLimit, flags.Lookup(FlagAuthAuditOutboxTCTailLimit)),
		v.BindPFlag(keyAuthAuditOutboxTCTailInterval, flags.Lookup(FlagAuthAuditOutboxTCTailInterval)),
		v.BindPFlag(keyAuthAuditOutboxTCTailCut, flags.Lookup(FlagAuthAuditOutboxTCTailCut)),
		v.BindPFlag(keyAuthAuditOutboxTCWorkerCount, flags.Lookup(FlagAuthAuditOutboxTCWorkerCount)),
		v.BindPFlag(keyAuthAuditOutboxTCDataCapacity, flags.Lookup(FlagAuthAuditOutboxTCDataCapacity)),
		v.BindPFlag(keyAuthAuditOutboxTCCompleteProcessing, flags.Lookup(FlagAuthAuditOutboxTCCompleteProcess)),
		v.BindPFlag(keyAuthAuditOutboxTCStartInterval, flags.Lookup(FlagAuthAuditOutboxTCStartInterval)),
		v.BindPFlag(keyAuthAuditOutboxTCScheduleInterval, flags.Lookup(FlagAuthAuditOutboxTCScheduleInterval)),
		v.BindPFlag(keyAuthAuditOutboxTCStopTimeout, flags.Lookup(FlagAuthAuditOutboxTCStopTimeout)),
		// data audit outbox tail cutter
		v.BindPFlag(keyDataAuditOutboxTCTailLimit, flags.Lookup(FlagDataAuditOutboxTCTailLimit)),
		v.BindPFlag(keyDataAuditOutboxTCTailInterval, flags.Lookup(FlagDataAuditOutboxTCTailInterval)),
		v.BindPFlag(keyDataAuditOutboxTCTailCut, flags.Lookup(FlagDataAuditOutboxTCTailCut)),
		v.BindPFlag(keyDataAuditOutboxTCWorkerCount, flags.Lookup(FlagDataAuditOutboxTCWorkerCount)),
		v.BindPFlag(keyDataAuditOutboxTCDataCapacity, flags.Lookup(FlagDataAuditOutboxTCDataCapacity)),
		v.BindPFlag(keyDataAuditOutboxTCCompleteProcessing, flags.Lookup(FlagDataAuditOutboxTCCompleteProcess)),
		v.BindPFlag(keyDataAuditOutboxTCStartInterval, flags.Lookup(FlagDataAuditOutboxTCStartInterval)),
		v.BindPFlag(keyDataAuditOutboxTCScheduleInterval, flags.Lookup(FlagDataAuditOutboxTCScheduleInterval)),
		v.BindPFlag(keyDataAuditOutboxTCStopTimeout, flags.Lookup(FlagDataAuditOutboxTCStopTimeout)),
	)
	if err != nil {
		return errs.NewConfigError("bind flags with keys", err)
	}

	return nil
}
