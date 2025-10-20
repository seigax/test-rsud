-- +migrate Up
CREATE TABLE verification_type (
   id SERIAL NOT NULL PRIMARY KEY,
   code VARCHAR(20) NOT NULL UNIQUE,
   name varchar(255),
   "desc" text,
   is_active_flag boolean,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   created_by integer,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_by integer,
   deleted_at TIMESTAMP
);

-- +migrate Down
DROP TABLE IF EXISTS verification_type;