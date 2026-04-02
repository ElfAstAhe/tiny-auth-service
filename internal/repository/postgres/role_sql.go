package postgres

// SQL запросы
const (
	sqlRoleFind string = `
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
	sqlRoleFindByName string = `
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
	sqlRoleList string = `
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
	sqlRoleCreate string = `
insert into roles (
    id,
    name,
    description,
    deleted,
    created_at,
    updated_at
)
values($1, $2, $3, false, $4, $5)
returning id, name, description, deleted, created_at, updated_at
`
	sqlRoleChange string = `
update
    roles
set
    name = $2,
    description = $3,
    updated_at = $4
where
    id = $1
returning id, name, description, deleted, created_at, updated_at
`
	sqlRoleDelete string = `
update
    roles
set
    deleted = true
where
    id = $1
`
)
