-- +migrate Up
insert into role_menu(id, role_id, menu_id, created_at, updated_at) values
(1, 1,  1,  NOW(), NOW()),
(2, 1,  2,  NOW(), NOW()),
(3, 1,  6,  NOW(), NOW()),
(4, 1,  7,  NOW(), NOW()),
(5, 1,  27,  NOW(), NOW()),
(6, 1,  8,  NOW(), NOW()),
(7, 1,  9,  NOW(), NOW()),
(8, 1,  10,  NOW(), NOW()),
(9, 1,  11,  NOW(), NOW()),
(10, 1,  12,  NOW(), NOW())
;

SELECT setval('role_menu_id_seq', (SELECT max(id) FROM "role_menu"));

-- +migrate Down
delete from role_menu where id in (1,2,3,4,5,6,7,8,9,10);