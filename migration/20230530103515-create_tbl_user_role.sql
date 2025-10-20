-- +migrate Up
CREATE TABLE IF NOT EXISTS user_role(
    id SERIAL NOT NULL PRIMARY KEY,
    user_id INT,
    role_id INT,
    created_at TIMESTAMPTZ NOT NULL,
    created_by INT,
    updated_at TIMESTAMPTZ NOT NULL,
    updated_by INT,
    deleted_at TIMESTAMPTZ
);

-- +migrate Down
DROP TABLE IF EXISTS user_role;