
-- +migrate Up
ALTER TABLE `courses` ADD `syllabus_year` INT NOT NULL;
-- +migrate Down
