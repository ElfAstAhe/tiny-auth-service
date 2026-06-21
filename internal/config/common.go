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
	FlagAMQPSenderURL                string = "amqp-sender-url"
	FlagAMQPSenderAddress            string = "amqp-sender-address"
	FlagAMQPSenderTargetName         string = "amqp-sender-target-name"
	FlagAMQPSenderUsername           string = "amqp-sender-username"
	FlagAMQPSenderPassword           string = "amqp-sender-password"
	FlagAMQPSenderInsecureSkipVerify string = "amqp-sender-insecure-skip-verify"
	FlagAMQPSenderConnectTimeout     string = "amqp-sender-connect-timeout"
	FlagAMQPSenderWriteTimeout       string = "amqp-sender-write-timeout"
	FlagAMQPSenderNotifyTimeout      string = "amqp-sender-notify-timeout"
	FlagAMQPSenderShutdownTimeout    string = "amqp-sender-shutdown-timeout"
)
