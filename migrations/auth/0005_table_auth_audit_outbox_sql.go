package auth

const (
	sqlCreateTableAuthAuditOutbox string = `
create table if not exists auth_audit_outbox (
    id varchar(50) not null,
    username varchar(100) not null,
    content_type varchar(50) not null,
    payload bytea null,
    status varchar(50) not null default 'pending',
    created_at timestamptz null default now(),
    update_at timestamptz null,
    --
    constraint auth_audit_outbox_pk primary key(id)
)
`
	sqlDropTableAuthAuditOutbox string = `
drop table if exists auth_audit_outbox cascade
`
	sqlCreateIndexAuthAuditOutboxPending string = `
create index if not exists idx_auth_audit_outbox_pending on auth_audit_outbox (status asc, created_at asc) where status = 'PENDING'
`
	sqlDropIndexAuthAuditOutboxPending string = `
drop index if exists idx_auth_audit_outbox_pending
`
	sqlCreateIndexAuthAuditOutboxProcessed string = `
create index if not exists idx_auth_audit_outbox_processed on auth_audit_outbox (status asc, created_at asc) where status in ('SUCCESS', 'FAILED')
`
	sqlDropIndexAuthAuditOutboxProcessed string = `
drop index if exists idx_auth_audit_outbox_processed
`
)
