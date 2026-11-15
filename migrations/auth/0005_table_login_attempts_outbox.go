package auth

import (
	"context"
	"database/sql"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/pressly/goose/v3"
)

func up0005(ctx context.Context, db *sql.DB) error {
	if err := createTableLoginAttemptsOutbox(ctx, db); err != nil {
		return err
	}

	return createIndexLoginAttemptsOutboxPending(ctx, db)
}

func createTableLoginAttemptsOutbox(ctx context.Context, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, sqlCreateTableLoginAttemptsOutbox); err != nil {
		return errs.NewDBMigrationError("failed to create table login_attempts_outbox", err)
	}

	return nil
}

func createIndexLoginAttemptsOutboxPending(ctx context.Context, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, sqlCreateIndexLoginAttemptsOutboxPending); err != nil {
		return errs.NewDBMigrationError("failed to create index idx_login_attempts_outbox_pending", err)
	}

	return nil
}

func down0005(ctx context.Context, db *sql.DB) error {
	if err := dropIndexLoginAttemptsOutboxPending(ctx, db); err != nil {
		return err
	}

	return dropTableLoginAttemptsOutbox(ctx, db)
}

func dropTableLoginAttemptsOutbox(ctx context.Context, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, sqlDropTableLoginAttemptsOutbox); err != nil {
		return errs.NewDBMigrationError("failed to drop table login_attempts_outbox", err)
	}

	return nil
}

func dropIndexLoginAttemptsOutboxPending(ctx context.Context, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, sqlDropIndexLoginAttemptsOutboxPending); err != nil {
		return errs.NewDBMigrationError("failed to drop index idx_login_attempts_outbox_pending", err)
	}

	return nil
}

func init() {
	goose.AddMigrationNoTxContext(up0005, down0005)
}
