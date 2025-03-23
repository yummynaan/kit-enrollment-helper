
-- +migrate Up
ALTER TABLE `courses` DROP INDEX `courses_idx`;
ALTER TABLE `courses` ADD UNIQUE INDEX `courses_idx` (`timetable_id`, `title`, `class`, `year`, `semester`, `day`, `syllabus_year`);
-- +migrate Down
