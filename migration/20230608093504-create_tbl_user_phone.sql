-- +migrate Up
CREATE TABLE IF NOT EXISTS user_phone(
    id SERIAL NOT NULL PRIMARY KEY,
    user_id INT,
    phone_number varchar(20) NOT NULL UNIQUE,
    is_active_flag char(1) DEFAULT 'N',
    created_at TIMESTAMPTZ NOT NULL,
    created_by INT,
    updated_at TIMESTAMPTZ NOT NULL,
    updated_by INT,
    deleted_at TIMESTAMPTZ
);

-- +migrate Down
DROP TABLE IF EXISTS user_phone;