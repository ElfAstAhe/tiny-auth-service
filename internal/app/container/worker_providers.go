package container

import (
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/helper"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	libworker "github.com/ElfAstAhe/go-service-template/pkg/transport/worker"
	"github.com/ElfAstAhe/tiny-auth-service/internal/config"
	"github.com/ElfAstAhe/tiny-auth-service/internal/transport/worker"
	pkgworker "github.com/ElfAstAhe/tiny-auth-service/pkg/transport/worker"
)

func (wc *WorkerContainer) providerTokenRefresher() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(wc.GetName(), "provider: retrieve instance failed", err)
	}
	logInst, err := container.GetInstance[logger.Logger](InstanceLogger)
	if err != nil {
		return nil, errs.NewContainerError(wc.GetName(), "provider: retrieve instance failed", err)
	}
	jwtHelperInst, err := container.GetInstance[*helper.JWTHelper](InstanceJWTHelper)
	if err != nil {
		return nil, errs.NewContainerError(wc.GetName(), "provider: retrieve instance failed", err)
	}

	return worker.NewTokenRefresher(
		jwtHelperInst,
		nil, // <--- setup at the end
		confInst.Credentials,
		pkgworker.NewBaseTokenRefresherConfig(
			libworker.NewBaseSchedulerConfig(
				100*time.Millisecond,
				confInst.Credentials.ScheduleInterval,
				confInst.App.DefShutdownTimeout,
			),
			confInst.Credentials.ErrorScheduleInterval,
		),
		logInst,
	), nil
}
