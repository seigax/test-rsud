-- +migrate Up
insert into menu(id, code, parent_menu_id, name, level, url, icon, order_number, is_active_flag,created_at,updated_at) values
(1, 'MENU-0001', 0, 'Dashboard', 'Parent', '/dashboard', '', 1, 'Y', NOW(), NOW()),
(2, 'MENU-0002', 0, 'Administrasi', 'Parent', '/administrasi', '', 2, 'Y', NOW(), NOW()),
(3, 'MENU-0003', 0, 'Master Data', 'Parent', '/master-data', '', 4, 'Y', NOW(), NOW()),
(4, 'MENU-0004', 2, 'Sistem', 'Child', '/sistem', '', 1, 'Y', NOW(), NOW()),
(5, 'MENU-0005', 2, 'Keamanan', 'Child', '/keamanan', '', 2, 'Y', NOW(), NOW()),
(6, 'MENU-0006', 6, 'Parameter System', 'Child', '/parameter-sistem', '', 1, 'Y', NOW(), NOW()),
(7, 'MENU-0007', 6, 'Pesan Kesalahan', 'Child', '/pesan-kesalahan', '', 2, 'Y', NOW(), NOW()),
(8, 'MENU-0008', 7, 'Menu', 'Child', '/menu', '', 1, 'Y', NOW(), NOW()),
(9, 'MENU-0009', 7, 'Role', 'Child', '/role', '', 2, 'Y', NOW(), NOW()),
(10, 'MENU-0010', 7, 'User', 'Child', '/user', '', 3, 'Y', NOW(), NOW()),
(11, 'MENU-0011', 4, 'Wilayah', 'Child', '/wilayah', '', 1, 'Y', NOW(), NOW()),
(12, 'MENU-0012', 4, 'Pusat Bantuan', 'Child', '/pusat-bantuan', '', 2, 'Y', NOW(), NOW()),
(13, 'MENU-0013', 2, 'Audit Trail', 'Child', '/audit-trail', '', 2, 'Y', NOW(), NOW())
;

SELECT setval('menu_id_seq', (SELECT max(id) FROM "menu"));

-- +migrate Down
delete from menu where id in (1,2,3,4,5,6,7,8,9,10,11,12,13);