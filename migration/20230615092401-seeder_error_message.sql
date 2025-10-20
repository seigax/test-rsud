-- +migrate Up
insert into error_message(code, type, application_name, message, is_active_flag,created_at,updated_at) values
('1000', 'Error', 'Admin Panel', 'Forbidden', 'Y', NOW(), NOW()),
('1001', 'Error', 'Admin Panel', 'Invalid Parameter', 'Y', NOW(), NOW()),
('1002', 'Error', 'Admin Panel', 'Not Found', 'Y', NOW(), NOW()),
('1003', 'Error', 'Admin Panel', 'Internal Server Error', 'Y', NOW(), NOW()),
('1004', 'Error', 'Admin Panel', 'You need a valid token!', 'Y', NOW(), NOW()),
('1005', 'Error', 'Admin Panel', 'Email Already Registered', 'Y', NOW(), NOW()),
('1006', 'Error', 'Admin Panel', 'Phone Already Registered', 'Y', NOW(), NOW()),
('1007', 'Error', 'Admin Panel', 'Wrong Email Or Password', 'Y', NOW(), NOW()),
('1008', 'Error', 'Admin Panel', 'Token Expired!', 'Y', NOW(), NOW()),
('1009', 'Error', 'Admin Panel', 'Only Internal User Can Login!', 'Y', NOW(), NOW());

-- +migrate Down