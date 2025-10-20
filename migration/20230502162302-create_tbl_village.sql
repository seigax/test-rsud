-- +migrate Up
CREATE TABLE IF NOT EXISTS village(
    id SERIAL NOT NULL PRIMARY KEY,
    district_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(20) NOT NULL UNIQUE,
    postal_code VARCHAR(10),
    is_active_flag char(1) DEFAULT 'N',
    created_at TIMESTAMPTZ NOT NULL,
    created_by INT,
    updated_at TIMESTAMPTZ NOT NULL,
    updated_by INT,
    deleted_at TIMESTAMPTZ
);

-- +migrate Down
DROP TABLE IF EXISTS village;