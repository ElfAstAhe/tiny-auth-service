package tiny_auth_service

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
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
