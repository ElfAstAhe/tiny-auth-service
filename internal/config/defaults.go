package config

import (
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/config"
)

// app
const (
	defaultAppEnv             config.AppEnv = config.AppEnvDevelopment
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

// amqp sender (FQQN artemis style)
const (
	defaultLoginAttemptsSenderTargetName string = "tiny.auth::login.attempts"
	defaultLoginAttemptsSenderUsername   string = "svc-auth"
	defaultLoginAttemptsSenderPassword   string = "test"
)
