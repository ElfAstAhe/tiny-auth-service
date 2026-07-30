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

		// svc-creds
		res.String(FlagCredsUsername, "", "client token refresher username")
		res.String(FlagCredsPassword, "", "client token refresher password")
		res.Duration(FlagCredsScheduleInterval, defaultCredsScheduleInterval, "client token refresher schedule interval")
		res.Duration(FlagCredsErrorScheduleInterval, defaultCredsErrorScheduleInterval, "client token refresher error schedule interval")

		// data-audit-client
		res.String(FlagDataAuditClientBaseURL, "", "data audit service base url")
		res.Duration(FlagDataAuditClientTimeout, defaultAuditClientTimeout, "data audit client timeout")
		res.Int(FlagDataAuditClientWorkerCount, defaultAuditClientWorkerCount, "data audit client worker count")
		res.Int(FlagDataAuditClientDataCapacity, defaultAuditClientDataCapacity, "data audit client data capacity")
		res.Bool(FlagDataAuditClientCompleteProcessing, true, "data audit client complete processing")
		res.Duration(FlagDataAuditClientShutdownTimeout, defaultAuditClientShutdownTimeout, "data audit client shutdown timeout")

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
		res.String(FlagDBDSN, config.DefaultDBDSN, "database dsn")
		res.String(FlagDBDriver, config.DefaultDBDriver, "database driver name/alias")
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

		// amqp connector
		res.String(FlagAMQPConnectorURL, config.DefaultAMQPConnectorURL, "connector AMQP server URL")
		res.String(FlagAMQPConnectorUsername, defaultAMQPConnectorUsername, "connector AMQP server username")
		res.String(FlagAMQPConnectorPassword, defaultAMQPConnectorPassword, "connector AMQP server password")
		res.Duration(FlagAMQPConnectorConnectTimeout, config.DefaultAMQPConnectorConnectTimeout, "connector AMQP connect timeout")
		res.Duration(FlagAMQPConnectorIdleTimeout, config.DefaultAMQPConnectorIdleTimeout, "connector AMQP idle timeout")
		res.Duration(FlagAMQPConnectorWriteTimeout, config.DefaultAMQPConnectorWriteTimeout, "connector AMQP write timeout")
		res.Duration(FlagAMQPConnectorShutdownTimeout, config.DefaultAMQPConnectorShutdownTimeout, "connector AMQP shutdown timeout")

		// amqp sender
		res.String(FlagLoginAttemptsSenderTargetName, defaultLoginAttemptsSenderTargetName, "login attempts sender queue/tipic name")
		res.Duration(FlagLoginAttemptsSenderConnectTimeout, config.DefaultAMQPSenderConnectTimeout, "login attempts sender connect timeout")
		res.Duration(FlagLoginAttemptsSenderNotifyTimeout, defaultLoginAttemptsNotifyTimeout, "login attempts sender notify timeout")
		res.Duration(FlagLoginAttemptsSenderShutdownTimeout, config.DefaultAMQPSenderShutdownTimeout, "login attempts sender shutdown timeout")
		res.Int(FlagLoginAttemptsSenderPublishMaxTryAttempts, config.DefaultAMQPSenderPublishMaxTryAttempts, "login attempts sender max send tries")
		res.Duration(FlagLoginAttemptsSenderPublishBaseRetryDelay, config.DefaultAMQPSenderPublishBaseRetryDelay, "login attempts sender publish base retry delay")
		res.Duration(FlagLoginAttemptsSenderPublishMaxRetryDelay, config.DefaultAMQPSenderPublishMaxRetryDelay, "login attempts sender publish max retry delay")
	}

	// Парсинг
	err = res.Parse(os.Args[1:])
	if err != nil {
		return nil, err
	}

	return res, err
}
