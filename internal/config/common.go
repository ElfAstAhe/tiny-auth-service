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

// Auth-audit-client
const (
	FlagAuthAuditClientBaseURL            string = "auth-audit-client-base-url"
	FlagAuthAuditClientTimeout            string = "auth-audit-client-timeout"
	FlagAuthAuditClientWorkerCount        string = "auth-audit-client-worker-count"
	FlagAuthAuditClientDataCapacity       string = "auth-audit-client-data-capacity"
	FlagAuthAuditClientCompleteProcessing string = "auth-audit-client-complete-processing"
	FlagAuthAuditClientShutdownTimeout    string = "auth-audit-client-shutdown-timeout"
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

// amqp sender
const (
	FlagLoginAttemptsSenderURL                string = "login-attempts-sender-url"
	FlagLoginAttemptsSenderTargetName         string = "login-attempts-sender-target-name"
	FlagLoginAttemptsSenderUsername           string = "login-attempts-sender-username"
	FlagLoginAttemptsSenderPassword           string = "login-attempts-sender-password"
	FlagLoginAttemptsSenderInsecureSkipVerify string = "login-attempts-sender-insecure-skip-verify"
	FlagLoginAttemptsSenderConnectTimeout     string = "login-attempts-sender-connect-timeout"
	FlagLoginAttemptsSenderWriteTimeout       string = "login-attempts-sender-write-timeout"
	FlagLoginAttemptsSenderNotifyTimeout      string = "login-attempts-sender-notify-timeout"
	FlagLoginAttemptsSenderShutdownTimeout    string = "login-attempts-sender-shutdown-timeout"
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

// auth audit client
const (
	keyAuthAuditClientBaseURL            string = "auth_audit_client.base_url"
	keyAuthAuditClientTimeout            string = "auth_audit_client.timeout"
	keyAuthAuditClientWorkerCount        string = "auth_audit_client.worker_count"
	keyAuthAuditClientDataCapacity       string = "auth_audit_client.data_capacity"
	keyAuthAuditClientCompleteProcessing string = "auth_audit_client.complete_processing"
	keyAuthAuditClientShutdownTimeout    string = "auth_audit_client.shutdown_timeout"
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

// amqp sender
const (
	keyLoginAttemptsSenderURL                string = "login_attempts_sender.url"
	keyLoginAttemptsSenderTargetName         string = "login_attempts_sender.target_name"
	keyLoginAttemptsSenderUsername           string = "login_attempts_sender.username"
	keyLoginAttemptsSenderPassword           string = "login_attempts_sender.password"
	keyLoginAttemptsSenderInsecureSkipVerify string = "login_attempts_sender.insecure_skip_verify"
	keyLoginAttemptsSenderConnectTimeout     string = "login_attempts_sender.connect_timeout"
	keyLoginAttemptsSenderWriteTimeout       string = "login_attempts_sender.write_timeout"
	keyLoginAttemptsSenderNotifyTimeout      string = "login_attempts_sender.notify_timeout"
	keyLoginAttemptsSenderShutdownTimeout    string = "login_attempts_sender.shutdown_timeout"
)
