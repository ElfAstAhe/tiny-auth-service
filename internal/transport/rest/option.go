package rest

import (
	"github.com/ElfAstAhe/go-service-template/pkg/auth"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	libhttp "github.com/ElfAstAhe/go-service-template/pkg/transport/http"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/config"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade"
	"github.com/hellofresh/health-go/v5"
)

// AppRouterOptions encapsulates configuration credentials, facades, core helpers,
// and diagnostics tooling required to spin up the HTTP transport delivery router.
type AppRouterOptions struct {
	Conf            *config.Config
	Logger          logger.Logger
	AuthHelper      auth.Helper
	Health          *health.Health
	Healthz         libhttp.HealthzFunc
	Readyz          libhttp.ReadyzFunc
	AuthFacade      facade.AuthFacade
	UserFacade      facade.UserFacade
	UserAdminFacade facade.UserAdminFacade
	RoleAdminFacade facade.RoleAdminFacade
}

// Validate executes strict fail-fast validation logic across all mandatory router dependency bounds.
func (aro *AppRouterOptions) Validate() error {
	if utils.IsNil(aro.Conf) {
		return errs.NewTlCommonError("validate", "conf not applied", nil)
	}
	if utils.IsNil(aro.Logger) {
		return errs.NewTlCommonError("validate", "logger not applied", nil)
	}
	if utils.IsNil(aro.AuthHelper) {
		return errs.NewTlCommonError("validate", "auth helper not applied", nil)
	}
	if utils.IsNil(aro.Health) {
		return errs.NewTlCommonError("validate", "health not applied", nil)
	}
	//if utils.IsNil(aro.Healthz) {
	//    return errs.NewTlCommonError("validate", "healthz not applied", nil)
	//}
	//if utils.IsNil(aro.Readyz) {
	//    return errs.NewTlCommonError("validate", "readyz not applied", nil)
	//}
	if utils.IsNil(aro.AuthFacade) {
		return errs.NewTlCommonError("validate", "auth facade not applied", nil)
	}
	if utils.IsNil(aro.UserFacade) {
		return errs.NewTlCommonError("validate", "user facade not applied", nil)
	}
	if utils.IsNil(aro.UserAdminFacade) {
		return errs.NewTlCommonError("validate", "user admin facade not applied", nil)
	}
	if utils.IsNil(aro.RoleAdminFacade) {
		return errs.NewTlCommonError("validate", "role admin facade not applied", nil)
	}

	return nil
}

// Option defines a functional configuration closure pattern designed to lazily populate AppRouterOptions credentials.
type Option func(*AppRouterOptions)

// WithConfig returns an Option configuring the application layout Config credentials block.
func WithConfig(conf *config.Config) Option {
	return func(aro *AppRouterOptions) {
		aro.Conf = conf
	}
}

// WithLogger returns an Option mapping the structured logging subsystem handle into the router dependencies graph.
func WithLogger(logger logger.Logger) Option {
	return func(aro *AppRouterOptions) {
		aro.Logger = logger
	}
}

// WithAuthHelper returns an Option mounting the core cryptographic security and payload validation utility engine.
func WithAuthHelper(helper auth.Helper) Option {
	return func(aro *AppRouterOptions) {
		aro.AuthHelper = helper
	}
}

// WithHealth returns an Option specifying the structural health-go diagnostics provider monitor.
func WithHealth(health *health.Health) Option {
	return func(aro *AppRouterOptions) {
		aro.Health = health
	}
}

// WithHealthz returns an Option assigning the generic liveness probe verification callback.
func WithHealthz(healthz libhttp.HealthzFunc) Option {
	return func(aro *AppRouterOptions) {
		aro.Healthz = healthz
	}
}

// WithReadyz returns an Option assigning the abstract network readiness check criteria callback.
func WithReadyz(readyz libhttp.ReadyzFunc) Option {
	return func(aro *AppRouterOptions) {
		aro.Readyz = readyz
	}
}

// WithAuthFacade returns an Option mapping structural security session orchestrator boundaries.
func WithAuthFacade(facade facade.AuthFacade) Option {
	return func(aro *AppRouterOptions) {
		aro.AuthFacade = facade
	}
}

// WithUserFacade returns an Option embedding core user profile business scenarios controllers.
func WithUserFacade(userFacade facade.UserFacade) Option {
	return func(aro *AppRouterOptions) {
		aro.UserFacade = userFacade
	}
}

// WithUserAdminFacade returns an Option setting up identity moderation facade boundaries.
func WithUserAdminFacade(userAdminFacade facade.UserAdminFacade) Option {
	return func(aro *AppRouterOptions) {
		aro.UserAdminFacade = userAdminFacade
	}
}

// WithRoleAdminFacade returns an Option structuring administrative RBAC validation facade components.
func WithRoleAdminFacade(roleAdminFacade facade.RoleAdminFacade) Option {
	return func(aro *AppRouterOptions) {
		aro.RoleAdminFacade = roleAdminFacade
	}
}
