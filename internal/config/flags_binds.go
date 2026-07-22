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
		// svc-creds
		v.BindPFlag(keySvcCredsUsername, flags.Lookup(FlagCredsUsername)),
		v.BindPFlag(keySvcCredsPassword, flags.Lookup(FlagCredsPassword)),
		v.BindPFlag(keySvcCredsScheduleInterval, flags.Lookup(FlagCredsScheduleInterval)),
		v.BindPFlag(keySvcCredsErrorScheduleInterval, flags.Lookup(FlagCredsErrorScheduleInterval)),
		// data-audit-client
		v.BindPFlag(keyDataAuditClientBaseURL, flags.Lookup(FlagDataAuditClientBaseURL)),
		v.BindPFlag(keyDataAuditClientTimeout, flags.Lookup(FlagDataAuditClientTimeout)),
		v.BindPFlag(keyDataAuditClientWorkerCount, flags.Lookup(FlagDataAuditClientWorkerCount)),
		v.BindPFlag(keyDataAuditClientDataCapacity, flags.Lookup(FlagDataAuditClientDataCapacity)),
		v.BindPFlag(keyDataAuditClientCompleteProcessing, flags.Lookup(FlagDataAuditClientCompleteProcessing)),
		v.BindPFlag(keyDataAuditClientShutdownTimeout, flags.Lookup(FlagDataAuditClientShutdownTimeout)),
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
		// amqp connector
		v.BindPFlag(keyAMQPConnectorURL, flags.Lookup(FlagAMQPConnectorURL)),
		v.BindPFlag(keyAMQPConnectorUsername, flags.Lookup(FlagAMQPConnectorUsername)),
		v.BindPFlag(keyAMQPConnectorPassword, flags.Lookup(FlagAMQPConnectorPassword)),
		v.BindPFlag(keyAMQPConnectorConnectTimeout, flags.Lookup(FlagAMQPConnectorConnectTimeout)),
		v.BindPFlag(keyAMQPConnectorWriteTimeout, flags.Lookup(FlagAMQPConnectorWriteTimeout)),
		v.BindPFlag(keyAMQPConnectorIdleTimeout, flags.Lookup(FlagAMQPConnectorIdleTimeout)),
		v.BindPFlag(keyAMQPConnectorShutdownTimeout, flags.Lookup(FlagAMQPConnectorShutdownTimeout)),
		// amqp sender
		v.BindPFlag(keyLoginAttemptsSenderTargetName, flags.Lookup(FlagLoginAttemptsSenderTargetName)),
		v.BindPFlag(keyLoginAttemptsSenderConnectTimeout, flags.Lookup(FlagLoginAttemptsSenderConnectTimeout)),
		v.BindPFlag(keyLoginAttemptsSenderNotifyTimeout, flags.Lookup(FlagLoginAttemptsSenderNotifyTimeout)),
		v.BindPFlag(keyLoginAttemptsSenderShutdownTimeout, flags.Lookup(FlagLoginAttemptsSenderShutdownTimeout)),
		v.BindPFlag(keyLoginAttemptsSenderPublishMaxTryAttempts, flags.Lookup(FlagLoginAttemptsSenderPublishMaxTryAttempts)),
		v.BindPFlag(keyLoginAttemptsSenderPublishBaseRetryDelay, flags.Lookup(FlagLoginAttemptsSenderPublishBaseRetryDelay)),
		v.BindPFlag(keyLoginAttemptsSenderPublishMaxRetryDelay, flags.Lookup(FlagLoginAttemptsSenderPublishMaxRetryDelay)),
	)
	if err != nil {
		return errs.NewConfigError("bind flags with keys", err)
	}

	return nil
}
