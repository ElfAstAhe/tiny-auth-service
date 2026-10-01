package config

import (
	"strings"
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
)

// ServiceCredentialsConfig encapsulates configuration criteria parameters payload
// managing automated service-to-machine background login operations and scheduler intervals.
type ServiceCredentialsConfig struct {
	Username              string        `mapstructure:"username" json:"username,omitempty" yaml:"username,omitempty"`
	Password              string        `mapstructure:"password" json:"password,omitempty" yaml:"password,omitempty"`
	ScheduleInterval      time.Duration `mapstructure:"schedule_interval" json:"schedule_interval,omitempty" yaml:"schedule_interval,omitempty"`
	ErrorScheduleInterval time.Duration `mapstructure:"error_schedule_interval" json:"error_schedule_interval,omitempty" yaml:"error_schedule_interval,omitempty"`
}

// NewServiceCredentialsConfig acts as a factory constructor allocating credentials records structures.
func NewServiceCredentialsConfig(username, password string) *ServiceCredentialsConfig {
	return &ServiceCredentialsConfig{
		Username: username,
		Password: password,
	}
}

// NewDefaultServiceCredentialsConfig returns a zero-allocated configuration schema blueprint stub.
func NewDefaultServiceCredentialsConfig() *ServiceCredentialsConfig {
	return NewServiceCredentialsConfig("", "")
}

// Validate executes strict fail-fast validation logic across structural setup parameters prior to application startup bounds.
func (scc *ServiceCredentialsConfig) Validate() error {
	if strings.TrimSpace(scc.Username) == "" {
		return errs.NewConfigValidateError("svc_creds", "username", "must not be empty", nil)
	}
	if strings.TrimSpace(scc.Password) == "" {
		return errs.NewConfigValidateError("svc_creds", "password", "must not be empty", nil)
	}

	return nil
}
