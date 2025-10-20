-- +migrate Up
CREATE TABLE IF NOT EXISTS system_parameter(
    id SERIAL NOT NULL PRIMARY KEY,
    code VARCHAR(20) NOT NULL UNIQUE,
    parameter_name VARCHAR(255) NOT NULL,
    data_type VARCHAR(30),
    message VARCHAR(255),
    is_active_flag char(1) DEFAULT 'N',
    created_at TIMESTAMPTZ NOT NULL,
    created_by INT,
    updated_at TIMESTAMPTZ NOT NULL,
    updated_by INT,
    deleted_at TIMESTAMPTZ
);

-- +migrate Down
DROP TABLE IF EXISTS system_parameter;