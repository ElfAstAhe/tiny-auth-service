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

type Option func(*AppRouterOptions)

func WithConfig(conf *config.Config) Option {
	return func(aro *AppRouterOptions) {
		aro.Conf = conf
	}
}

func WithLogger(logger logger.Logger) Option {
	return func(aro *AppRouterOptions) {
		aro.Logger = logger
	}
}

func WithAuthHelper(helper auth.Helper) Option {
	return func(aro *AppRouterOptions) {
		aro.AuthHelper = helper
	}
}

func WithHealth(health *health.Health) Option {
	return func(aro *AppRouterOptions) {
		aro.Health = health
	}
}

func WithHealthz(healthz libhttp.HealthzFunc) Option {
	return func(aro *AppRouterOptions) {
		aro.Healthz = healthz
	}
}

func WithReadyz(readyz libhttp.ReadyzFunc) Option {
	return func(aro *AppRouterOptions) {
		aro.Readyz = readyz
	}
}

func WithAuthFacade(facade facade.AuthFacade) Option {
	return func(aro *AppRouterOptions) {
		aro.AuthFacade = facade
	}
}

func WithUserFacade(userFacade facade.UserFacade) Option {
	return func(aro *AppRouterOptions) {
		aro.UserFacade = userFacade
	}
}

func WithUserAdminFacade(userAdminFacade facade.UserAdminFacade) Option {
	return func(aro *AppRouterOptions) {
		aro.UserAdminFacade = userAdminFacade
	}
}

func WithRoleAdminFacade(roleAdminFacade facade.RoleAdminFacade) Option {
	return func(aro *AppRouterOptions) {
		aro.RoleAdminFacade = roleAdminFacade
	}
}
