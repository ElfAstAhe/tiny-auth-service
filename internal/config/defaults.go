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

// db
const (
	defaultDBDriver string = "postgres"
	defaultDBDSN    string = "postgres://svc_auth:password@localhost:5432/test?sslmode=disable&search_path=auth_db"
)

// audit client
const (
	defaultDataAuditClientTimeout            time.Duration = 5 * time.Second
	defaultDataAuditClientWorkerCount        int           = 4
	defaultDataAuditClientDataCapacity       int           = 10000
	defaultDataAuditClientCompleteProcessing bool          = true
	defaultDataAuditClientShutdownTimeout    time.Duration = 15 * time.Second
)

// amqp connector
const (
	defaultAMQPConnectorUsername string = "svc-auth"
	defaultAMQPConnectorPassword string = "test"
)

// login attempts
const (
	// defaultLoginAttemptsSenderKind sender kind, values: amqp, kafka
	defaultLoginAttemptsSenderKind string = "amqp"
	// defaultLoginAttemptsSenderNotifyTimeout
	defaultLoginAttemptsSenderNotifyTimeout time.Duration = 2 * time.Second
)

// amqp login attempts target name (FQQN artemis style)
const (
	defaultLoginAttemptsSenderAMQPConfigTargetName string = "tiny.auth::login.attempts"
)

// kafka login attempts target name
const (
	defaultLoginAttemptsSenderKafkaConfigTargetName string = "tiny.auth.login.attempts"
)

//goland:noinspection DuplicatedCode
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
	v.SetDefault(keyDataAuditClientTimeout, defaultDataAuditClientTimeout)
	v.SetDefault(keyDataAuditClientWorkerCount, defaultDataAuditClientWorkerCount)
	v.SetDefault(keyDataAuditClientDataCapacity, defaultDataAuditClientDataCapacity)
	v.SetDefault(keyDataAuditClientCompleteProcessing, defaultDataAuditClientCompleteProcessing)
	v.SetDefault(keyDataAuditClientShutdownTimeout, defaultDataAuditClientShutdownTimeout)

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
	v.SetDefault(config.KeyDBDriver, defaultDBDriver)
	v.SetDefault(config.KeyDBDSN, defaultDBDSN)
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

	// login attempts
	v.SetDefault(keyLoginAttemptsSenderKind, defaultLoginAttemptsSenderKind)
	v.SetDefault(keyLoginAttemptsSenderNotifyTimeout, defaultLoginAttemptsSenderNotifyTimeout)

	// amqp login attempts
	v.SetDefault(keyLoginAttemptsSenderAMQPConfigTargetName, defaultLoginAttemptsSenderAMQPConfigTargetName)
	v.SetDefault(keyLoginAttemptsSenderAMQPConfigConnectTimeout, config.DefaultAMQPSenderConnectTimeout)
	v.SetDefault(keyLoginAttemptsSenderAMQPConfigShutdownTimeout, config.DefaultAMQPSenderShutdownTimeout)
	v.SetDefault(keyLoginAttemptsSenderAMQPConfigPublishMaxTryAttempts, config.DefaultAMQPSenderPublishMaxTryAttempts)
	v.SetDefault(keyLoginAttemptsSenderAMQPConfigPublishBaseRetryDelay, config.DefaultAMQPSenderPublishBaseRetryDelay)
	v.SetDefault(keyLoginAttemptsSenderAMQPConfigPublishMaxRetryDelay, config.DefaultAMQPSenderPublishMaxRetryDelay)

	// kafka login attempts
	v.SetDefault(keyLoginAttemptsSenderKafkaConfigBrokers, config.DefaultKafkaBrokers)
	v.SetDefault(keyLoginAttemptsSenderKafkaConfigTargetName, defaultLoginAttemptsSenderKafkaConfigTargetName)
	v.SetDefault(keyLoginAttemptsSenderKafkaConfigConnectTimeout, config.DefaultKafkaSenderConnectTimeout)
	v.SetDefault(keyLoginAttemptsSenderKafkaConfigIdleTimeout, config.DefaultKafkaSenderIdleTimeout)
	v.SetDefault(keyLoginAttemptsSenderKafkaConfigShutdownTimeout, config.DefaultKafkaSenderShutdownTimeout)
	v.SetDefault(keyLoginAttemptsSenderKafkaConfigPublishMaxTryAttempts, config.DefaultKafkaSenderPublishMaxTryAttempts)
	v.SetDefault(keyLoginAttemptsSenderKafkaConfigPublishBaseRetryDelay, config.DefaultKafkaSenderPublishBaseRetryDelay)
	v.SetDefault(keyLoginAttemptsSenderKafkaConfigPublishMaxRetryDelay, config.DefaultKafkaSenderPublishMaxRetryDelay)
	v.SetDefault(keyLoginAttemptsSenderKafkaConfigBatchSize, config.DefaultKafkaSenderBatchSize)
	v.SetDefault(keyLoginAttemptsSenderKafkaConfigBatchBytes, config.DefaultKafkaSenderBatchBytes)
	v.SetDefault(keyLoginAttemptsSenderKafkaConfigBatchTimeout, config.DefaultKafkaSenderBatchTimeout)
	v.SetDefault(keyLoginAttemptsSenderKafkaConfigWriteTimeout, config.DefaultKafkaSenderWriteTimeout)
	v.SetDefault(keyLoginAttemptsSenderKafkaConfigRequiredAcks, config.DefaultKafkaSenderRequiredAcks)
}
