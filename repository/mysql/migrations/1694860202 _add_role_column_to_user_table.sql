
-- MYSQL 8.0 set the role value to `user for all old records
-- Don't change order of the enum values
-- TODO find a better solution instead of keeping order
-- +migrate Up
ALTER TABLE `users` ADD COLUMN `role` ENUM('user','admin') NOT NULL;

-- +migrate Down
ALTER TABLE `users` DROP COLUMN `role`;