package postgres

const (
	sqlUserAdminFind string = `
select
    id,
    name,
    user_type,
    password_hash,
    public_key,
    private_key,
    active,
    deleted,
    created_at,
    updated_at
from
    users
where
    id = $1
`
	sqlUserAdminFindByName string = `
select
    id,
    name,
    user_type,
    password_hash,
    public_key,
    private_key,
    active,
    deleted,
    created_at,
    updated_at
from
    users
where
    name = $1
`
	sqlUserAdminList string = `
select
    id,
    name,
    user_type,
    password_hash,
    public_key,
    private_key,
    active,
    deleted,
    created_at,
    updated_at
from
    users
order by
    id asc
offset $2
limit $1
`
	sqlUserAdminCreate string = `
insert into users (
    id,
    name,
    user_type,
    password_hash,
    public_key,
    private_key,
    active,
    deleted,
    created_at,
    updated_at
)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
returning id, name, user_type, password_hash, public_key, private_key, active, deleted, created_at, updated_at
`
	sqlUserAdminChange string = `
update
    users
set
    name = $2,
    user_type = $3,
    password_hash = $4,
    public_key = $5,
    private_key = $6,
    active = $7,
    deleted = $8,
    updated_at = $9
where
    id = $1
returning id, name, user_type, password_hash, public_key, private_key, active, deleted, created_at, updated_at
`
	sqlUserAdminDelete string = `
delete
from
    users
where
    id = $1
`
)
