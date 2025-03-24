
-- +migrate Up
ALTER TABLE `users` DROP COLUMN `password_hash`;
-- +migrate Down
