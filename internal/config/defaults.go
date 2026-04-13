package config

import (
	"time"
)

// app
const (
	defaultAppEnv             AppEnv        = AppEnvDevelopment
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

// audit client
const (
	defaultAuditClientTimeout            time.Duration = 5 * time.Second
	defaultAuditClientWorkerCount        int           = 4
	defaultAuditClientDataCapacity       int           = 10000
	defaultAuditClientCompleteProcessing bool          = true
	defaultAuditClientShutdownTimeout    time.Duration = 15 * time.Second
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
