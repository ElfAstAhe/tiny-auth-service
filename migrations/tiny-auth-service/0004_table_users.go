package tiny_auth_service

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

const (
	sql0004001AddColumn string = `
alter table if exists users add column if not exists user_type varchar(50) null default 'guest'
`
	sql0004001DropColumn string = `
alter table if exists users drop column if exists user_type
`
)

func up0004(ctx context.Context, db *sql.DB) error {
	return alterTableUsersAddColumnUserType(ctx, db)
}

func down0004(ctx context.Context, db *sql.DB) error {
	return alterTableUsersDropColumnUserType(ctx, db)
}

func alterTableUsersAddColumnUserType(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sql0004001AddColumn)

	return err
}

func alterTableUsersDropColumnUserType(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, sql0004001DropColumn)

	return err
}

func init() {
	goose.AddMigrationNoTxContext(up0004, down0004)
}
