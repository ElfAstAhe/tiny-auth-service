package config

import (
	"log"
	"strings"

	libconf "github.com/ElfAstAhe/go-service-template/pkg/config"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	libextconf "github.com/ElfAstAhe/go-service-template/pkg/ext/config"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/spf13/viper"
)

const (
	// ApplicationName defines the global system identification token for the microservice node.
	ApplicationName string = "tiny-auth-service"
)

// Config aggregates multi-layered environment settings blocks, infrastructure configurations blueprints,
// cross-cutting telemetry parameters, and credentials profiles required to spin up application services.
type Config struct {
	App                       *AppConfig                   `mapstructure:"app" json:"app,omitempty" yaml:"app,omitempty"`
	Auth                      *libconf.AuthConfig          `mapstructure:"auth" json:"auth,omitempty" yaml:"auth,omitempty"`
	HTTP                      *libconf.HTTPConfig          `mapstructure:"http" json:"http,omitempty" yaml:"http,omitempty"`
	GRPC                      *libconf.GRPCConfig          `mapstructure:"grpc" json:"grpc,omitempty" yaml:"grpc,omitempty"`
	Log                       *libconf.LogConfig           `mapstructure:"log" json:"log,omitempty" yaml:"log,omitempty"`
	DB                        *libconf.DBConfig            `mapstructure:"db" json:"db,omitempty" yaml:"db,omitempty"`
	Telemetry                 *libconf.TelemetryConfig     `mapstructure:"telemetry" json:"telemetry,omitempty" yaml:"telemetry,omitempty"`
	LoginAttemptsConnector    *libconf.AMQPConnectorConfig `mapstructure:"login_attempts_connector" json:"login_attempts_connector,omitempty" yaml:"login_attempts_connector,omitempty"`
	LoginAttemptsSender       *LoginAttemptsConfig         `mapstructure:"login_attempts_sender" json:"login_attempts_sender,omitempty" yaml:"login_attempts_sender,omitempty"`
	DataAuditConnector        *libconf.AMQPConnectorConfig `mapstructure:"data_audit_connector" json:"data_audit_connector,omitempty" yaml:"data_audit_connector,omitempty"`
	DataAuditSender           *DataAuditConfig             `mapstructure:"data_audit_sender" json:"data_audit_sender,omitempty" yaml:"data_audit_sender,omitempty"`
	DataAuditOutboxTailCutter *libextconf.TailCutterConfig `mapstructure:"data_audit_outbox_tail_cutter" json:"data_audit_outbox_tail_cutter,omitempty" yaml:"data_audit_outbox_tail_cutter,omitempty"`
	AuthAuditOutboxTailCutter *libextconf.TailCutterConfig `mapstructure:"auth_audit_outbox_tail_cutter" json:"auth_audit_outbox_tail_cutter,omitempty" yaml:"auth_audit_outbox_tail_cutter,omitempty"`
}

// Linker params injected via compiler build automation pipelines (-ldflags).
var (
	AppVersion   string
	AppBuildTime string
)

// NewConfig acts as a complete constructor function mapping parameters across individual configuration scopes.
func NewConfig(
	app *AppConfig,
	auth *libconf.AuthConfig,
	HTTP *libconf.HTTPConfig,
	GRPC *libconf.GRPCConfig,
	log *libconf.LogConfig,
	db *libconf.DBConfig,
	telemetry *libconf.TelemetryConfig,
	loginAttemptsConnector *libconf.AMQPConnectorConfig,
	loginAttemptsSender *LoginAttemptsConfig,
	dataAuditConnector *libconf.AMQPConnectorConfig,
	dataAuditSender *DataAuditConfig,
	dataAuditOutboxTailCutter *libextconf.TailCutterConfig,
	authAuditOutboxTailCutter *libextconf.TailCutterConfig,
) *Config {
	return &Config{
		App:                       app,
		Auth:                      auth,
		HTTP:                      HTTP,
		GRPC:                      GRPC,
		Log:                       log,
		DB:                        db,
		Telemetry:                 telemetry,
		LoginAttemptsConnector:    loginAttemptsConnector,
		LoginAttemptsSender:       loginAttemptsSender,
		DataAuditConnector:        dataAuditConnector,
		DataAuditSender:           dataAuditSender,
		DataAuditOutboxTailCutter: dataAuditOutboxTailCutter,
		AuthAuditOutboxTailCutter: authAuditOutboxTailCutter,
	}
}

// NewDefaultConfig initializes the application settings schema pre-populating safe fallbacks from frameworks defaults.
func NewDefaultConfig() *Config {
	return NewConfig(
		NewDefaultAppConfig(),
		libconf.NewDefaultAuthConfig(),
		libconf.NewDefaultHTTPConfig(),
		libconf.NewDefaultGRPCConfig(),
		libconf.NewDefaultLogConfig(),
		libconf.NewDefaultDBConfig(),
		libconf.NewDefaultTelemetryConfig(),
		libconf.NewDefaultAMQPConnectorConfig(),
		NewDefaultLoginAttemptsConfig(),
		libconf.NewDefaultAMQPConnectorConfig(),
		NewDefaultDataAuditConfig(),
		libextconf.NewDefaultTailCutterConfig(),
		libextconf.NewDefaultTailCutterConfig(),
	)
}

// NewEmptyConfig builds a zero-allocated nested reference tree mapping empty configurations stubs.
//
//goland:noinspection GoUnusedExportedFunction
func NewEmptyConfig() *Config {
	return &Config{
		App:                    &AppConfig{},
		Auth:                   &libconf.AuthConfig{},
		HTTP:                   &libconf.HTTPConfig{},
		GRPC:                   &libconf.GRPCConfig{},
		Log:                    &libconf.LogConfig{},
		DB:                     &libconf.DBConfig{},
		Telemetry:              &libconf.TelemetryConfig{},
		LoginAttemptsConnector: &libconf.AMQPConnectorConfig{},
		LoginAttemptsSender: &LoginAttemptsConfig{
			AMQPConfig:  &libconf.AMQPSenderConfig{},
			KafkaConfig: &libconf.KafkaSenderConfig{},
		},
		DataAuditConnector: &libconf.AMQPConnectorConfig{},
		DataAuditSender: &DataAuditConfig{
			AMQPConfig:  &libconf.AMQPSenderConfig{},
			KafkaConfig: &libconf.KafkaSenderConfig{},
		},
		DataAuditOutboxTailCutter: &libextconf.TailCutterConfig{
			Dispatcher: &libconf.SchedulerDispatcherConfig{},
		},
		AuthAuditOutboxTailCutter: &libextconf.TailCutterConfig{
			Dispatcher: &libconf.SchedulerDispatcherConfig{},
		},
	}
}

// Validate dynamically maps fields onto an anonymous validation contract array loop to execute fast composite assertions.
func (c *Config) Validate() error {
	validators := []interface {
		Validate() error
	}{
		c.App,
		c.Auth,
		c.HTTP,
		c.GRPC,
		c.Log,
		c.DB,
		c.Telemetry,
		c.LoginAttemptsConnector,
		c.LoginAttemptsSender,
		c.DataAuditConnector,
		c.DataAuditSender,
		c.DataAuditOutboxTailCutter,
		c.AuthAuditOutboxTailCutter,
	}

	for _, validator := range validators {
		if utils.IsNil(validator) {
			return errs.NewConfigValidateError("validator item", "validator", "is required", nil)
		}
		if err := validator.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// Load собирает конфигурацию из: Flags -> ENV -> YAML -> Defaults
func Load() (*Config, error) {
	v := viper.New()

	// 1. Установка значений по умолчанию (Defaults)
	applyDefaults(v)

	// 2. Настройка Флагов (pflag для длинных имен --port)
	pFlagSet, err := initFLags()
	if err != nil {
		return nil, errs.NewConfigError("failed init flags", err)
	}

	// 3. Привязываем все флаги к Viper
	if err := bindFlags(pFlagSet, v); err != nil {
		return nil, err
	}

	// 4. Настройка Переменных окружения (ENV)
	// Используем твой механизм AutomaticEnv
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()

	// 5. Поддержка ENV для пути к конфигу
	err = v.BindEnv(FlagConfig, EnvConfig)
	if err != nil {
		return nil, errs.NewConfigError("failed to bind env", err)
	}

	// 6. Чтение файла конфигурации
	cfgFile := v.GetString(FlagConfig)
	v.SetConfigFile(cfgFile)

	if err := v.ReadInConfig(); err != nil {
		// Если файла нет — это предупреждение, но не фатальная ошибка (могут быть ENV)
		log.Printf("WARN: config file not found [%s]\n", cfgFile)
	}

	// 7. Маппинг в структуру
	var cfg = NewDefaultConfig()
	if err := v.Unmarshal(cfg); err != nil {
		return nil, errs.NewConfigError("failed to unmarshal config struct", err)
	}

	// 8. Итоговая валидация по всем слоям
	if err := cfg.Validate(); err != nil {
		return nil, errs.NewConfigError("config validation failed", err)
	}

	return cfg, nil
}
