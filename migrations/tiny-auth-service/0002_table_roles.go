package tiny_auth_service

import (
	"context"
	"database/sql"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/pressly/goose/v3"
)

const (
	sqlCreateTableRoles string = `
create table if not exists roles (
    id varchar(50) not null,
    name varchar(100) not null,
    description varchar(512) not null default '',
    deleted bool not null default false,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    constraint roles_pk primary key (id),
    constraint roles_uk unique(name)
)
`
	sqlDropTableRoles string = `drop table if exists roles`

	sqlCreateIndexRolesActive string = `create index if not exists idx_roles_active on roles (deleted asc, name asc)`
	sqlDropIndexRolesActive   string = `drop table if exists idx_roles_active`
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
