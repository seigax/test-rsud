-- +migrate Up
insert into permission(id, name, description,created_at,updated_at) values
(1, 'Dashboard:ACCESS', 'access dashboard menu', NOW(), NOW()),
(2, 'Role:ACCESS', 'access role menu', NOW(), NOW()),
(3, 'Permission:ACCESS', 'access permission menu', NOW(), NOW()),
(4, 'Region:ACCESS', 'access region menu', NOW(), NOW());

-- +migrate Down