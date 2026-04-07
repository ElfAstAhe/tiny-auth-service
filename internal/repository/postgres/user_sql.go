package postgres

const (
	sqlUserFind string = `
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
	sqlUserFindByName string = `
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
	sqlUserList string = `
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
	sqlUserCreate string = `
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
values($1, $2, $3, $4, $5, $6, $7, false, $8, $9)
returning id, name, user_type, password_hash, public_key, private_key, active, deleted, created_at, updated_at
`
	sqlUserChange string = `
update
    users
set
    user_type = $2,
    password_hash = $3,
    public_key = $4,
    private_key = $5,
    active = $6,
    deleted = $7,
    updated_at = $8
where
    id = $1
returning id, name, user_type, password_hash, public_key, private_key, active, deleted, created_at, updated_at
`
	sqlUserDelete string = `
update
    users
set
    deleted = true
where
    id = $1
`
)
