-- +migrate Up
CREATE TABLE IF NOT EXISTS role_menu(
    id SERIAL NOT NULL PRIMARY KEY,
    role_id int NOT NULL,
    menu_id int NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    created_by INT,
    updated_at TIMESTAMPTZ NOT NULL,
    updated_by INT,
    deleted_at TIMESTAMPTZ
);

-- +migrate Down
DROP TABLE IF EXISTS role_menu;