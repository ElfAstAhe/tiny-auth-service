package app

import (
	"context"
	"net/http"
	"sync"

	"github.com/ElfAstAhe/go-service-template/pkg/db"
	"github.com/ElfAstAhe/go-service-template/pkg/helper"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/go-service-template/pkg/transport"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/config"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade"
	grpcsvc "github.com/ElfAstAhe/tiny-auth-service/internal/transport/grpc"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
	"github.com/hellofresh/health-go/v5"
	"google.golang.org/grpc"
)

type App struct {
	// app context
	ctx    context.Context
	cancel context.CancelFunc
	// app config
	config *config.Config
	// logging
	logger logger.Logger
	// telemetry
	telemetryShutdown func(ctx context.Context) error

	// helpers
	hashCipher   utils.Cipher
	dataCipher   utils.Cipher
	cipherHelper helper.Cipher
	keysHelper   helper.RSAKeys
	jwtHelper    *helper.JWTHelper

	// DB
	db db.DB

	// infra
	wg sync.WaitGroup

	// checkers
	health *health.Health

	// tx
	tm db.TransactionManager

	// http
	httpRouter transport.HTTPRouter
	httpServer *http.Server

	// gRPC
	grpcAuthService      *grpcsvc.AuthGRPCService
	grpcRoleAdminService *grpcsvc.RoleAdminGRPCService
	grpcUserAdminService *grpcsvc.UserAdminGRPCService
	grpcServer           *grpc.Server

	// use cases
	userAdminGetUC       usecase.UserAdminGetUseCase
	userAdminGetByNameUC usecase.UserAdminGetNameUseCase
	userAdminListUC      usecase.UserAdminListUseCase
	userAdminSaveUC      usecase.UserAdminSaveUseCase
	userAdminDeleteUC    usecase.UserAdminDeleteUseCase

	roleAdminGetUC       usecase.RoleAdminGetUseCase
	roleAdminGetByNameUC usecase.RoleAdminGetNameUseCase
	roleAdminListUC      usecase.RoleAdminListUseCase
	roleAdminSaveUC      usecase.RoleAdminSaveUseCase
	roleAdminDeleteUC    usecase.RoleAdminDeleteUseCase

	// facade
	roleAdminFacade facade.RoleAdminFacade
	userAdminFacade facade.UserAdminFacade
}

func NewApp(config *config.Config, logger logger.Logger) *App {
	appCtx, appCancel := context.WithCancel(context.Background())

	return &App{
		ctx:    appCtx,
		cancel: appCancel,
		config: config,
		logger: logger,
	}
}

// Stop - метод остановки приложения
func (app *App) Stop() {
	app.cancel()
}

func (app *App) WaitForStop() {
	app.wg.Wait()
}
