package tiny_auth_service

const (
	sqlCreateTableUsers string = `
create table if not exists users (
    id varchar(50) not null,
    name varchar(100) not null,
    password_hash varchar(100) not null default '',
    public_key text not null default '',
    private_key text not null default '',
    active bool not null default true,
    deleted bool not null default false,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    constraint users_pk primary key (id),
    constraint users_uk unique (name)
)
`
	sqlDropTableUsers string = `drop table if exists users`

	sqlCreateIndexUsersAlive string = `create index if not exists idx_users_alive on users (deleted asc, active desc, name asc)`
	sqlDropIndexUsersAlive   string = `drop index if exists idx_users_alive`
)
