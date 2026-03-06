package tiny_auth_service

import (
	"context"
	"database/sql"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/pressly/goose/v3"
)

const (
	sqlCreateTableUsers string = `
create table if not exists users (
    id varchar(50) not null,
    name varchar(100) not null,
    password_hash varchar(100) not null,
    active bool null default true,
    deleted bool null default false,
    created_at timestamptz default now(),
    updated_at timestamptz default now(),
    constraint users_pk primary key (id),
    constraint users_uk unique (name)
)
`
	sqlDropTableUsers string = `drop table if exists users`

	sqlCreateIndexUsersAlive string = `create index if not exists idx_users_alive on users (deleted asc, active desc, name asc)`
	sqlDropIndexUsersAlive   string = `drop index if exists idx_users_alive`
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
