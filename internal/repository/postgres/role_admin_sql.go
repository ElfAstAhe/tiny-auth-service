package postgres

const (
	sqlRoleAdminFind string = `
select
    id,
    name,
    description,
    deleted,
    created_at,
    updated_at
from
    roles
where
    id = $1
`
	sqlRoleAdminFindByName string = `
select
    id,
    name,
    description,
    deleted,
    created_at,
    updated_at
from
    roles
where
    name = $1
`
	sqlRoleAdminList string = `
select
    id,
    name,
    description,
    deleted,
    created_at,
    updated_at
from
    roles
order by
    id asc
offset $2
limit $1
`
	sqlRoleAdminCreate string = `
insert into roles (
    id,
    name,
    description,
    deleted,
    created_at,
    updated_at
)
values ($1, $2, $3, false, $4, $5)
returning id, name, description, deleted, created_at, updated_at
`
	sqlRoleAdminChange string = `
update
    roles
set
    name = $2,
    description = $3,
    deleted = $4,
    updated_at = $5
where
    id = $1
returning id, name, description, deleted, created_at, updated_at
`
	sqlRoleAdminDelete string = `
delete
from
    roles
where
    id = $1
`
)
