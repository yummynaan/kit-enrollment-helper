
-- +migrate Up
ALTER TABLE `courses` MODIFY `timetable_id` VARCHAR(16) NOT NULL;
ALTER TABLE `courses` MODIFY `class` VARCHAR(16) NOT NULL;
ALTER TABLE `courses` MODIFY `day` VARCHAR(16) NOT NULL;
-- +migrate Down
