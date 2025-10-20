-- +migrate Up 
CREATE TABLE user_verification (
   id SERIAL NOT NULL PRIMARY KEY,
   verification_code VARCHAR(20) NOT NULL UNIQUE,
   user_id integer,
   verification_type_id integer,
   communication_device_type_code VARCHAR(50),
   verification_status_flag CHAR(1),
   expired_at TIMESTAMP,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   created_by integer,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_by integer,
   deleted_at TIMESTAMP
);

-- +migrate Down
DROP TABLE IF EXISTS user_verification;