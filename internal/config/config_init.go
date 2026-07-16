package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/ElfAstAhe/go-service-template/pkg/config"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
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
