-- +migrate Up
CREATE TABLE IF NOT EXISTS menu(
    id SERIAL NOT NULL PRIMARY KEY,
    code VARCHAR(20) NOT NULL UNIQUE,
    parent_menu_id int,
    name VARCHAR(255),
    level VARCHAR(50),
    url VARCHAR(255) NOT NULL,
    icon VARCHAR(255),
    order_number int,
    is_active_flag char(1) DEFAULT 'N',
    created_at TIMESTAMPTZ NOT NULL,
    created_by INT,
    updated_at TIMESTAMPTZ NOT NULL,
    updated_by INT,
    deleted_at TIMESTAMPTZ
);

-- +migrate Down
DROP TABLE IF EXISTS menu;