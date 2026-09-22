package config

// FlagConfig - файл конфигурации
const FlagConfig = "config-path"

// App config flags
const (
	FlagAppNodeName           string = "app-node-name"
	FlagAppMaxListLimit       string = "app-max-list-limit"
	FlagAppTokenIssuer        string = "app-token-issuer"
	FlagAppCipherKey          string = "app-cipher-key"
	FlagAppDefShutdownTimeout string = "app-def-shutdown-timeout"
)

// Creds
const (
	FlagCredsUsername              string = "svc-creds-username"
	FlagCredsPassword              string = "svc-creds-password"
	FlagCredsScheduleInterval      string = "svc-creds-schedule-interval"
	FlagCredsErrorScheduleInterval string = "svc-creds-error-schedule-interval"
)

// Data-audit-client
const (
	FlagDataAuditClientBaseURL            string = "data-audit-client-base-url"
	FlagDataAuditClientTimeout            string = "data-audit-client-timeout"
	FlagDataAuditClientWorkerCount        string = "data-audit-client-worker-count"
	FlagDataAuditClientDataCapacity       string = "data-audit-client-data-capacity"
	FlagDataAuditClientCompleteProcessing string = "data-audit-client-complete-processing"
	FlagDataAuditClientShutdownTimeout    string = "data-audit-client-shutdown-timeout"
)

// Auth config flags
const (
	FlagAuthJWTSecret          string = "auth-jwt-secret"
	FlagAuthJWTSigningMethod   string = "auth-jwt-signing-method"
	FlagAuthAccessTokenTTL     string = "auth-access-token-ttl"
	FlagAuthRefreshTokenTTL    string = "auth-refresh-token-ttl"
	FlagAuthRSAPrivateKeyPath  string = "auth-rsa-private-key-path"
	FlagAuthMasterPasswordSalt string = "auth-master-password-salt"
)

// DB config flags
const (
	FlagDBDSN             string = "db-dsn"
	FlagDBDriver          string = "db-driver"
	FlagDBMaxOpenConns    string = "db-max-open-conns"
	FlagDBMaxIdleConns    string = "db-max-idle-conns"
	FlagDBMaxIdleLifetime string = "db-max-idle-lifetime"
	FlagDBConnTimeout     string = "db-conn-timeout"
)

// gRPC config flags
const (
	FlagGRPCAddress          string = "grpc-address"
	FlagGRPCMaxConnIdle      string = "grpc-max-conn-idle"
	FlagGRPCMaxConnAge       string = "grpc-max-conn-age"
	FlagGRPCMaxConnAgeGrace  string = "grpc-max-conn-age-grace"
	FlagGRPCTimeout          string = "grpc-timeout"
	FlagGRPCKeepAliveTime    string = "grpc-keep-alive-time"
	FlagGRPCKeepAliveTimeout string = "grpc-keep-alive-timeout"
	FlagGRPCShutdownTimeout  string = "grpc-shutdown-timeout"
)

// http config flags
const (
	FlagHTTPAddress            string = "http-address"
	FlagHTTPReadTimeout        string = "http-read-timeout"
	FlagHTTPWriteTimeout       string = "http-write-timeout"
	FlagHTTPIdleTimeout        string = "http-idle-timeout"
	FlagHTTPShutdownTimeout    string = "http-shutdown-timeout"
	FlagHTTPPrivateKeyPath     string = "http-private-key-path"
	FlagHTTPCertificatePath    string = "http-certificate-path"
	FlagHTTPSecure             string = "http-secure"
	FlagHTTPMaxRequestBodySize string = "http-max-request-body-size"
)

// log config flags
const (
	FlagLogLevel  string = "log-level"
	FlagLogFormat string = "log-format"
)

// telemetry
const (
	FlagTelemetryEnabled          string = "telemetry-enabled"
	FlagTelemetryServiceName      string = "telemetry-service-name"
	FlagTelemetryExporterEndpoint string = "telemetry-exporter-endpoint"
	FlagTelemetrySampleRate       string = "telemetry-sample-rate"
	FlagTelemetryTimeout          string = "telemetry-timeout"
)

// EnvConfig - файл конфигурации
const EnvConfig string = "CONFIG_PATH"

// amqp connector
const (
	FlagAMQPConnectorURL             string = "amqp-connector-url"
	FlagAMQPConnectorUsername        string = "amqp-connector-username"
	FlagAMQPConnectorPassword        string = "amqp-connector-password"
	FlagAMQPConnectorConnectTimeout  string = "amqp-connector-connect-timeout"
	FlagAMQPConnectorWriteTimeout    string = "amqp-connector-write-timeout"
	FlagAMQPConnectorIdleTimeout     string = "amqp-connector-idle-timeout"
	FlagAMQPConnectorShutdownTimeout string = "amqp-connector-shutdown-timeout"
)

const (
	FlagLoginAttemptsSenderKind          string = "login-attempts-sender-kind"
	FlagLoginAttemptsSenderNotifyTimeout string = "login-attempts-sender-notify-timeout"
)

// amqp login attempts sender
const (
	FlagLoginAttemptsSenderAMQPConfigTargetName            string = "login-attempts-sender-amqp-target-name"
	FlagLoginAttemptsSenderAMQPConfigConnectTimeout        string = "login-attempts-sender-amqp-connect-timeout"
	FlagLoginAttemptsSenderAMQPConfigShutdownTimeout       string = "login-attempts-sender-amqp-shutdown-timeout"
	FlagLoginAttemptsSenderAMQPConfigPublishMaxTryAttempts string = "login-attempts-sender-amqp-publish-max-try-attempts"
	FlagLoginAttemptsSenderAMQPConfigPublishBaseRetryDelay string = "login-attempts-sender-amqp-publish-base-retry-delay"
	FlagLoginAttemptsSenderAMQPConfigPublishMaxRetryDelay  string = "login-attempts-sender-amqp-publish-max-retry-delay"
)

// kafka login attempts sender
const (
	FlagLoginAttemptsSenderKafkaConfigBrokers               string = "login-attempts-sender-kafka-brokers"
	FlagLoginAttemptsSenderKafkaConfigTargetName            string = "login-attempts-sender-kafka-target-name"
	FlagLoginAttemptsSenderKafkaConfigConnectTimeout        string = "login-attempts-sender-kafka-connect-timeout"
	FlagLoginAttemptsSenderKafkaConfigShutdownTimeout       string = "login-attempts-sender-kafka-shutdown-timeout"
	FlagLoginAttemptsSenderKafkaConfigPublishMaxTryAttempts string = "login-attempts-sender-kafka-publish-max-try-attempts"
	FlagLoginAttemptsSenderKafkaConfigPublishMaxRetryDelay  string = "login-attempts-sender-kafka-publish-max-retry-delay"
	FlagLoginAttemptsSenderKafkaConfigPublishBaseRetryDelay string = "login-attempts-sender-kafka-publish-base-retry-delay"
	FlagLoginAttemptsSenderKafkaConfigUsername              string = "login-attempts-sender-kafka-username"
	FlagLoginAttemptsSenderKafkaConfigPassword              string = "login-attempts-sender-kafka-password"
	FlagLoginAttemptsSenderKafkaConfigBatchSize             string = "login-attempts-sender-kafka-batch-size"
	FlagLoginAttemptsSenderKafkaConfigBatchBytes            string = "login-attempts-sender-kafka-batch-bytes"
	FlagLoginAttemptsSenderKafkaConfigBatchTimeout          string = "login-attempts-sender-kafka-batch-timeout"
	FlagLoginAttemptsSenderKafkaConfigWriteTimeout          string = "login-attempts-sender-kafka-write-timeout"
	FlagLoginAttemptsSenderKafkaConfigRequiredAcks          string = "login-attempts-sender-kafka-required-acks"
)

// app
const (
	keyAppEnv                string = "app.env"
	keyAppNodeName           string = "app.node_name"
	keyAppMaxListLimit       string = "app.max_list_limit"
	keyAppTokenIssuer        string = "app.token_issuer"
	keyAppCipherKey          string = "app.cipher_key"
	keyAppDefShutdownTimeout string = "app.def_shutdown_timeout"
)

// service credentials
const (
	keySvcCredsUsername              string = "svc_creds.username"
	keySvcCredsPassword              string = "svc_creds.password"
	keySvcCredsScheduleInterval      string = "svc_creds.schedule_interval"
	keySvcCredsErrorScheduleInterval string = "svc_creds.error_schedule_interval"
)

// data audit client
const (
	keyDataAuditClientBaseURL            string = "data_audit_client.base_url"
	keyDataAuditClientTimeout            string = "data_audit_client.timeout"
	keyDataAuditClientWorkerCount        string = "data_audit_client.worker_count"
	keyDataAuditClientDataCapacity       string = "data_audit_client.data_capacity"
	keyDataAuditClientCompleteProcessing string = "data_audit_client.complete_processing"
	keyDataAuditClientShutdownTimeout    string = "data_audit_client.shutdown_timeout"
)

// amqp connector
const (
	keyAMQPConnectorURL             string = "amqp_connector.url"
	keyAMQPConnectorUsername        string = "amqp_connector.username"
	keyAMQPConnectorPassword        string = "amqp_connector.password"
	keyAMQPConnectorConnectTimeout  string = "amqp_connector.connect_timeout"
	keyAMQPConnectorWriteTimeout    string = "amqp_connector.write_timeout"
	keyAMQPConnectorIdleTimeout     string = "amqp_connector.idle_timeout"
	keyAMQPConnectorShutdownTimeout string = "amqp_connector.shutdown_timeout"
)

// login attempts sender
const (
	keyLoginAttemptsSenderKind          string = "login_attempts_sender.kind"
	keyLoginAttemptsSenderNotifyTimeout string = "login_attempts_sender.notify_timeout"
)

// amqp login attempts sender
const (
	keyLoginAttemptsSenderAMQPConfigTargetName            string = "login_attempts_sender.amqp_config.target_name"
	keyLoginAttemptsSenderAMQPConfigConnectTimeout        string = "login_attempts_sender.amqp_config.connect_timeout"
	keyLoginAttemptsSenderAMQPConfigShutdownTimeout       string = "login_attempts_sender.amqp_config.shutdown_timeout"
	keyLoginAttemptsSenderAMQPConfigPublishMaxTryAttempts string = "login_attempts_sender.amqp_config.publish_max_try_attempts"
	keyLoginAttemptsSenderAMQPConfigPublishBaseRetryDelay string = "login_attempts_sender.amqp_config.publish_base_retry_delay"
	keyLoginAttemptsSenderAMQPConfigPublishMaxRetryDelay  string = "login_attempts_sender.amqp_config.publish_max_retry_delay"
)

// kafka login attempts
const (
	keyLoginAttemptsSenderKafkaConfigBrokers               string = "login_attempts_sender.kafka_config.brokers"
	keyLoginAttemptsSenderKafkaConfigTargetName            string = "login_attempts_sender.kafka_config.target_name"
	keyLoginAttemptsSenderKafkaConfigConnectTimeout        string = "login_attempts_sender.kafka_config.connect_timeout"
	keyLoginAttemptsSenderKafkaConfigShutdownTimeout       string = "login_attempts_sender.kafka_config.shutdown_timeout"
	keyLoginAttemptsSenderKafkaConfigPublishMaxTryAttempts string = "login_attempts_sender.kafka_config.publish_max_try_attempts"
	keyLoginAttemptsSenderKafkaConfigPublishMaxRetryDelay  string = "login_attempts_sender.kafka_config.publish_max_retry_delay"
	keyLoginAttemptsSenderKafkaConfigPublishBaseRetryDelay string = "login_attempts_sender.kafka_config.publish_base_retry_delay"
	keyLoginAttemptsSenderKafkaConfigUsername              string = "login_attempts_sender.kafka_config.username"
	keyLoginAttemptsSenderKafkaConfigPassword              string = "login_attempts_sender.kafka_config.password"
	keyLoginAttemptsSenderKafkaConfigBatchSize             string = "login_attempts_sender.kafka_config.batch_size"
	keyLoginAttemptsSenderKafkaConfigBatchBytes            string = "login_attempts_sender.kafka_config.batch_bytes"
	keyLoginAttemptsSenderKafkaConfigBatchTimeout          string = "login_attempts_sender.kafka_config.batch_timeout"
	keyLoginAttemptsSenderKafkaConfigWriteTimeout          string = "login_attempts_sender.kafka_config.write_timeout"
	keyLoginAttemptsSenderKafkaConfigRequiredAcks          string = "login_attempts_sender.kafka_config.required_acks"
)
