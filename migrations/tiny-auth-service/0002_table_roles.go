package tiny_auth_service

import (
	"context"
	"database/sql"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/pressly/goose/v3"
)

func up0002(ctx context.Context, db *sql.DB) error {
	if err := createTableRoles(ctx, db); err != nil {
		return err
	}
	if err := createIndexRolesAlive(ctx, db); err != nil {
		return err
	}

	return nil
}

func createTableRoles(ctx context.Context, db *sql.DB) error {
	_, err := db.Exec(sqlCreateTableRoles)
	if err != nil {
		return errs.NewDBMigrationError("create table roles", err)
	}

	return nil
}

func createIndexRolesAlive(ctx context.Context, db *sql.DB) error {
	_, err := db.Exec(sqlCreateIndexRolesActive)
	if err != nil {
		return errs.NewDBMigrationError("create index roles active", err)
	}

	return nil
}

func down0002(ctx context.Context, db *sql.DB) error {
	if err := dropIndexRolesActive(ctx, db); err != nil {
		return err
	}
	if err := dropTableRoles(ctx, db); err != nil {
		return err
	}

	return nil
}

func dropTableRoles(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sqlDropTableRoles)
	if err != nil {
		return errs.NewDBMigrationError("drop table roles", err)
	}

	return nil
}

func dropIndexRolesActive(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sqlDropIndexRolesActive)
	if err != nil {
		return errs.NewDBMigrationError("drop index roles active", err)
	}

	return nil
}

func init() {
	goose.AddMigrationNoTxContext(up0002, down0002)
}
