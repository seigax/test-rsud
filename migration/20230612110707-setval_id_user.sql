-- +migrate Up 
SELECT setval('user_id_seq', (SELECT max(id) FROM "user"));

-- +migrate Down

