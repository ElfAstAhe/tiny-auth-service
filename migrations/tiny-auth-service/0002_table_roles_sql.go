package tiny_auth_service

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
