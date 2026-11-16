package config

import (
	"fmt"
	"os"

	libconf "github.com/ElfAstAhe/go-service-template/pkg/config"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/ext/config"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
	"github.com/spf13/pflag"
)

func initFLags() (res *pflag.FlagSet, err error) {
	defer func() {
		if r := recover(); r != nil {
			// Проверяем, является ли r ошибкой
			recoveryErr, ok := r.(error)
			if !ok {
				// Если это строка или что-то другое, приводим к виду error вручную
				recoveryErr = errs.NewConfigError(fmt.Sprintf("panic [%v] recovery", r), nil)
			}
			res = nil
			err = errs.NewConfigError("parse cli flags panic", recoveryErr)
		}
	}()

	res = pflag.NewFlagSet("cmd flags", pflag.PanicOnError)

	// Используем константы Flag...
	{
		// app
		res.String(FlagConfig, "config/config.yaml", "path to config file")
		res.String(libconf.FlagAppEnv, string(defaultAppEnv), "application environment")
		res.Duration(libconf.FlagAppInitTimeout, libconf.DefaultAppInitTimeout, "application init timeout")
		res.Duration(libconf.FlagAppStopTimeout, libconf.DefaultAppStopTimeout, "application stop timeout")
		res.Duration(libconf.FlagAppCloseTimeout, libconf.DefaultAppCloseTimeout, "application close timeout")
		res.String(FlagAppNodeName, defaultAppNodeName, "application node name")
		res.Int(FlagAppMaxListLimit, usecase.DefaultMaxLimit, "max list limit")
		res.String(FlagAppTokenIssuer, defaultTokenIssuer, "token issuer")
		res.String(FlagAppCipherKey, "", "cipher key")
		res.Duration(FlagAppDefShutdownTimeout, defaultDefShutdownTimeout, "default shutdown timeout")

		// Auth
		res.String(FlagAuthJWTSecret, "", "JWT secret")
		res.String(FlagAuthJWTSigningMethod, libconf.DefaultAuthSigningMethod, "JWT signing method")
		res.Duration(FlagAuthAccessTokenTTL, libconf.DefaultAuthAccessTokenTTL, "JWT access token TTL")
		res.Duration(FlagAuthRefreshTokenTTL, libconf.DefaultAuthRefreshTokenTTL, "JWT refresh token TTL")
		res.String(FlagAuthRSAPrivateKeyPath, "", "RSA private key path")
		res.String(FlagAuthMasterPasswordSalt, "", "master password salt")

		// HTTP
		res.String(FlagHTTPAddress, libconf.DefaultHTTPAddress, "http address")
		res.Duration(FlagHTTPReadTimeout, libconf.DefaultHTTPReadTimeout, "http read timeout")
		res.Duration(FlagHTTPWriteTimeout, libconf.DefaultHTTPWriteTimeout, "http write timeout")
		res.Duration(FlagHTTPIdleTimeout, libconf.DefaultHTTPIdleTimeout, "http idle timeout")
		res.Duration(FlagHTTPShutdownTimeout, libconf.DefaultHTTPShutdownTimeout, "http shutdown timeout")
		res.String(FlagHTTPPrivateKeyPath, "", "http private key path")
		res.String(FlagHTTPCertificatePath, "", "http certificate path")
		res.Bool(FlagHTTPSecure, libconf.DefaultHTTPSecure, "http secure mode")
		res.Int(FlagHTTPMaxRequestBodySize, libconf.DefaultHTTPMaxRequestBodySize, "http max request body size")

		// gRPC
		res.String(FlagGRPCAddress, libconf.DefaultGRPCAddress, "gRPC address")
		res.Duration(FlagGRPCMaxConnIdle, libconf.DefaultGRPCMaxConnIdle, "gRPC max connection idle timeout")
		res.Duration(FlagGRPCMaxConnAge, libconf.DefaultGRPCMaxConnAge, "gRPC max connection age timeout")
		res.Duration(FlagGRPCMaxConnAgeGrace, libconf.DefaultGRPCMaxConnAgeGrace, "gRPC max connection age grace timeout")
		res.Duration(FlagGRPCTimeout, libconf.DefaultGRPCTimeout, "gRPC timeout")
		res.Duration(FlagGRPCKeepAliveTime, libconf.DefaultGRPCKeepAliveTime, "gRPC keep alive timeout")
		res.Duration(FlagGRPCKeepAliveTimeout, libconf.DefaultGRPCKeepAliveTimeout, "gRPC keep alive timeout")
		res.Duration(FlagGRPCShutdownTimeout, libconf.DefaultGRPCShutdownTimeout, "gRPC shutdown timeout")

		// DB
		res.String(FlagDBDSN, defaultDBDSN, "database dsn")
		res.String(FlagDBDriver, defaultDBDriver, "database driver name/alias")
		res.Int(FlagDBMaxOpenConns, libconf.DefaultDBMaxOpenConns, "db max open connections")
		res.Int(FlagDBMaxIdleConns, libconf.DefaultDBMaxIdleConns, "db max idle connections")
		res.Duration(FlagDBMaxIdleLifetime, libconf.DefaultDBConnMaxIdleLifetime, "db max idle connection lifetime")
		res.Duration(FlagDBConnTimeout, libconf.DefaultDBConnTimeout, "db connection timeout)")

		// Log
		res.String(FlagLogLevel, libconf.DefaultLogLevel, "log level")
		res.String(FlagLogFormat, libconf.DefaultLogFormat, "log format")

		// Telemetry
		res.Bool(FlagTelemetryEnabled, libconf.DefaultTelemetryEnabled, "telemetry enabled")
		res.String(FlagTelemetryServiceName, "", "telemetry service name")
		res.String(FlagTelemetryExporterEndpoint, libconf.DefaultTelemetryExporterEndpoint, "telemetry exporter endpoint")
		res.Float64(FlagTelemetrySampleRate, libconf.DefaultTelemetrySampleRate, "telemetry sample rate")
		res.Duration(FlagTelemetryTimeout, libconf.DefaultTelemetryTimeout, "telemetry timeout")

		// login attempts sender amqp connector
		res.String(FlagLoginAttemptsSenderAMQPConnectorURL, libconf.DefaultAMQPConnectorURL, "login attempts sender connector AMQP server URL")
		res.String(FlagLoginAttemptsSenderAMQPConnectorUsername, defaultLoginAttemptsSenderAMQPConnectorUsername, "login attempts sender connector AMQP server username")
		res.String(FlagLoginAttemptsSenderAMQPConnectorPassword, defaultLoginAttemptsSenderAMQPConnectorPassword, "login attempts sender connector AMQP server password")
		res.Duration(FlagLoginAttemptsSenderAMQPConnectorConnectTimeout, libconf.DefaultAMQPConnectorConnectTimeout, "login attempts sender connector AMQP connect timeout")
		res.Duration(FlagLoginAttemptsSenderAMQPConnectorIdleTimeout, libconf.DefaultAMQPConnectorIdleTimeout, "login attempts sender connector AMQP idle timeout")
		res.Duration(FlagLoginAttemptsSenderAMQPConnectorWriteTimeout, libconf.DefaultAMQPConnectorWriteTimeout, "login attempts sender connector AMQP write timeout")
		res.Duration(FlagLoginAttemptsSenderAMQPConnectorShutdownTimeout, libconf.DefaultAMQPConnectorShutdownTimeout, "login attempts sender connector AMQP shutdown timeout")

		// login attempts sender
		res.String(FlagLoginAttemptsSenderKind, defaultLoginAttemptsSenderKind, "login attempts sender kind (accepted values: amqp, kafka)")
		res.Duration(FlagLoginAttemptsSenderNotifyTimeout, defaultLoginAttemptsSenderNotifyTimeout, "login attempts sender notify timeout")

		// login attempts sender amqp
		res.String(FlagLoginAttemptsSenderAMQPTargetName, defaultLoginAttemptsSenderAMQPTargetName, "login attempts sender AMQP queue/topic name")
		res.Duration(FlagLoginAttemptsSenderAMQPConnectTimeout, libconf.DefaultAMQPSenderConnectTimeout, "login attempts sender AMQP connect timeout")
		res.Duration(FlagLoginAttemptsSenderAMQPShutdownTimeout, libconf.DefaultAMQPSenderShutdownTimeout, "login attempts sender AMQP shutdown timeout")
		res.Int(FlagLoginAttemptsSenderAMQPPublishMaxTryAttempts, libconf.DefaultAMQPSenderPublishMaxTryAttempts, "login attempts sender AMQP max send tries")
		res.Duration(FlagLoginAttemptsSenderAMQPPublishBaseRetryDelay, libconf.DefaultAMQPSenderPublishBaseRetryDelay, "login attempts sender AMQP publish base retry delay")
		res.Duration(FlagLoginAttemptsSenderAMQPPublishMaxRetryDelay, libconf.DefaultAMQPSenderPublishMaxRetryDelay, "login attempts sender AMQP publish max retry delay")

		// login attempts sender kafka
		res.StringSlice(FlagLoginAttemptsSenderKafkaBrokers, libconf.DefaultKafkaBrokers, "login attempts sender kafka brokers, separated by comma, like: localhost:9092,localhost:9093")
		res.String(FlagLoginAttemptsSenderKafkaTargetName, "", "login attempts sender kafka queue/topic name")
		res.Duration(FlagLoginAttemptsSenderKafkaConnectTimeout, libconf.DefaultKafkaSenderConnectTimeout, "login attempts sender kafka connect timeout")
		res.Duration(FlagLoginAttemptsSenderKafkaIdleTimeout, libconf.DefaultKafkaSenderIdleTimeout, "login attempts sender kafka idle timeout")
		res.Duration(FlagLoginAttemptsSenderKafkaShutdownTimeout, libconf.DefaultKafkaSenderShutdownTimeout, "login attempts sender kafka shutdown timeout")
		res.Int(FlagLoginAttemptsSenderKafkaPublishMaxTryAttempts, libconf.DefaultKafkaSenderPublishMaxTryAttempts, "login attempts sender kafka max try attempts")
		res.Duration(FlagLoginAttemptsSenderKafkaPublishMaxRetryDelay, libconf.DefaultKafkaSenderPublishMaxRetryDelay, "login attempts sender kafka max retry delay")
		res.Duration(FlagLoginAttemptsSenderKafkaPublishBaseRetryDelay, libconf.DefaultKafkaSenderPublishBaseRetryDelay, "login attempts sender kafka publish base retry delay")
		res.String(FlagLoginAttemptsSenderKafkaUsername, "", "login attempts sender kafka username")
		res.String(FlagLoginAttemptsSenderKafkaPassword, "", "login attempts sender kafka password")
		res.Int(FlagLoginAttemptsSenderKafkaBatchSize, libconf.DefaultKafkaSenderBatchSize, "login attempts sender kafka batch size")
		res.Int(FlagLoginAttemptsSenderKafkaBatchBytes, libconf.DefaultKafkaSenderBatchBytes, "login attempts sender kafka batch bytes")
		res.Duration(FlagLoginAttemptsSenderKafkaBatchTimeout, libconf.DefaultKafkaSenderBatchTimeout, "login attempts sender kafka batch timeout")
		res.Duration(FlagLoginAttemptsSenderKafkaWriteTimeout, libconf.DefaultKafkaSenderWriteTimeout, "login attempts sender kafka write timeout")
		res.Int(FlagLoginAttemptsSenderKafkaRequiredAcks, libconf.DefaultKafkaSenderRequiredAcks, "login attempts sender kafka required acks")

		// data audit sender amqp connector
		res.String(FlagDataAuditSenderAMQPConnectorURL, libconf.DefaultAMQPConnectorURL, "data audit sender connector AMQP server URL")
		res.String(FlagDataAuditSenderAMQPConnectorUsername, defaultDataAuditSenderAMQPConnectorUsername, "data audit sender connector AMQP server username")
		res.String(FlagDataAuditSenderAMQPConnectorPassword, defaultDataAuditSenderAMQPConnectorPassword, "data audit sender connector AMQP server password")
		res.Duration(FlagDataAuditSenderAMQPConnectorConnectTimeout, libconf.DefaultAMQPConnectorConnectTimeout, "data audit sender connector AMQP connect timeout")
		res.Duration(FlagDataAuditSenderAMQPConnectorIdleTimeout, libconf.DefaultAMQPConnectorIdleTimeout, "data audit sender connector AMQP idle timeout")
		res.Duration(FlagDataAuditSenderAMQPConnectorWriteTimeout, libconf.DefaultAMQPConnectorWriteTimeout, "data audit sender connector AMQP write timeout")
		res.Duration(FlagDataAuditSenderAMQPConnectorShutdownTimeout, libconf.DefaultAMQPConnectorShutdownTimeout, "data audit sender connector AMQP shutdown timeout")

		// data audit sender
		res.String(FlagDataAuditSenderKind, defaultDataAuditSenderKind, "data audit sender kind (accepted values: amqp, kafka)")
		res.Duration(FlagDataAuditSenderNotifyTimeout, defaultDataAuditSenderNotifyTimeout, "data audit sender notify timeout")

		// data audit sender amqp
		res.String(FlagDataAuditSenderAMQPTargetName, defaultDataAuditSenderAMQPTargetName, "data audit sender AMQP queue/topic name")
		res.Duration(FlagDataAuditSenderAMQPConnectTimeout, libconf.DefaultAMQPSenderConnectTimeout, "data audit sender AMQP connect timeout")
		res.Duration(FlagDataAuditSenderAMQPShutdownTimeout, libconf.DefaultAMQPSenderShutdownTimeout, "data audit sender AMQP shutdown timeout")
		res.Int(FlagDataAuditSenderAMQPPublishMaxTryAttempts, libconf.DefaultAMQPSenderPublishMaxTryAttempts, "data audit sender AMQP max send tries")
		res.Duration(FlagDataAuditSenderAMQPPublishBaseRetryDelay, libconf.DefaultAMQPSenderPublishBaseRetryDelay, "data audit sender AMQP publish base retry delay")
		res.Duration(FlagDataAuditSenderAMQPPublishMaxRetryDelay, libconf.DefaultAMQPSenderPublishMaxRetryDelay, "data audit sender AMQP publish max retry delay")

		// data audit sender kafka
		res.StringSlice(FlagDataAuditSenderKafkaBrokers, libconf.DefaultKafkaBrokers, "data audit sender kafka brokers, separated by comma, like: localhost:9092,localhost:9093")
		res.String(FlagDataAuditSenderKafkaTargetName, "", "data audit sender kafka queue/topic name")
		res.Duration(FlagDataAuditSenderKafkaConnectTimeout, libconf.DefaultKafkaSenderConnectTimeout, "data audit sender kafka connect timeout")
		res.Duration(FlagDataAuditSenderKafkaIdleTimeout, libconf.DefaultKafkaSenderIdleTimeout, "data audit sender kafka idle timeout")
		res.Duration(FlagDataAuditSenderKafkaShutdownTimeout, libconf.DefaultKafkaSenderShutdownTimeout, "data audit sender kafka shutdown timeout")
		res.Int(FlagDataAuditSenderKafkaPublishMaxTryAttempts, libconf.DefaultKafkaSenderPublishMaxTryAttempts, "data audit sender kafka max try attempts")
		res.Duration(FlagDataAuditSenderKafkaPublishMaxRetryDelay, libconf.DefaultKafkaSenderPublishMaxRetryDelay, "data audit sender kafka max retry delay")
		res.Duration(FlagDataAuditSenderKafkaPublishBaseRetryDelay, libconf.DefaultKafkaSenderPublishBaseRetryDelay, "data audit sender kafka publish base retry delay")
		res.String(FlagDataAuditSenderKafkaUsername, "", "data audit sender kafka username")
		res.String(FlagDataAuditSenderKafkaPassword, "", "data audit sender kafka password")
		res.Int(FlagDataAuditSenderKafkaBatchSize, libconf.DefaultKafkaSenderBatchSize, "data audit sender kafka batch size")
		res.Int(FlagDataAuditSenderKafkaBatchBytes, libconf.DefaultKafkaSenderBatchBytes, "data audit sender kafka batch bytes")
		res.Duration(FlagDataAuditSenderKafkaBatchTimeout, libconf.DefaultKafkaSenderBatchTimeout, "data audit sender kafka batch timeout")
		res.Duration(FlagDataAuditSenderKafkaWriteTimeout, libconf.DefaultKafkaSenderWriteTimeout, "data audit sender kafka write timeout")
		res.Int(FlagDataAuditSenderKafkaRequiredAcks, libconf.DefaultKafkaSenderRequiredAcks, "data audit sender kafka required acks")

		// auth audit outbox tail cutter
		res.Int(FlagAuthAuditOutboxTCTailLimit, config.DefaultTailCutterTailLimit, "auth audit outbox tail cutter tail limit")
		res.Duration(FlagAuthAuditOutboxTCTailInterval, defaultAuthAuditOutboxTailCutterTailInterval, "auth audit outbox tail cutter tail interval")
		res.Bool(FlagAuthAuditOutboxTCTailCut, defaultAuthAuditOutboxTailCutterTailCut, "auth audit outbox tail cutter tail cut flag")
		res.Int(FlagAuthAuditOutboxTCWorkerCount, config.DefaultTailCutterWorkerCount, "auth audit outbox tail cutter worker count")
		res.Int(FlagAuthAuditOutboxTCDataCapacity, config.DefaultTailCutterDataCapacity, "auth audit outbox tail cutter data capacity")
		res.Bool(FlagAuthAuditOutboxTCCompleteProcess, config.DefaultTailCutterCompleteProcessing, "auth audit outbox tail cutter complete process flag")
		res.Duration(FlagAuthAuditOutboxTCStartInterval, config.DefaultTailCutterStartInterval, "auth audit outbox tail cutter start interval")
		res.Duration(FlagAuthAuditOutboxTCScheduleInterval, config.DefaultTailCutterScheduleInterval, "auth audit outbox tail cutter schedule interval")
		res.Duration(FlagAuthAuditOutboxTCStopTimeout, config.DefaultTailCutterStopTimeout, "auth audit outbox tail cutter stop timeout")

		// data audit outbox tail cutter
		res.Int(FlagDataAuditOutboxTCTailLimit, config.DefaultTailCutterTailLimit, "data audit outbox tail cutter tail limit")
		res.Duration(FlagDataAuditOutboxTCTailInterval, defaultDataAuditOutboxTailCutterTailInterval, "data audit outbox tail cutter tail interval")
		res.Bool(FlagDataAuditOutboxTCTailCut, defaultDataAuditOutboxTailCutterTailCut, "data audit outbox tail cutter tail cut flag")
		res.Int(FlagDataAuditOutboxTCWorkerCount, config.DefaultTailCutterWorkerCount, "data audit outbox tail cutter worker count")
		res.Int(FlagDataAuditOutboxTCDataCapacity, config.DefaultTailCutterDataCapacity, "data audit outbox tail cutter data capacity")
		res.Bool(FlagDataAuditOutboxTCCompleteProcess, config.DefaultTailCutterCompleteProcessing, "data audit outbox tail cutter complete process flag")
		res.Duration(FlagDataAuditOutboxTCStartInterval, config.DefaultTailCutterStartInterval, "data audit outbox tail cutter start interval")
		res.Duration(FlagDataAuditOutboxTCScheduleInterval, config.DefaultTailCutterScheduleInterval, "data audit outbox tail cutter schedule interval")
		res.Duration(FlagDataAuditOutboxTCStopTimeout, config.DefaultTailCutterStopTimeout, "data audit outbox tail cutter stop timeout")
	}

	// Парсинг
	err = res.Parse(os.Args[1:])
	if err != nil {
		return nil, err
	}

	return res, err
}
