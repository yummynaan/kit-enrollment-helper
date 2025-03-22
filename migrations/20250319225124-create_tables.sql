
-- +migrate Up
CREATE TABLE IF NOT EXISTS `departments` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `faculty` VARCHAR(32) NOT NULL,
  `field` VARCHAR(32) NOT NULL,
  `program` VARCHAR(32) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `departments_idx` (`faculty`, `field`, `program`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `users` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(50) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `password_hash` VARCHAR(255) NOT NULL,
  `enrollment_year` INT NOT NULL,
  `department_id` INT UNSIGNED NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `users_idx` (`email`),
  FOREIGN KEY (`department_id`) REFERENCES `departments` (`id`) ON DELETE RESTRICT
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `courses` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `timetable_id` VARCHAR(16),
  `title` VARCHAR(64) NOT NULL,
  `class` VARCHAR(16),
  `type` VARCHAR(16) NOT NULL,
  `credits` INT NOT NULL,
  `instructors` VARCHAR(256) NOT NULL,
  `year` VARCHAR(16) NOT NULL,
  `semester` VARCHAR(16) NOT NULL,
  `day` VARCHAR(16),
  PRIMARY KEY (`id`),
  UNIQUE INDEX `courses_idx` (`title`, `class`, `year`, `semester`, `day`)
) ENGINE = InnoDB;
-- +migrate Down
