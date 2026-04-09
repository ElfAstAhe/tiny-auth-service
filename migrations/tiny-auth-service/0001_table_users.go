package tiny_auth_service

import (
	"context"
	"database/sql"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/pressly/goose/v3"
)

func up0001(ctx context.Context, db *sql.DB) error {
	if err := createTableUsers(ctx, db); err != nil {
		return err
	}
	if err := createIndexUsersAlive(ctx, db); err != nil {
		return err
	}

	return nil
}

func createTableUsers(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sqlCreateTableUsers)
	if err != nil {
		return errs.NewDBMigrationError("create table users", err)
	}

	return nil
}

func createIndexUsersAlive(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sqlCreateIndexUsersAlive)
	if err != nil {
		return errs.NewDBMigrationError("create index users alive", err)
	}

	return nil
}

func down0001(ctx context.Context, db *sql.DB) error {
	if err := dropIndexUsersAlive(ctx, db); err != nil {
		return err
	}
	if err := dropTableUsers(ctx, db); err != nil {
		return err
	}

	return nil
}

func dropTableUsers(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sqlDropTableUsers)
	if err != nil {
		return errs.NewDBMigrationError("drop table users", err)
	}

	return nil
}

func dropIndexUsersAlive(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sqlDropIndexUsersAlive)
	if err != nil {
		return errs.NewDBMigrationError("drop index users alive", err)
	}

	return nil
}

func init() {
	goose.AddMigrationNoTxContext(up0001, down0001)
}
