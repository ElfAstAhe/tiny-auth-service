package container

import (
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/auth"
	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/go-service-template/pkg/transport/http"
	"github.com/ElfAstAhe/tiny-auth-service/internal/config"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade"
	"github.com/ElfAstAhe/tiny-auth-service/internal/transport/rest"
	"github.com/hellofresh/health-go/v5"
)

//goland:noinspection DuplicatedCode
func (hc *HTTPContainer) providerChiRouter() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(hc.GetName(), "provider: retrieve instance failed", err)
	}
	logInst, err := container.GetInstance[logger.Logger](InstanceLogger)
	if err != nil {
		return nil, errs.NewContainerError(hc.GetName(), "provider: retrieve instance failed", err)
	}
	readyzInst, err := container.GetInstance[func() bool](InstanceApplicationReady)
	if err != nil {
		return nil, errs.NewContainerError(hc.GetName(), "provider: retrieve instance failed", err)
	}
	authHelperInst, err := container.GetInstance[auth.Helper](InstanceAuthHelper)
	if err != nil {
		return nil, errs.NewContainerError(hc.GetName(), "provider: retrieve instance failed", err)
	}
	authFacadeInst, err := container.GetInstance[facade.AuthFacade](InstanceAuthFacade)
	if err != nil {
		return nil, errs.NewContainerError(hc.GetName(), "provider: retrieve instance failed", err)
	}
	userFacadeInst, err := container.GetInstance[facade.UserFacade](InstanceUserFacade)
	if err != nil {
		return nil, errs.NewContainerError(hc.GetName(), "provider: retrieve instance failed", err)
	}
	userAdminFacadeInst, err := container.GetInstance[facade.UserAdminFacade](InstanceUserAdminFacade)
	if err != nil {
		return nil, errs.NewContainerError(hc.GetName(), "provider: retrieve instance failed", err)
	}
	roleAdminFacadeInst, err := container.GetInstance[facade.RoleAdminFacade](InstanceRoleAdminFacade)
	if err != nil {
		return nil, errs.NewContainerError(hc.GetName(), "provider: retrieve instance failed", err)
	}
	healthInst, err := container.GetInstance[*health.Health](InstanceHealthStatus)
	if err != nil {
		return nil, errs.NewContainerError(hc.GetName(), "provider: retrieve instance failed", err)
	}

	return rest.NewAppRouter(
		rest.WithConfig(confInst),
		rest.WithLogger(logInst),
		rest.WithAuthHelper(authHelperInst),
		rest.WithHealth(healthInst),
		rest.WithReadyz(readyzInst),
		rest.WithAuthFacade(authFacadeInst),
		rest.WithUserFacade(userFacadeInst),
		rest.WithUserAdminFacade(userAdminFacadeInst),
		rest.WithRoleAdminFacade(roleAdminFacadeInst),
	)
}

//goland:noinspection DuplicatedCode
func (hc *HTTPContainer) providerHTTPRunner() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(hc.GetName(), "provider: retrieve instance failed", err)
	}
	logInst, err := container.GetInstance[logger.Logger](InstanceLogger)
	if err != nil {
		return nil, errs.NewContainerError(hc.GetName(), "provider: retrieve instance failed", err)
	}
	routerInst, err := container.GetInstance[http.Router](InstanceHTTPRouter)
	if err != nil {
		return nil, errs.NewContainerError(hc.GetName(), "provider: retrieve instance failed", err)
	}

	runner, err := http.NewRunner(
		http.WithName("main-http-server"),
		http.WithConfig(confInst.HTTP),
		http.WithLogger("http_server", logInst),
		http.WithRouter(routerInst),
	)
	if err != nil {
		return nil, errs.NewContainerError(hc.GetName(), fmt.Sprintf("provider: create %s failed", InstanceHTTPRunner), err)
	}

	return runner, nil
}
