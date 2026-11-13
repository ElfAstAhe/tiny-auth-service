package config

import (
	"fmt"
	"os"

	"github.com/ElfAstAhe/go-service-template/pkg/config"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
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
		res.String(config.FlagAppEnv, string(defaultAppEnv), "application environment")
		res.Duration(config.FlagAppInitTimeout, config.DefaultAppInitTimeout, "application init timeout")
		res.Duration(config.FlagAppStopTimeout, config.DefaultAppStopTimeout, "application stop timeout")
		res.Duration(config.FlagAppCloseTimeout, config.DefaultAppCloseTimeout, "application close timeout")
		res.String(FlagAppNodeName, defaultAppNodeName, "application node name")
		res.Int(FlagAppMaxListLimit, usecase.DefaultMaxLimit, "max list limit")
		res.String(FlagAppTokenIssuer, defaultTokenIssuer, "token issuer")
		res.String(FlagAppCipherKey, "", "cipher key")
		res.Duration(FlagAppDefShutdownTimeout, defaultDefShutdownTimeout, "default shutdown timeout")

		// Auth
		res.String(FlagAuthJWTSecret, "", "JWT secret")
		res.String(FlagAuthJWTSigningMethod, config.DefaultAuthSigningMethod, "JWT signing method")
		res.Duration(FlagAuthAccessTokenTTL, config.DefaultAuthAccessTokenTTL, "JWT access token TTL")
		res.Duration(FlagAuthRefreshTokenTTL, config.DefaultAuthRefreshTokenTTL, "JWT refresh token TTL")
		res.String(FlagAuthRSAPrivateKeyPath, "", "RSA private key path")
		res.String(FlagAuthMasterPasswordSalt, "", "master password salt")

		// HTTP
		res.String(FlagHTTPAddress, config.DefaultHTTPAddress, "http address")
		res.Duration(FlagHTTPReadTimeout, config.DefaultHTTPReadTimeout, "http read timeout")
		res.Duration(FlagHTTPWriteTimeout, config.DefaultHTTPWriteTimeout, "http write timeout")
		res.Duration(FlagHTTPIdleTimeout, config.DefaultHTTPIdleTimeout, "http idle timeout")
		res.Duration(FlagHTTPShutdownTimeout, config.DefaultHTTPShutdownTimeout, "http shutdown timeout")
		res.String(FlagHTTPPrivateKeyPath, "", "http private key path")
		res.String(FlagHTTPCertificatePath, "", "http certificate path")
		res.Bool(FlagHTTPSecure, config.DefaultHTTPSecure, "http secure mode")
		res.Int(FlagHTTPMaxRequestBodySize, config.DefaultHTTPMaxRequestBodySize, "http max request body size")

		// gRPC
		res.String(FlagGRPCAddress, config.DefaultGRPCAddress, "gRPC address")
		res.Duration(FlagGRPCMaxConnIdle, config.DefaultGRPCMaxConnIdle, "gRPC max connection idle timeout")
		res.Duration(FlagGRPCMaxConnAge, config.DefaultGRPCMaxConnAge, "gRPC max connection age timeout")
		res.Duration(FlagGRPCMaxConnAgeGrace, config.DefaultGRPCMaxConnAgeGrace, "gRPC max connection age grace timeout")
		res.Duration(FlagGRPCTimeout, config.DefaultGRPCTimeout, "gRPC timeout")
		res.Duration(FlagGRPCKeepAliveTime, config.DefaultGRPCKeepAliveTime, "gRPC keep alive timeout")
		res.Duration(FlagGRPCKeepAliveTimeout, config.DefaultGRPCKeepAliveTimeout, "gRPC keep alive timeout")
		res.Duration(FlagGRPCShutdownTimeout, config.DefaultGRPCShutdownTimeout, "gRPC shutdown timeout")

		// DB
		res.String(FlagDBDSN, defaultDBDSN, "database dsn")
		res.String(FlagDBDriver, defaultDBDriver, "database driver name/alias")
		res.Int(FlagDBMaxOpenConns, config.DefaultDBMaxOpenConns, "db max open connections")
		res.Int(FlagDBMaxIdleConns, config.DefaultDBMaxIdleConns, "db max idle connections")
		res.Duration(FlagDBMaxIdleLifetime, config.DefaultDBConnMaxIdleLifetime, "db max idle connection lifetime")
		res.Duration(FlagDBConnTimeout, config.DefaultDBConnTimeout, "db connection timeout)")

		// Log
		res.String(FlagLogLevel, config.DefaultLogLevel, "log level")
		res.String(FlagLogFormat, config.DefaultLogFormat, "log format")

		// Telemetry
		res.Bool(FlagTelemetryEnabled, config.DefaultTelemetryEnabled, "telemetry enabled")
		res.String(FlagTelemetryServiceName, "", "telemetry service name")
		res.String(FlagTelemetryExporterEndpoint, config.DefaultTelemetryExporterEndpoint, "telemetry exporter endpoint")
		res.Float64(FlagTelemetrySampleRate, config.DefaultTelemetrySampleRate, "telemetry sample rate")
		res.Duration(FlagTelemetryTimeout, config.DefaultTelemetryTimeout, "telemetry timeout")

		// login attempts sender amqp connector
		res.String(FlagLoginAttemptsSenderAMQPConnectorURL, config.DefaultAMQPConnectorURL, "login attempts sender connector AMQP server URL")
		res.String(FlagLoginAttemptsSenderAMQPConnectorUsername, defaultLoginAttemptsSenderAMQPConnectorUsername, "login attempts sender connector AMQP server username")
		res.String(FlagLoginAttemptsSenderAMQPConnectorPassword, defaultLoginAttemptsSenderAMQPConnectorPassword, "login attempts sender connector AMQP server password")
		res.Duration(FlagLoginAttemptsSenderAMQPConnectorConnectTimeout, config.DefaultAMQPConnectorConnectTimeout, "login attempts sender connector AMQP connect timeout")
		res.Duration(FlagLoginAttemptsSenderAMQPConnectorIdleTimeout, config.DefaultAMQPConnectorIdleTimeout, "login attempts sender connector AMQP idle timeout")
		res.Duration(FlagLoginAttemptsSenderAMQPConnectorWriteTimeout, config.DefaultAMQPConnectorWriteTimeout, "login attempts sender connector AMQP write timeout")
		res.Duration(FlagLoginAttemptsSenderAMQPConnectorShutdownTimeout, config.DefaultAMQPConnectorShutdownTimeout, "login attempts sender connector AMQP shutdown timeout")

		// login attempts sender
		res.String(FlagLoginAttemptsSenderKind, defaultLoginAttemptsSenderKind, "login attempts sender kind (accepted values: amqp, kafka)")
		res.Duration(FlagLoginAttemptsSenderNotifyTimeout, defaultLoginAttemptsSenderNotifyTimeout, "login attempts sender notify timeout")

		// login attempts sender amqp
		res.String(FlagLoginAttemptsSenderAMQPTargetName, defaultLoginAttemptsSenderAMQPTargetName, "login attempts sender AMQP queue/topic name")
		res.Duration(FlagLoginAttemptsSenderAMQPConnectTimeout, config.DefaultAMQPSenderConnectTimeout, "login attempts sender AMQP connect timeout")
		res.Duration(FlagLoginAttemptsSenderAMQPShutdownTimeout, config.DefaultAMQPSenderShutdownTimeout, "login attempts sender AMQP shutdown timeout")
		res.Int(FlagLoginAttemptsSenderAMQPPublishMaxTryAttempts, config.DefaultAMQPSenderPublishMaxTryAttempts, "login attempts sender AMQP max send tries")
		res.Duration(FlagLoginAttemptsSenderAMQPPublishBaseRetryDelay, config.DefaultAMQPSenderPublishBaseRetryDelay, "login attempts sender AMQP publish base retry delay")
		res.Duration(FlagLoginAttemptsSenderAMQPPublishMaxRetryDelay, config.DefaultAMQPSenderPublishMaxRetryDelay, "login attempts sender AMQP publish max retry delay")

		// login attempts sender kafka
		res.StringSlice(FlagLoginAttemptsSenderKafkaBrokers, config.DefaultKafkaBrokers, "login attempts sender kafka brokers, separated by comma, like: localhost:9092,localhost:9093")
		res.String(FlagLoginAttemptsSenderKafkaTargetName, "", "login attempts sender kafka queue/topic name")
		res.Duration(FlagLoginAttemptsSenderKafkaConnectTimeout, config.DefaultKafkaSenderConnectTimeout, "login attempts sender kafka connect timeout")
		res.Duration(FlagLoginAttemptsSenderKafkaIdleTimeout, config.DefaultKafkaSenderIdleTimeout, "login attempts sender kafka idle timeout")
		res.Duration(FlagLoginAttemptsSenderKafkaShutdownTimeout, config.DefaultKafkaSenderShutdownTimeout, "login attempts sender kafka shutdown timeout")
		res.Int(FlagLoginAttemptsSenderKafkaPublishMaxTryAttempts, config.DefaultKafkaSenderPublishMaxTryAttempts, "login attempts sender kafka max try attempts")
		res.Duration(FlagLoginAttemptsSenderKafkaPublishMaxRetryDelay, config.DefaultKafkaSenderPublishMaxRetryDelay, "login attempts sender kafka max retry delay")
		res.Duration(FlagLoginAttemptsSenderKafkaPublishBaseRetryDelay, config.DefaultKafkaSenderPublishBaseRetryDelay, "login attempts sender kafka publish base retry delay")
		res.String(FlagLoginAttemptsSenderKafkaUsername, "", "login attempts sender kafka username")
		res.String(FlagLoginAttemptsSenderKafkaPassword, "", "login attempts sender kafka password")
		res.Int(FlagLoginAttemptsSenderKafkaBatchSize, config.DefaultKafkaSenderBatchSize, "login attempts sender kafka batch size")
		res.Int(FlagLoginAttemptsSenderKafkaBatchBytes, config.DefaultKafkaSenderBatchBytes, "login attempts sender kafka batch bytes")
		res.Duration(FlagLoginAttemptsSenderKafkaBatchTimeout, config.DefaultKafkaSenderBatchTimeout, "login attempts sender kafka batch timeout")
		res.Duration(FlagLoginAttemptsSenderKafkaWriteTimeout, config.DefaultKafkaSenderWriteTimeout, "login attempts sender kafka write timeout")
		res.Int(FlagLoginAttemptsSenderKafkaRequiredAcks, config.DefaultKafkaSenderRequiredAcks, "login attempts sender kafka required acks")

		// data audit sender amqp connector
		res.String(FlagDataAuditSenderAMQPConnectorURL, config.DefaultAMQPConnectorURL, "data audit sender connector AMQP server URL")
		res.String(FlagDataAuditSenderAMQPConnectorUsername, defaultDataAuditSenderAMQPConnectorUsername, "data audit sender connector AMQP server username")
		res.String(FlagDataAuditSenderAMQPConnectorPassword, defaultDataAuditSenderAMQPConnectorPassword, "data audit sender connector AMQP server password")
		res.Duration(FlagDataAuditSenderAMQPConnectorConnectTimeout, config.DefaultAMQPConnectorConnectTimeout, "data audit sender connector AMQP connect timeout")
		res.Duration(FlagDataAuditSenderAMQPConnectorIdleTimeout, config.DefaultAMQPConnectorIdleTimeout, "data audit sender connector AMQP idle timeout")
		res.Duration(FlagDataAuditSenderAMQPConnectorWriteTimeout, config.DefaultAMQPConnectorWriteTimeout, "data audit sender connector AMQP write timeout")
		res.Duration(FlagDataAuditSenderAMQPConnectorShutdownTimeout, config.DefaultAMQPConnectorShutdownTimeout, "data audit sender connector AMQP shutdown timeout")

		// data audit sender
		res.String(FlagDataAuditSenderKind, defaultDataAuditSenderKind, "data audit sender kind (accepted values: amqp, kafka)")
		res.Duration(FlagDataAuditSenderNotifyTimeout, defaultDataAuditSenderNotifyTimeout, "data audit sender notify timeout")

		// data audit sender amqp
		res.String(FlagDataAuditSenderAMQPTargetName, defaultDataAuditSenderAMQPTargetName, "data audit sender AMQP queue/topic name")
		res.Duration(FlagDataAuditSenderAMQPConnectTimeout, config.DefaultAMQPSenderConnectTimeout, "data audit sender AMQP connect timeout")
		res.Duration(FlagDataAuditSenderAMQPShutdownTimeout, config.DefaultAMQPSenderShutdownTimeout, "data audit sender AMQP shutdown timeout")
		res.Int(FlagDataAuditSenderAMQPPublishMaxTryAttempts, config.DefaultAMQPSenderPublishMaxTryAttempts, "data audit sender AMQP max send tries")
		res.Duration(FlagDataAuditSenderAMQPPublishBaseRetryDelay, config.DefaultAMQPSenderPublishBaseRetryDelay, "data audit sender AMQP publish base retry delay")
		res.Duration(FlagDataAuditSenderAMQPPublishMaxRetryDelay, config.DefaultAMQPSenderPublishMaxRetryDelay, "data audit sender AMQP publish max retry delay")

		// data audit sender kafka
		res.StringSlice(FlagDataAuditSenderKafkaBrokers, config.DefaultKafkaBrokers, "data audit sender kafka brokers, separated by comma, like: localhost:9092,localhost:9093")
		res.String(FlagDataAuditSenderKafkaTargetName, "", "data audit sender kafka queue/topic name")
		res.Duration(FlagDataAuditSenderKafkaConnectTimeout, config.DefaultKafkaSenderConnectTimeout, "data audit sender kafka connect timeout")
		res.Duration(FlagDataAuditSenderKafkaIdleTimeout, config.DefaultKafkaSenderIdleTimeout, "data audit sender kafka idle timeout")
		res.Duration(FlagDataAuditSenderKafkaShutdownTimeout, config.DefaultKafkaSenderShutdownTimeout, "data audit sender kafka shutdown timeout")
		res.Int(FlagDataAuditSenderKafkaPublishMaxTryAttempts, config.DefaultKafkaSenderPublishMaxTryAttempts, "data audit sender kafka max try attempts")
		res.Duration(FlagDataAuditSenderKafkaPublishMaxRetryDelay, config.DefaultKafkaSenderPublishMaxRetryDelay, "data audit sender kafka max retry delay")
		res.Duration(FlagDataAuditSenderKafkaPublishBaseRetryDelay, config.DefaultKafkaSenderPublishBaseRetryDelay, "data audit sender kafka publish base retry delay")
		res.String(FlagDataAuditSenderKafkaUsername, "", "data audit sender kafka username")
		res.String(FlagDataAuditSenderKafkaPassword, "", "data audit sender kafka password")
		res.Int(FlagDataAuditSenderKafkaBatchSize, config.DefaultKafkaSenderBatchSize, "data audit sender kafka batch size")
		res.Int(FlagDataAuditSenderKafkaBatchBytes, config.DefaultKafkaSenderBatchBytes, "data audit sender kafka batch bytes")
		res.Duration(FlagDataAuditSenderKafkaBatchTimeout, config.DefaultKafkaSenderBatchTimeout, "data audit sender kafka batch timeout")
		res.Duration(FlagDataAuditSenderKafkaWriteTimeout, config.DefaultKafkaSenderWriteTimeout, "data audit sender kafka write timeout")
		res.Int(FlagDataAuditSenderKafkaRequiredAcks, config.DefaultKafkaSenderRequiredAcks, "data audit sender kafka required acks")
	}

	// Парсинг
	err = res.Parse(os.Args[1:])
	if err != nil {
		return nil, err
	}

	return res, err
}
