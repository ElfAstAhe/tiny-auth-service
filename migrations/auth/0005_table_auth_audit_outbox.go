package auth

import (
	"context"
	"database/sql"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/pressly/goose/v3"
)

func up0005(ctx context.Context, db *sql.DB) error {
	if err := createTableAuthAuditOutbox(ctx, db); err != nil {
		return err
	}
	if err := createIndexAuthAuditOutboxProcessed(ctx, db); err != nil {
		return err
	}

	return createIndexAuthAuditOutboxPending(ctx, db)
}

func createTableAuthAuditOutbox(ctx context.Context, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, sqlCreateTableAuthAuditOutbox); err != nil {
		return errs.NewDBMigrationError("failed to create table auth_audit_outbox", err)
	}

	return nil
}

func createIndexAuthAuditOutboxPending(ctx context.Context, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, sqlCreateIndexAuthAuditOutboxPending); err != nil {
		return errs.NewDBMigrationError("failed to create index idx_auth_audit_outbox_pending", err)
	}

	return nil
}

func createIndexAuthAuditOutboxProcessed(ctx context.Context, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, sqlCreateIndexAuthAuditOutboxProcessed); err != nil {
		return errs.NewDBMigrationError("failed to create index idx_auth_audit_outbox_processed", err)
	}

	return nil
}

func down0005(ctx context.Context, db *sql.DB) error {
	if err := dropIndexAuthAuditOutboxPending(ctx, db); err != nil {
		return err
	}
	if err := dropIndexAuthAuditOutboxProcessed(ctx, db); err != nil {
		return err
	}

	return dropTableAuthAuditOutbox(ctx, db)
}

func dropTableAuthAuditOutbox(ctx context.Context, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, sqlDropTableAuthAuditOutbox); err != nil {
		return errs.NewDBMigrationError("failed to drop table auth_audit_outbox", err)
	}

	return nil
}

func dropIndexAuthAuditOutboxPending(ctx context.Context, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, sqlDropIndexAuthAuditOutboxPending); err != nil {
		return errs.NewDBMigrationError("failed to drop index idx_auth_audit_outbox_pending", err)
	}

	return nil
}

func dropIndexAuthAuditOutboxProcessed(ctx context.Context, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, sqlDropIndexAuthAuditOutboxProcessed); err != nil {
		return errs.NewDBMigrationError("failed to create index idx_auth_audit_outbox_processed", err)
	}

	return nil
}

func init() {
	goose.AddMigrationNoTxContext(up0005, down0005)
}
