package app

import (
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/db"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	libworker "github.com/ElfAstAhe/go-service-template/pkg/transport/worker"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/client/rest"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade"
	authaudit "github.com/ElfAstAhe/tiny-auth-service/internal/facade/audit"
	dataaudit "github.com/ElfAstAhe/tiny-auth-service/internal/repository/audit"
	"github.com/ElfAstAhe/tiny-auth-service/internal/repository/metrics"
	"github.com/ElfAstAhe/tiny-auth-service/internal/repository/postgres"
	"github.com/ElfAstAhe/tiny-auth-service/internal/repository/trace"
	"github.com/ElfAstAhe/tiny-auth-service/internal/transport/worker"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase/telemetry"
	pkgworker "github.com/ElfAstAhe/tiny-auth-service/pkg/transport/worker"
)

//goland:noinspection DuplicatedCode
func (app *App) initDependencies() error {
	var err error
	// transaction manager
	app.tm = db.NewTxManager(app.db)

	var (
		roleAdminRepo      domain.RoleAdminRepository
		userRolesAdminRepo domain.UserRolesAdminRepository
		userAdminRepo      domain.UserAdminRepository

		roleRepo      domain.RoleRepository
		userRolesRepo domain.UserRolesRepository
		userRepo      domain.UserRepository

		loginUC       usecase.LoginUseCase
		loginSimpleUC usecase.LoginSimpleUseCase

		registerUC       usecase.RegisterUseCase
		profileUC        usecase.ProfileUseCase
		changeKeysUC     usecase.ChangeKeysUseCase
		changePasswordUC usecase.ChangePasswordUseCase

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

		tokenRefresher *worker.TokenRefresher
	)
	// workers
	{
		tokenRefresher = worker.NewTokenRefresher(
			app.jwtHelper,
			nil, // <--- setup at the end
			app.config.Credentials,
			pkgworker.NewBaseTokenRefresherConfig(
				libworker.NewBaseSchedulerConfig(
					100*time.Millisecond,
					app.config.Credentials.ScheduleInterval,
				),
				app.config.Credentials.ErrorScheduleInterval,
			),
			app.logger,
		)
		app.tokenRefresher = tokenRefresher
	}
	// clients
	{
		// auth audit
		authAuditConf, err := rest.NewAuditClientConfig(
			app.config.AuthAuditClient.BaseURL,
			app.config.AuthAuditClient.Timeout,
			libworker.NewBasePoolConfig(
				app.config.AuthAuditClient.WorkerCount,
				app.config.AuthAuditClient.DataCapacity,
				app.config.AuthAuditClient.CompleteProcessing,
			),
		)
		if err != nil {
			return errs.NewCommonError("create auth audit config failed", err)
		}
		app.authAuditClient = rest.NewAuthAuditClient(authAuditConf, tokenRefresher, app.logger)

		// data audit
		dataAuditConf, err := rest.NewAuditClientConfig(
			app.config.DataAuditClient.BaseURL,
			app.config.DataAuditClient.Timeout,
			libworker.NewBasePoolConfig(
				app.config.DataAuditClient.WorkerCount,
				app.config.DataAuditClient.DataCapacity,
				app.config.DataAuditClient.CompleteProcessing,
			),
		)
		if err != nil {
			return errs.NewCommonError("create data audit config failed", err)
		}
		app.dataAuditClient = rest.NewDataAuditClient(dataAuditConf, tokenRefresher, app.logger)
	}
	// repositories
	{
		// role repo
		roleRepo, err = postgres.NewRolePgRepository(app.db, app.db)
		if err != nil {
			return err
		}
		roleRepo = dataaudit.NewRoleRepository(
			app.config.App.NodeName,
			trace.NewRoleTraceRepository(metrics.NewRoleMetricsRepository(roleRepo)),
			app.dataAuditClient,
			app.logger,
		)
		// user roles repo
		userRolesRepo, err = postgres.NewUserRolesPgRepository(app.db, app.db)
		if err != nil {
			return err
		}
		// user roles metrics repo
		userRolesRepo = trace.NewUserRolesTraceRepository(metrics.NewUserRolesMetricsRepository(userRolesRepo))
		// user repo
		userRepo, err = postgres.NewUserPgRepository(app.db, app.db, app.hashCipher, app.cipherHelper, userRolesRepo)
		if err != nil {
			return err
		}
		userRepo = dataaudit.NewUserRepository(
			app.config.App.NodeName,
			trace.NewUserTraceRepository(metrics.NewUserMetricsRepository(userRepo)),
			app.dataAuditClient,
			app.logger,
		)
		// role admin repo
		roleAdminRepo, err = postgres.NewRoleAdminPgRepository(app.db, app.db)
		if err != nil {
			return err
		}
		roleAdminRepo = dataaudit.NewRoleAdminRepository(
			app.config.App.NodeName,
			trace.NewRoleAdminTraceRepository(metrics.NewRoleAdminMetricsRepository(roleAdminRepo)),
			app.dataAuditClient,
			app.logger,
		)
		// user roles admin repo
		userRolesAdminRepo, err = postgres.NewUserRolesAdminPgRepository(app.db, app.db)
		if err != nil {
			return err
		}
		userRolesAdminRepo = trace.NewUserRolesAdminTraceRepository(metrics.NewUserRolesAdminMetricsRepository(userRolesAdminRepo))
		// user admin repo
		userAdminRepo, err = postgres.NewUserAdminPgRepository(app.db, app.db, app.cipherHelper, app.hashCipher, userRolesAdminRepo)
		if err != nil {
			return err
		}
		userAdminRepo = dataaudit.NewUserAdminRepository(
			app.config.App.NodeName,
			trace.NewUserAdminTraceRepository(metrics.NewUserAdminMetricsRepository(userAdminRepo)),
			app.dataAuditClient,
			app.logger,
		)
	}
	// use cases
	{
		// auth
		loginUC = telemetry.NewLoginTraceUseCase("LoginUseCase", usecase.NewLoginUseCase(app.hashCipher, app.keysHelper, app.authHelper, userRepo))
		loginSimpleUC = telemetry.NewLoginSimpleTraceUseCase("LoginSimpleUseCase", usecase.NewLoginSimpleUseCase(app.hashCipher, app.authHelper, userRepo))
		// users
		registerUC = telemetry.NewRegisterTraceUseCase("RegisterUseCase", usecase.NewRegisterUseCase(app.tm, app.hashCipher, app.keysHelper, userRepo))
		profileUC = telemetry.NewProfileTraceUseCase("ProfileUseCase", usecase.NewProfileUseCase(userRepo))
		changeKeysUC = telemetry.NewChangeKeysTraceUseCase("ChangeKeysUseCase", usecase.NewChangeKeysUseCase(app.keysHelper, app.tm, userRepo))
		changePasswordUC = telemetry.NewChangePasswordTraceUseCase("ChangePasswordUseCase", usecase.NewChangePasswordUseCase(app.hashCipher, app.tm, userRepo))
		// role admin
		roleAdminGetUC = telemetry.NewRoleAdminGetTraceUseCase("RoleAdminGetUseCase", usecase.NewRoleAdminGetUseCase(roleAdminRepo))
		roleAdminGetByNameUC = telemetry.NewROleAdminGetNameTraceUseCase("RoleAdminGetNameUseCase", usecase.NewRoleAdminGetNameUseCase(roleAdminRepo))
		roleAdminListUC = telemetry.NewRoleAdminListTraceUseCase("RoleAdminListUseCase", usecase.NewRoleAdminListUseCase(roleAdminRepo, app.config.App.MaxListLimit))
		roleAdminSaveUC = telemetry.NewRoleAdminSaveTraceUseCase("RoleAdminSaveUseCase", usecase.NewRoleAdminSaveUseCase(app.tm, roleAdminRepo))
		roleAdminDeleteUC = telemetry.NewRoleAdminDeleteTraceUseCase("RoleAdminDeleteUseCase", usecase.NewRoleAdminDeleteUseCase(app.tm, roleAdminRepo))
		// user admin
		userAdminGetUC = telemetry.NewUserAdminGetTraceUseCase("UserAdminGetUseCase", usecase.NewUserAdminGetUseCase(userAdminRepo))
		userAdminGetByNameUC = telemetry.NewUserAdminGetNameTraceUseCase("UserAdminGetNameUseCase", usecase.NewUserAdminGetNameUseCase(userAdminRepo))
		userAdminListUC = telemetry.NewUserAdminListTraceUseCase("UserAdminListUseCase", usecase.NewUserAdminListUseCase(userAdminRepo, app.config.App.MaxListLimit))
		userAdminSaveUC = telemetry.NewUserAdminSaveTraceUseCase("UserAdminSaveUseCase", usecase.NewUserAdminSaveUseCase(app.tm, app.hashCipher, app.keysHelper, userAdminRepo))
		userAdminDeleteUC = telemetry.NewUserAdminDeleteTraceUseCase("UserAdminDeleteUseCase", usecase.NewUserAdminDeleteUseCase(app.tm, userAdminRepo))
	}
	// workers post setup
	{
		tokenRefresher.SetSimpleLoginUC(loginSimpleUC)
	}
	// facades
	{
		// auth
		app.authFacade = authaudit.NewAuthFacade(
			app.authAuditClient,
			app.config.App.NodeName,
			facade.NewAuthFacade(
				app.jwtHelper,
				loginUC,
				loginSimpleUC,
			),
		)
		// user
		app.userFacade = facade.NewUserFacade(
			app.authHelper,
			registerUC,
			profileUC,
			changePasswordUC,
			changeKeysUC,
		)
		// role admin
		app.roleAdminFacade = facade.NewRoleAdminFacade(
			app.authHelper,
			roleAdminGetUC,
			roleAdminGetByNameUC,
			roleAdminListUC,
			roleAdminSaveUC,
			roleAdminDeleteUC,
			app.config.App.MaxListLimit,
		)
		// user admin
		app.userAdminFacade = facade.NewUserAdminFacade(
			app.authHelper,
			userAdminGetUC,
			userAdminGetByNameUC,
			userAdminListUC,
			userAdminSaveUC,
			userAdminDeleteUC,
			app.config.App.MaxListLimit,
		)
	}

	return nil
}
