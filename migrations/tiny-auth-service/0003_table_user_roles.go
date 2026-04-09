package tiny_auth_service

import (
	"context"
	"database/sql"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/pressly/goose/v3"
)

func up0003(ctx context.Context, db *sql.DB) error {
	if err := createTableUserRoles(ctx, db); err != nil {
		return errs.NewDBMigrationError("create user roles table", err)
	}
	if err := createIndexUserRoles(ctx, db); err != nil {
		return errs.NewDBMigrationError("create user roles index", err)
	}

	return nil
}

func createTableUserRoles(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sqlCreateTableUserRoles)
	if err != nil {
		return errs.NewDBMigrationError("create user roles table", err)
	}

	return nil
}

func createIndexUserRoles(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sqlCreateIndexUserRoles)
	if err != nil {
		return errs.NewDBMigrationError("create user roles index", err)
	}

	return nil
}

func down0003(ctx context.Context, db *sql.DB) error {
	if err := dropIndexUserRoles(ctx, db); err != nil {
		return errs.NewDBMigrationError("drop user roles index", err)
	}
	if err := dropTableUserRoles(ctx, db); err != nil {
		return errs.NewDBMigrationError("drop user roles table", err)
	}

	return nil
}

func dropTableUserRoles(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sqlDropTableUserRoles)
	if err != nil {
		return errs.NewDBMigrationError("drop user roles table", err)
	}

	return nil
}

func dropIndexUserRoles(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sqlDropIndexUserRoles)
	if err != nil {
		return errs.NewDBMigrationError("drop user roles index", err)
	}

	return nil
}

func init() {
	goose.AddMigrationNoTxContext(up0003, down0003)
}
