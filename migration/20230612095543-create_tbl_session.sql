-- +migrate Up
CREATE TABLE IF NOT EXISTS session(
    id SERIAL NOT NULL PRIMARY KEY,
    user_id int NOT NULL,
    token TEXT NOT NULL UNIQUE,
    is_login_with_biometric_flag char(1) DEFAULT 'N',
    created_at TIMESTAMPTZ NOT NULL,
    created_by INT,
    updated_at TIMESTAMPTZ NOT NULL,
    updated_by INT,
    deleted_at TIMESTAMPTZ,
    expired_at TIMESTAMPTZ
);

-- +migrate Down
DROP TABLE IF EXISTS session;