-- +migrate Up
CREATE TABLE IF NOT EXISTS error_message(
    id SERIAL NOT NULL PRIMARY KEY,
    code VARCHAR(20) NOT NULL UNIQUE,
    type VARCHAR(50),
    application_name VARCHAR(255) NOT NULL,
    message VARCHAR(255),
    is_active_flag char(1) DEFAULT 'N',
    created_at TIMESTAMPTZ NOT NULL,
    created_by INT,
    updated_at TIMESTAMPTZ NOT NULL,
    updated_by INT,
    deleted_at TIMESTAMPTZ
);

-- +migrate Down
DROP TABLE IF EXISTS error_message;