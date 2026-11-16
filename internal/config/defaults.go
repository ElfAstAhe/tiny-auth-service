package config

import (
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/config"
	libextconf "github.com/ElfAstAhe/go-service-template/pkg/ext/config"
	"github.com/spf13/viper"
)

// app
//
//nolint:gosec // G101 : app conf dev/test defaults
const (
	defaultAppEnv             config.AppEnv = config.AppEnvDevelopment
	defaultAppNodeName        string        = ApplicationName
	defaultMaxListLimit       int           = 100
	defaultTokenIssuer        string        = "tiny-auth-service"
	defaultDefShutdownTimeout time.Duration = 15 * time.Second
)

// creds
const (
	defaultTokenRefresherScheduleInterval      time.Duration = 5 * time.Minute
	defaultTokenRefresherErrorScheduleInterval time.Duration = 3 * time.Second
)

// db
const (
	defaultDBDriver string = "postgres"
	//nolint:gosec // G101 : dev/test env default params
	defaultDBDSN string = "postgres://svc_auth:password@localhost:5432/test?sslmode=disable&search_path=auth_db"
)

// data audit amqp connector
const (
	defaultDataAuditSenderAMQPConnectorUsername string = "svc-auth"
	defaultDataAuditSenderAMQPConnectorPassword string = "test"
)

// data audit connector
const (
	// defaultDataAuditSenderKind sender kind, valus: amqp, kafka
	defaultDataAuditSenderKind string = "amqp"
	// defaultDataAuditSenderNotifyTimeout
	defaultDataAuditSenderNotifyTimeout time.Duration = 2 * time.Second
)

// data audit amqp target name
const (
	defaultDataAuditSenderAMQPTargetName string = "tiny.audit::data.audit"
)

// data audit kafka target name
const (
	defaultDataAuditSenderKafkaTargetName string = "tiny.audit.data.audit"
)

// amqp connector
const (
	defaultLoginAttemptsSenderAMQPConnectorUsername string = "svc-auth"
	defaultLoginAttemptsSenderAMQPConnectorPassword string = "test"
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
	defaultLoginAttemptsSenderAMQPTargetName string = "tiny.auth::login.attempts"
)

// kafka login attempts target name
const (
	defaultLoginAttemptsSenderKafkaTargetName string = "tiny.auth.login.attempts"
)

// auth audit outbox tail cutter
const (
	defaultAuthAuditOutboxTailCutterTailInterval time.Duration = 7 * 24 * time.Hour
	defaultAuthAuditOutboxTailCutterTailCut      bool          = true
)

// data audit outbox tail cutter
const (
	defaultDataAuditOutboxTailCutterTailInterval time.Duration = 7 * 24 * time.Hour
	defaultDataAuditOutboxTailCutterTailCut      bool          = true
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

	// login attempts sender amqp connector
	v.SetDefault(keyLoginAttemptsSenderAMQPConnectorURL, config.DefaultAMQPConnectorURL)
	v.SetDefault(keyLoginAttemptsSenderAMQPConnectorUsername, defaultLoginAttemptsSenderAMQPConnectorUsername)
	v.SetDefault(keyLoginAttemptsSenderAMQPConnectorPassword, defaultLoginAttemptsSenderAMQPConnectorPassword)
	v.SetDefault(keyLoginAttemptsSenderAMQPConnectorConnectTimeout, config.DefaultAMQPSenderConnectTimeout)
	v.SetDefault(keyLoginAttemptsSenderAMQPConnectorWriteTimeout, config.DefaultAMQPConnectorWriteTimeout)
	v.SetDefault(keyLoginAttemptsSenderAMQPConnectorIdleTimeout, config.DefaultAMQPConnectorIdleTimeout)
	v.SetDefault(keyLoginAttemptsSenderAMQPConnectorShutdownTimeout, config.DefaultAMQPConnectorShutdownTimeout)

	// login attempts sender
	v.SetDefault(keyLoginAttemptsSenderKind, defaultLoginAttemptsSenderKind)
	v.SetDefault(keyLoginAttemptsSenderNotifyTimeout, defaultLoginAttemptsSenderNotifyTimeout)

	// login attempts sender amqp
	v.SetDefault(keyLoginAttemptsSenderAMQPTargetName, defaultLoginAttemptsSenderAMQPTargetName)
	v.SetDefault(keyLoginAttemptsSenderAMQPConnectTimeout, config.DefaultAMQPSenderConnectTimeout)
	v.SetDefault(keyLoginAttemptsSenderAMQPShutdownTimeout, config.DefaultAMQPSenderShutdownTimeout)
	v.SetDefault(keyLoginAttemptsSenderAMQPPublishMaxTryAttempts, config.DefaultAMQPSenderPublishMaxTryAttempts)
	v.SetDefault(keyLoginAttemptsSenderAMQPPublishBaseRetryDelay, config.DefaultAMQPSenderPublishBaseRetryDelay)
	v.SetDefault(keyLoginAttemptsSenderAMQPPublishMaxRetryDelay, config.DefaultAMQPSenderPublishMaxRetryDelay)

	// login attempts sender kafka
	v.SetDefault(keyLoginAttemptsSenderKafkaBrokers, config.DefaultKafkaBrokers)
	v.SetDefault(keyLoginAttemptsSenderKafkaTargetName, defaultLoginAttemptsSenderKafkaTargetName)
	v.SetDefault(keyLoginAttemptsSenderKafkaConnectTimeout, config.DefaultKafkaSenderConnectTimeout)
	v.SetDefault(keyLoginAttemptsSenderKafkaIdleTimeout, config.DefaultKafkaSenderIdleTimeout)
	v.SetDefault(keyLoginAttemptsSenderKafkaShutdownTimeout, config.DefaultKafkaSenderShutdownTimeout)
	v.SetDefault(keyLoginAttemptsSenderKafkaPublishMaxTryAttempts, config.DefaultKafkaSenderPublishMaxTryAttempts)
	v.SetDefault(keyLoginAttemptsSenderKafkaPublishBaseRetryDelay, config.DefaultKafkaSenderPublishBaseRetryDelay)
	v.SetDefault(keyLoginAttemptsSenderKafkaPublishMaxRetryDelay, config.DefaultKafkaSenderPublishMaxRetryDelay)
	v.SetDefault(keyLoginAttemptsSenderKafkaBatchSize, config.DefaultKafkaSenderBatchSize)
	v.SetDefault(keyLoginAttemptsSenderKafkaBatchBytes, config.DefaultKafkaSenderBatchBytes)
	v.SetDefault(keyLoginAttemptsSenderKafkaBatchTimeout, config.DefaultKafkaSenderBatchTimeout)
	v.SetDefault(keyLoginAttemptsSenderKafkaWriteTimeout, config.DefaultKafkaSenderWriteTimeout)
	v.SetDefault(keyLoginAttemptsSenderKafkaRequiredAcks, config.DefaultKafkaSenderRequiredAcks)

	// data audit amqp connector
	v.SetDefault(keyDataAuditSenderAMQPConnectorURL, config.DefaultAMQPConnectorURL)
	v.SetDefault(keyDataAuditSenderAMQPConnectorUsername, defaultDataAuditSenderAMQPConnectorUsername)
	v.SetDefault(keyDataAuditSenderAMQPConnectorPassword, defaultDataAuditSenderAMQPConnectorPassword)
	v.SetDefault(keyDataAuditSenderAMQPConnectorConnectTimeout, config.DefaultAMQPSenderConnectTimeout)
	v.SetDefault(keyDataAuditSenderAMQPConnectorWriteTimeout, config.DefaultAMQPConnectorWriteTimeout)
	v.SetDefault(keyDataAuditSenderAMQPConnectorIdleTimeout, config.DefaultAMQPConnectorIdleTimeout)
	v.SetDefault(keyDataAuditSenderAMQPConnectorShutdownTimeout, config.DefaultAMQPConnectorShutdownTimeout)

	// data audit sender
	v.SetDefault(keyDataAuditSenderKind, defaultDataAuditSenderKind)
	v.SetDefault(keyDataAuditSenderNotifyTimeout, defaultDataAuditSenderNotifyTimeout)

	// data audit sender amqp
	v.SetDefault(keyDataAuditSenderAMQPTargetName, defaultDataAuditSenderAMQPTargetName)
	v.SetDefault(keyDataAuditSenderAMQPConnectTimeout, config.DefaultAMQPSenderConnectTimeout)
	v.SetDefault(keyDataAuditSenderAMQPShutdownTimeout, config.DefaultAMQPSenderShutdownTimeout)
	v.SetDefault(keyDataAuditSenderAMQPPublishMaxTryAttempts, config.DefaultAMQPSenderPublishMaxTryAttempts)
	v.SetDefault(keyDataAuditSenderAMQPPublishBaseRetryDelay, config.DefaultAMQPSenderPublishBaseRetryDelay)
	v.SetDefault(keyDataAuditSenderAMQPPublishMaxRetryDelay, config.DefaultAMQPSenderPublishMaxRetryDelay)

	// data audit sender kafka
	v.SetDefault(keyDataAuditSenderKafkaBrokers, config.DefaultKafkaBrokers)
	v.SetDefault(keyDataAuditSenderKafkaTargetName, defaultDataAuditSenderKafkaTargetName)
	v.SetDefault(keyDataAuditSenderKafkaConnectTimeout, config.DefaultKafkaSenderConnectTimeout)
	v.SetDefault(keyDataAuditSenderKafkaIdleTimeout, config.DefaultKafkaSenderIdleTimeout)
	v.SetDefault(keyDataAuditSenderKafkaShutdownTimeout, config.DefaultKafkaSenderShutdownTimeout)
	v.SetDefault(keyDataAuditSenderKafkaPublishMaxTryAttempts, config.DefaultKafkaSenderPublishMaxTryAttempts)
	v.SetDefault(keyDataAuditSenderKafkaPublishBaseRetryDelay, config.DefaultKafkaSenderPublishBaseRetryDelay)
	v.SetDefault(keyDataAuditSenderKafkaPublishMaxRetryDelay, config.DefaultKafkaSenderPublishMaxRetryDelay)
	v.SetDefault(keyDataAuditSenderKafkaBatchSize, config.DefaultKafkaSenderBatchSize)
	v.SetDefault(keyDataAuditSenderKafkaBatchBytes, config.DefaultKafkaSenderBatchBytes)
	v.SetDefault(keyDataAuditSenderKafkaBatchTimeout, config.DefaultKafkaSenderBatchTimeout)
	v.SetDefault(keyDataAuditSenderKafkaWriteTimeout, config.DefaultKafkaSenderWriteTimeout)
	v.SetDefault(keyDataAuditSenderKafkaRequiredAcks, config.DefaultKafkaSenderRequiredAcks)

	// auth audit outbox tail cutter
	v.SetDefault(keyAuthAuditOutboxTCTailLimit, libextconf.DefaultTailCutterTailLimit)
	v.SetDefault(keyAuthAuditOutboxTCTailInterval, defaultAuthAuditOutboxTailCutterTailInterval)
	v.SetDefault(keyAuthAuditOutboxTCTailCut, defaultAuthAuditOutboxTailCutterTailCut)
	v.SetDefault(keyAuthAuditOutboxTCWorkerCount, libextconf.DefaultTailCutterWorkerCount)
	v.SetDefault(keyAuthAuditOutboxTCDataCapacity, libextconf.DefaultTailCutterDataCapacity)
	v.SetDefault(keyAuthAuditOutboxTCCompleteProcessing, libextconf.DefaultTailCutterCompleteProcessing)
	v.SetDefault(keyAuthAuditOutboxTCStartInterval, libextconf.DefaultTailCutterStartInterval)
	v.SetDefault(keyAuthAuditOutboxTCScheduleInterval, libextconf.DefaultTailCutterScheduleInterval)
	v.SetDefault(keyAuthAuditOutboxTCStopTimeout, libextconf.DefaultTailCutterStopTimeout)

	// data audit outbox tail cutter
	v.SetDefault(keyDataAuditOutboxTCTailLimit, libextconf.DefaultTailCutterTailLimit)
	v.SetDefault(keyDataAuditOutboxTCTailInterval, defaultDataAuditOutboxTailCutterTailInterval)
	v.SetDefault(keyDataAuditOutboxTCTailCut, defaultDataAuditOutboxTailCutterTailCut)
	v.SetDefault(keyDataAuditOutboxTCWorkerCount, libextconf.DefaultTailCutterWorkerCount)
	v.SetDefault(keyDataAuditOutboxTCDataCapacity, libextconf.DefaultTailCutterDataCapacity)
	v.SetDefault(keyDataAuditOutboxTCCompleteProcessing, libextconf.DefaultTailCutterCompleteProcessing)
	v.SetDefault(keyDataAuditOutboxTCStartInterval, libextconf.DefaultTailCutterStartInterval)
	v.SetDefault(keyDataAuditOutboxTCScheduleInterval, libextconf.DefaultTailCutterScheduleInterval)
	v.SetDefault(keyDataAuditOutboxTCStopTimeout, libextconf.DefaultTailCutterStopTimeout)
}
