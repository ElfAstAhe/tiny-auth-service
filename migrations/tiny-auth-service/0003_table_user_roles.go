package tiny_auth_service

import (
	"context"
	"database/sql"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/pressly/goose/v3"
)

const (
	sqlCreateTableUserRoles string = `
create table if not exists user_roles (
    user_id varchar(50) not null,
    role_id varchar(50) not null,
    created_at timestamptz not null default now(),
    constraint user_roles_fk_user foreign key (user_id) references users (id),
    constraint user_roles_fk_role foreign key (role_id) references roles (id),
    constraint user_roles_uk unique (user_id, role_id)
)
`
	sqlDropTableUserRoles string = `drop table if exists user_roles`

	sqlCreateIndexUserRoles string = `create index if not exists idx_user_roles on user_roles(user_id asc, role_id asc)`
	sqlDropIndexUserRoles   string = `drop index if exists idx_user_roles`
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
