package auth

import (
	"context"
	"database/sql"

	libextmigrpg "github.com/ElfAstAhe/go-service-template/pkg/ext/migrations/postgres"
	"github.com/pressly/goose/v3"
)

func up0006(ctx context.Context, db *sql.DB) error {
	if err := libextmigrpg.CreateTableDataAuditOutbox(ctx, db); err != nil {
		return err
	}
	if err := libextmigrpg.CreateIndexDataAuditOutboxProcessed(ctx, db); err != nil {
		return err
	}

	return libextmigrpg.CreateIndexDataAuditOutboxPending(ctx, db)
}

func down0006(ctx context.Context, db *sql.DB) error {
	if err := libextmigrpg.DropIndexDataAuditOutboxPending(ctx, db); err != nil {
		return err
	}
	if err := libextmigrpg.DropIndexDataAuditOutboxProcessed(ctx, db); err != nil {
		return err
	}

	return libextmigrpg.DropTableDataAuditOutbox(ctx, db)
}

func init() {
	goose.AddMigrationNoTxContext(up0006, down0006)
}
