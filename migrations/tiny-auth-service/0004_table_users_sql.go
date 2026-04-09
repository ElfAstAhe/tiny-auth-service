package tiny_auth_service

const (
	sql0004001AddColumn string = `
alter table if exists users add column if not exists user_type varchar(50) null default 'guest'
`
	sql0004001DropColumn string = `
alter table if exists users drop column if exists user_type
`
)
