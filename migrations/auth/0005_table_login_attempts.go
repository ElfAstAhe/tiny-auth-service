package auth

import (
	"context"
	"database/sql"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/pressly/goose/v3"
)

func up0005(ctx context.Context, db *sql.DB) error {
	if err := createTableLoginAttempts(ctx, db); err != nil {
		return err
	}

	return createIndexLoginAttemptsUnsent(ctx, db)
}

func createTableLoginAttempts(ctx context.Context, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, sqlCreateTableLoginAttempts); err != nil {
		return errs.NewDBMigrationError("create table login attempts failed", err)
	}

	return nil
}

func createIndexLoginAttemptsUnsent(ctx context.Context, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, sqlCreateIndexLoginAttemptsUnsent); err != nil {
		return errs.NewDBMigrationError("create index login attempts unsent failed", err)
	}

	return nil
}

func down0005(ctx context.Context, db *sql.DB) error {
	if err := dropIndexLoginAttemptsUnsent(ctx, db); err != nil {
		return err
	}

	return dropTableLoginAttempts(ctx, db)
}

func dropTableLoginAttempts(ctx context.Context, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, sqlDropTableLoginAttempts); err != nil {
		return errs.NewDBMigrationError("drop table login attempts failed", err)
	}

	return nil
}

func dropIndexLoginAttemptsUnsent(ctx context.Context, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, sqlDropIndexLoginAttemptsUnsent); err != nil {
		return errs.NewDBMigrationError("drop index login attempts unsent failed", err)
	}

	return nil
}

func init() {
	goose.AddMigrationNoTxContext(up0005, down0005)
}
