package postgres

const (
	sqlUserRolesListAll string = `
select
    r.id,
    r.name,
    r.description,
    r.deleted,
    r.created_at,
    r.updated_at
from
    user_roles ur
    inner join 
        roles r
        on
            r.id = ur.role_id
        and r.deleted = false
where
    ur.user_id = $1
`
	sqlUserRolesListAllByOwners string = `
select
    ur.user_id,
    r.id,
    r.name,
    r.description,
    r.deleted,
    r.created_at,
    r.updated_at
from
    user_roles ur
    inner join 
        roles r
        on
            r.id = ur.role_id
        and r.deleted = false
where
    ur.user_id = any($1)
order by
    1 asc, 2 asc
`
	sqlUserRolesDeleteAll string = `
delete from
    user_roles
where
    user_id = $1
`
	sqlUserRolesCreate string = `
insert into user_roles(
    user_id,
    role_id
)
values ($1, $2)
returning role_id
`
)
