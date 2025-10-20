-- +migrate Up
CREATE TABLE IF NOT EXISTS role_permission(
    id SERIAL NOT NULL PRIMARY KEY,
    id_role INT NOT NULL,
    id_permission INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    created_by INT,
    updated_at TIMESTAMPTZ NOT NULL,
    updated_by INT,
    deleted_at TIMESTAMPTZ
);

-- +migrate Down
DROP TABLE IF EXISTS role_permission;