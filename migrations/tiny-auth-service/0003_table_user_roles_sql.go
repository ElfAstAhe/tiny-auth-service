package tiny_auth_service

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
