CREATE TABLE `teachers` (
  `email` varchar(255) PRIMARY KEY
);

CREATE TABLE `students` (
  `email` varchar(255) PRIMARY KEY,
  `status` ENUM ('suspended', 'active') NOT NULL DEFAULT "active"
);

CREATE TABLE `registers` (
  `t_email` varchar(255) NOT NULL,
  `s_email` varchar(255) NOT NULL
);

CREATE INDEX `teachers_index_0` ON `teachers` (`email`);

CREATE INDEX `students_index_1` ON `students` (`email`);

ALTER TABLE `registers` ADD FOREIGN KEY (`t_email`) REFERENCES `teachers` (`email`);

ALTER TABLE `registers` ADD FOREIGN KEY (`s_email`) REFERENCES `students` (`email`);
