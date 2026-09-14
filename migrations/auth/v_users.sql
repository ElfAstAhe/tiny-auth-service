-- ATTENTION! Do not need in application, for manual usage only
create or replace view auth_db.v_users as (
select
    ur.user_id,
	u.name as username,
	u.active as user_active,
	u.deleted as user_deleted,
	ur.role_id,
	r.name as role_name,
	r.deleted as role_deleted
from
    auth_db.user_roles ur
	left outer join
	    auth_db.users u
		on
		    u.id = ur.user_id
	left outer join
	    auth_db.roles r
		on
		    r.id = ur.role_id
)
;
