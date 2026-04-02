package postgres

const (
	sqlUserRolesAdminFind string = `
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
where
    ur.user_id = $1
and ur.role_id = $2
`
	sqlUserRolesAdminListAll string = `
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
where
    ur.user_id = $1
`
	sqlUserRolesAdminListAllByOwners string = `
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
where
    ur.user_id = any($1)
`
	sqlUserRolesAdminCreate string = `
insert into user_roles (
    user_id,
    role_id
)
values ($1, $2)
returning role_id
`
	sqlUserRolesAdminDelete string = `
delete from user_roles where user_id = $1 and role_id = $2
`
	sqlUserRolesAdminDeleteAll string = `
delete from user_roles where user_id = $1
`
)
