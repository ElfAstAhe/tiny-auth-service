package auth

const (
	sqlCreateTableLoginAttemptsOutbox string = `
create table if not exists login_attempts_outbox (
    id varchar(50) not null,
    username varchar(100) not null,
    content_type varchar(50) not null,
    payload bytea null,
    status varchar(50) not null default 'pending',
    created_at timestamptz null default now(),
    update_at timestamptz null,
    --
    constraint login_attempts_outbox_pk primary key(id)
)
`
	sqlDropTableLoginAttemptsOutbox string = `
drop table if exists login_attempts_outbox cascade
`
	sqlCreateIndexLoginAttemptsOutboxPending string = `
create index if not exists idx_login_attempts_outbox_pending on login_attempts_outbox (status asc, created_at asc) where status = 'pending'
`
	sqlDropIndexLoginAttemptsOutboxPending string = `
drop index if exists idx_login_attempts_outbox_pending
`
)
