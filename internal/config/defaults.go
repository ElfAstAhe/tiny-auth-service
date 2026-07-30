package config

import (
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/config"
	"github.com/spf13/viper"
)

// app
const (
	defaultAppEnv             config.AppEnv = config.AppEnvDevelopment
	defaultAppNodeName        string        = ApplicationName
	defaultMaxListLimit       int           = 100
	defaultTokenIssuer        string        = "tiny-auth-service"
	defaultDefShutdownTimeout time.Duration = 15 * time.Second
)

// creds
const (
	defaultCredsScheduleInterval      time.Duration = 5 * time.Minute
	defaultCredsErrorScheduleInterval time.Duration = 3 * time.Second
)

// audit client
const (
	defaultAuditClientTimeout            time.Duration = 5 * time.Second
	defaultAuditClientWorkerCount        int           = 4
	defaultAuditClientDataCapacity       int           = 10000
	defaultAuditClientCompleteProcessing bool          = true
	defaultAuditClientShutdownTimeout    time.Duration = 15 * time.Second
)

// amqp connector
const (
	defaultAMQPConnectorUsername string = "svc-auth"
	defaultAMQPConnectorPassword string = "test"
)

// amqp login attempts sender (FQQN artemis style)
const (
	defaultLoginAttemptsSenderTargetName string        = "tiny.auth::login.attempts"
	defaultLoginAttemptsNotifyTimeout    time.Duration = 2 * time.Second
)

func applyDefaults(v *viper.Viper) {
	// App
	v.SetDefault(config.KeyAppEnv, defaultAppEnv)
	v.SetDefault(config.KeyAppInitTimeout, config.DefaultAppInitTimeout)
	v.SetDefault(config.KeyAppStopTimeout, config.DefaultAppStopTimeout)
	v.SetDefault(config.KeyAppCloseTimeout, config.DefaultAppCloseTimeout)
	v.SetDefault(keyAppNodeName, defaultAppNodeName)
	v.SetDefault(keyAppMaxListLimit, defaultMaxListLimit)
	v.SetDefault(keyAppTokenIssuer, defaultTokenIssuer)
	v.SetDefault(keyAppDefShutdownTimeout, defaultDefShutdownTimeout)

	// creds
	v.SetDefault(keySvcCredsScheduleInterval, defaultCredsScheduleInterval)
	v.SetDefault(keySvcCredsErrorScheduleInterval, defaultCredsErrorScheduleInterval)

	// data-audit-client
	v.SetDefault(keyDataAuditClientTimeout, defaultAuditClientTimeout)
	v.SetDefault(keyDataAuditClientWorkerCount, defaultAuditClientWorkerCount)
	v.SetDefault(keyDataAuditClientDataCapacity, defaultAuditClientDataCapacity)
	v.SetDefault(keyDataAuditClientCompleteProcessing, defaultAuditClientCompleteProcessing)
	v.SetDefault(keyDataAuditClientShutdownTimeout, defaultAuditClientShutdownTimeout)

	// Auth
	v.SetDefault(config.KeyAuthJWTSigningMethod, config.DefaultAuthSigningMethod)
	v.SetDefault(config.KeyAuthAccessTokenTTL, config.DefaultAuthAccessTokenTTL)
	v.SetDefault(config.KeyAuthRefreshTokenTTL, config.DefaultAuthRefreshTokenTTL)

	// HTTP
	v.SetDefault(config.KeyHTTPAddress, config.DefaultHTTPAddress)
	v.SetDefault(config.KeyHTTPReadTimeout, config.DefaultHTTPReadTimeout)
	v.SetDefault(config.KeyHTTPWriteTimeout, config.DefaultHTTPWriteTimeout)
	v.SetDefault(config.KeyHTTPIdleTimeout, config.DefaultHTTPIdleTimeout)
	v.SetDefault(config.KeyHTTPShutdownTimeout, config.DefaultHTTPShutdownTimeout)
	v.SetDefault(config.KeyHTTPSecure, config.DefaultHTTPSecure)
	v.SetDefault(config.KeyHTTPMaxRequestBodySize, config.DefaultHTTPMaxRequestBodySize)

	// gRPC
	v.SetDefault(config.KeyGRPCAddress, config.DefaultGRPCAddress)
	v.SetDefault(config.KeyGRPCMaxConnIdle, config.DefaultGRPCMaxConnIdle)
	v.SetDefault(config.KeyGRPCMaxConnAge, config.DefaultGRPCMaxConnAge)
	v.SetDefault(config.KeyGRPCMaxConnAgeGrace, config.DefaultGRPCMaxConnAgeGrace)
	v.SetDefault(config.KeyGRPCTimeout, config.DefaultGRPCTimeout)
	v.SetDefault(config.KeyGRPCKeepAliveTime, config.DefaultGRPCKeepAliveTime)
	v.SetDefault(config.KeyGRPCKeepAliveTimeout, config.DefaultGRPCKeepAliveTimeout)
	v.SetDefault(config.KeyGRPCShutdownTimeout, config.DefaultGRPCShutdownTimeout)

	// DB
	v.SetDefault(config.KeyDBDriver, config.DefaultDBDriver)
	v.SetDefault(config.KeyDBDSN, config.DefaultDBDSN)
	v.SetDefault(config.KeyDBMaxOpenConns, config.DefaultDBMaxOpenConns)
	v.SetDefault(config.KeyDBMaxIdleConns, config.DefaultDBMaxIdleConns)
	v.SetDefault(config.KeyDBConnMaxIdleLifetime, config.DefaultDBConnMaxIdleLifetime)
	v.SetDefault(config.KeyDBConnTimeout, config.DefaultDBConnTimeout)

	// Log
	v.SetDefault(config.KeyLogLevel, config.DefaultLogLevel)
	v.SetDefault(config.KeyLogFormat, config.DefaultLogFormat)

	// Telemetry
	v.SetDefault(config.KeyTelemetryEnabled, config.DefaultTelemetryEnabled)
	v.SetDefault(config.KeyTelemetryExporterEndpoint, config.DefaultTelemetryExporterEndpoint)
	v.SetDefault(config.KeyTelemetrySampleRate, config.DefaultTelemetrySampleRate)
	v.SetDefault(config.KeyTelemetryTimeout, config.DefaultTelemetryTimeout)

	// amqp connector
	v.SetDefault(keyAMQPConnectorURL, config.DefaultAMQPConnectorURL)
	v.SetDefault(keyAMQPConnectorUsername, defaultAMQPConnectorUsername)
	v.SetDefault(keyAMQPConnectorPassword, defaultAMQPConnectorPassword)
	v.SetDefault(keyAMQPConnectorConnectTimeout, config.DefaultAMQPSenderConnectTimeout)
	v.SetDefault(keyAMQPConnectorWriteTimeout, config.DefaultAMQPConnectorWriteTimeout)
	v.SetDefault(keyAMQPConnectorIdleTimeout, config.DefaultAMQPConnectorIdleTimeout)
	v.SetDefault(keyAMQPConnectorShutdownTimeout, config.DefaultAMQPConnectorShutdownTimeout)
	// amqp sender
	v.SetDefault(keyLoginAttemptsSenderTargetName, defaultLoginAttemptsSenderTargetName)
	v.SetDefault(keyLoginAttemptsSenderConnectTimeout, config.DefaultAMQPSenderConnectTimeout)
	v.SetDefault(keyLoginAttemptsSenderNotifyTimeout, defaultLoginAttemptsNotifyTimeout)
	v.SetDefault(keyLoginAttemptsSenderShutdownTimeout, config.DefaultAMQPSenderShutdownTimeout)
	v.SetDefault(keyLoginAttemptsSenderPublishMaxTryAttempts, config.DefaultAMQPSenderPublishMaxTryAttempts)
	v.SetDefault(keyLoginAttemptsSenderPublishBaseRetryDelay, config.DefaultAMQPSenderPublishBaseRetryDelay)
	v.SetDefault(keyLoginAttemptsSenderPublishMaxRetryDelay, config.DefaultAMQPSenderPublishMaxRetryDelay)
}
