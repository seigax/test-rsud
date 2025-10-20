-- +migrate Up
insert into "user"(id, code, name, email, encrypted_password, is_active_flag,created_at,updated_at, change_password_at, tnc_accepted_at) values
(1, 'USR-0000000000001', 'System Administrator', 'sysadmin@gmail.com', '$2a$04$dRCgjLXf7Inx.S07uwKw0egXbgonDwlLNV1Lpm1pPkokMpOyHzatC', 'Y', NOW(), NOW(), NOW(), NOW());

insert into user_role(user_id, role_id, created_at, updated_at) values
(1, 1, NOW(), NOW());

insert into user_phone(user_id, phone_number, is_active_flag, created_at, updated_at) values
(1, '082299999999', 'Y', NOW(), NOW());

-- +migrate Down
delete from "user" where id in (1);
delete from user_role where user_id in (1);
delete from user_phone where user_id in (1);