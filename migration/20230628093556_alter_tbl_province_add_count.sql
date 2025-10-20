-- +migrate Up
ALTER TABLE province ADD COLUMN total_city int;
ALTER TABLE province ADD COLUMN total_district int;
ALTER TABLE province ADD COLUMN total_village int;

-- +migrate Down
ALTER TABLE province DROP COLUMN total_city;
ALTER TABLE province DROP COLUMN total_district;
ALTER TABLE province DROP COLUMN total_village;