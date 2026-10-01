package container

import (
	"context"
	"errors"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/db"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/go-service-template/pkg/migration"
)

const (
	// InstanceDB defines the global lookup token key targeted for primary SQL СУБД client interface.
	InstanceDB string = "DB"
	// InstanceDBMigrator maps the configuration trajectory identifier for relational schema migrations utilities.
	InstanceDBMigrator string = "DBMigrator"
	// InstanceTM structures the lookup token mapping the ACID atomic transaction manager boundary interface.
	InstanceTM string = "transaction-manager"
)

// PgContainer database connection and data migrations
type PgContainer struct {
	*container.BaseLazyContainer // Generic framework-level baseline container orchestration handle
}

// Compile-time interface compliance verifications
var _ container.Container = (*PgContainer)(nil)
var _ container.LazyContainer = (*PgContainer)(nil)

// NewPgContainer acts as a factory constructor deploying structural configuration, logger, and orchestrator boundaries.
func NewPgContainer(
	orchestrator container.Orchestrator,
	log logger.Logger,
) *PgContainer {
	return &PgContainer{
		BaseLazyContainer: container.NewBaseLazyContainer(
			container.WithLazyName(DBContainerName),
			container.WithLazyOrchestrator(orchestrator),
			container.WithLazyLogger(log),
		),
	}
}

// Init triggers synchronous registration sequences linking database engines, performs connections ping checks, and applies schema migrations.
func (pc *PgContainer) Init(initCtx context.Context) error {
	// add providers
	err := errors.Join(
		pc.RegisterProvider(InstanceDB, pc.providerDB),
		pc.RegisterProvider(InstanceDBMigrator, pc.providerDBMigrator),
		pc.RegisterProvider(InstanceTM, pc.providerTM),
	)
	if err != nil {
		return errs.NewContainerError(pc.GetName(), "container init: register providers failed", err)
	}
	// init db instance
	dbInst, err := container.GetInstance[db.DB](InstanceDB)
	if err != nil {
		return errs.NewContainerError(pc.GetName(), "container init: init db failed", err)
	}
	// check db connection
	err = dbInst.Ping(initCtx)
	if err != nil {
		return errs.NewContainerError(pc.GetName(), "container init: check db failed", err)
	}
	// data migration
	migrator, err := container.GetInstance[migration.Migrator](InstanceDBMigrator)
	if err != nil {
		return errs.NewContainerError(pc.GetName(), "container init: init migrator failed", err)
	}
	// migrate up
	// BOOTSTRAP NOTICE: Synchronous migration execution blocks the container initialization path.
	// Long-running DDL modifications under heavy tables may exhaust context budget limits before network ports bind.
	err = migrator.Up(initCtx)
	if err != nil {
		return errs.NewContainerError(pc.GetName(), "container init: up migrator failed", err)
	}

	return nil
}
