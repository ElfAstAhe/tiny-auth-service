package container

import (
	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/ext/transport/worker/tailcutter"
	"github.com/ElfAstAhe/go-service-template/pkg/ext/usecase"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/tiny-auth-service/internal/config"
)

func (wc *WorkerContainer) providerDataAuditOutboxTailCutter() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(wc.GetName(), "provider: retrieve instance failed", err)
	}
	logInst, err := container.GetInstance[logger.Logger](InstanceLogger)
	if err != nil {
		return nil, errs.NewContainerError(wc.GetName(), "provider: retrieve instance failed", err)
	}
	tailGetUCInst, err := container.GetInstance[usecase.TailGetUseCase[string]](InstanceAuthAuditOutboxGetTailUC)
	if err != nil {
		return nil, errs.NewContainerError(wc.GetName(), "provider: retrieve instance failed", err)
	}
	tailCutUCInst, err := container.GetInstance[usecase.TailCutUseCase[string]](InstanceAuthAuditOutboxCutTailUC)
	if err != nil {
		return nil, errs.NewContainerError(wc.GetName(), "provider: retrieve instance failed", err)
	}

	res, err := tailcutter.New(
		tailcutter.WithName("auth-audit-outbox"),
		tailcutter.WithWorkerCount(confInst.AuthAuditOutboxTailCutter.Dispatcher.WorkerCount),
		tailcutter.WithDataCapacity(confInst.AuthAuditOutboxTailCutter.Dispatcher.DataCapacity),
		tailcutter.WithCompleteProcess(confInst.AuthAuditOutboxTailCutter.Dispatcher.CompleteProcess),
		tailcutter.WithStartInterval(confInst.AuthAuditOutboxTailCutter.Dispatcher.StartInterval),
		tailcutter.WithScheduleInterval(confInst.AuthAuditOutboxTailCutter.Dispatcher.ScheduleInterval),
		tailcutter.WithStopTimeout(confInst.AuthAuditOutboxTailCutter.Dispatcher.StopTimeout),
		tailcutter.WithDataInterval(confInst.AuthAuditOutboxTailCutter.TailInterval),
		tailcutter.WithCutEnabled(confInst.AuthAuditOutboxTailCutter.TailCut),
		tailcutter.WithLogger(logInst),
		tailcutter.WithTailGetUseCase(tailGetUCInst),
		tailcutter.WithTailCutUseCase(tailCutUCInst),
	)
	if err != nil {

	}

	return res, nil
}

func (wc *WorkerContainer) providerAuthAuditOutboxTailCutter() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(wc.GetName(), "provider: retrieve instance failed", err)
	}
	logInst, err := container.GetInstance[logger.Logger](InstanceLogger)
	if err != nil {
		return nil, errs.NewContainerError(wc.GetName(), "provider: retrieve instance failed", err)
	}
	tailGetUCInst, err := container.GetInstance[usecase.TailGetUseCase[string]](InstanceDataAuditOutboxGetTailUC)
	if err != nil {
		return nil, errs.NewContainerError(wc.GetName(), "provider: retrieve instance failed", err)
	}
	tailCutUCInst, err := container.GetInstance[usecase.TailCutUseCase[string]](InstanceDataAuditOutboxCutTailUC)
	if err != nil {
		return nil, errs.NewContainerError(wc.GetName(), "provider: retrieve instance failed", err)
	}

	res, err := tailcutter.New(
		tailcutter.WithName("data-audit-outbox"),
		tailcutter.WithWorkerCount(confInst.DataAuditOutboxTailCutter.Dispatcher.WorkerCount),
		tailcutter.WithDataCapacity(confInst.DataAuditOutboxTailCutter.Dispatcher.DataCapacity),
		tailcutter.WithCompleteProcess(confInst.DataAuditOutboxTailCutter.Dispatcher.CompleteProcess),
		tailcutter.WithStartInterval(confInst.DataAuditOutboxTailCutter.Dispatcher.StartInterval),
		tailcutter.WithScheduleInterval(confInst.DataAuditOutboxTailCutter.Dispatcher.ScheduleInterval),
		tailcutter.WithStopTimeout(confInst.DataAuditOutboxTailCutter.Dispatcher.StopTimeout),
		tailcutter.WithDataInterval(confInst.DataAuditOutboxTailCutter.TailInterval),
		tailcutter.WithCutEnabled(confInst.DataAuditOutboxTailCutter.TailCut),
		tailcutter.WithLogger(logInst),
		tailcutter.WithTailGetUseCase(tailGetUCInst),
		tailcutter.WithTailCutUseCase(tailCutUCInst),
	)
	if err != nil {

	}

	return res, nil
}
