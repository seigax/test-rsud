-- +migrate Up 
SELECT setval('role_id_seq', 999);

-- +migrate Down

