-- +migrate Up
ALTER TABLE users ADD COLUMN `password` text NOT NULL;

-- +migrate Down
ALTER TABLE `users` DROP COLUMN `password`;
