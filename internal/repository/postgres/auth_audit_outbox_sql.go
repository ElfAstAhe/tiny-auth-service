package postgres

const (
	SqlAuthAuditOutboxFind string = `
select
    id,
    username,
    content_type,
    payload,
    created_at,
    updated_at
from
    auth_audit_outbox
where
    id = $1
`
	SqlAuthAuditOutboxList string = `
select
    id,
    username,
    content_type,
    payload,
    created_at,
    updated_at
from
    auth_audit_outbox
order by
    created_at asc,
    id asc
offset $2
limit $1
`
	SqlAuthAuditOutboxCreate string = `
insert into auth_audit_outbox (
    id,
    username,
    content_type,
    payload,
    status,
    created_at
)
values (
    $1, $2, $3, $4, $5, $6
)
`
	SqlAuthAuditOutboxChange string = `
update
    auth_audit_outbox
set
    username = $2,
    content_type = $3,
    payload = $4,
    status = $5,
    updated_at = $6
where
    id = $1
`
	SqlAuthAuditOutboxDelete string = `
delete
from
    auth_audit_outbox
where
    id = $1
`
	SqlAuthAuditOutboxGetTail string = `
select
    id
from
    auth_audit_aoutbox
where
    status = any($1)
and updated_at < $2
limit $3
`
	SqlAuthAuditOutboxListPending string = `
select
    id,
    username,
    content_type,
    payload,
    status,
    created_at,
    updated_at
from
    auth_audit_outbox
where
    status = any($1)
order by
    created_at asc
limit $2
`
	SqlAuthAuditOutboxListProcessed string = `
select
    id,
    username,
    content_type,
    payload,
    status,
    created_at,
    updated_at
from
    auth_audit_outbox
where
    status = any($1)
and updated_at < $2
limit $3
`
	SqlAuthAuditOutboxChangeStatus string = `
update
    auth_audit_outbox
set
    status = $2,
    updated_at = $3
where
    id = $1
`
)
