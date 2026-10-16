package auth

const (
	sqlCreateTableLoginAttempts string = `
create table if not exists login_attempts (
    id varchar(50) not null,
    content_type varchar(50) not null,
    data text not null,
    created_at timestamptz null default now(),
    update_at timestamptz null,
    sent bool not null default false,
    --
    constraint login_attempts_pk primary key(id)
)
`
	sqlDropTableLoginAttempts string = `
drop table if exists login_attempts cascade
`
	sqlCreateIndexLoginAttemptsUnsent string = `
create index if not exists idx_login_attempts_unsent on login_attempts (created_at asc, sent asc)
`
	sqlDropIndexLoginAttemptsUnsent string = `
drop index if exists idx_login_attempts_unsent
`
)
