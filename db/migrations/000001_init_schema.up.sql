CREATE TABLE `teachers` (
  `id`  bigint PRIMARY KEY auto_increment,
  `email` varchar(255) NOT NULL
);

CREATE TABLE `students` (
  `id` bigint PRIMARY KEY auto_increment,
  `email` varchar(255) NOT NULL,
  `status` ENUM ('suspended', 'active') NOT NULL DEFAULT "active"
);

CREATE TABLE `registers` (
  `t_id` bigint NOT NULL,
  `s_id` bigint NOT NULL,
  PRIMARY KEY (`t_id`, `s_id`)
);

CREATE INDEX `teachers_index_0` ON `teachers` (`email`);

CREATE INDEX `students_index_1` ON `students` (`email`);

ALTER TABLE `registers` ADD FOREIGN KEY (`t_id`) REFERENCES `teachers` (`id`);

ALTER TABLE `registers` ADD FOREIGN KEY (`s_id`) REFERENCES `students` (`id`);
