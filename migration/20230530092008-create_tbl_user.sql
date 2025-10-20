-- +migrate Up
CREATE TABLE IF NOT EXISTS "user"(
    id SERIAL NOT NULL PRIMARY KEY,
    code VARCHAR(20) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE,
    photo_url VARCHAR(255),
    encrypted_password VARCHAR(255),
    change_password_at TIMESTAMPTZ NOT NULL,
    tnc_accepted_at TIMESTAMPTZ NOT NULL,
    login_with_biometric_flag char(1) DEFAULT 'N',
    is_active_flag char(1) DEFAULT 'N',
    created_at TIMESTAMPTZ NOT NULL,
    created_by INT,
    updated_at TIMESTAMPTZ NOT NULL,
    updated_by INT,
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_user_email ON "user" (email);

-- +migrate Down
DROP TABLE IF EXISTS "user";