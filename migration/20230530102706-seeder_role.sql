-- +migrate Up
insert into role(code, name, description, type, platform, is_active_flag,created_at,updated_at) values
('ROLE-0001', 'Administrator', 'System Administrator', 'Internal', 'Web Desktop', 'Y', NOW(), NOW());

-- +migrate Down
delete from role where code in ('ROLE-0001');