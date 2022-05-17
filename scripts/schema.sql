CREATE TABLE `regions`
(
    `id` bigint unsigned PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `name` varchar(255),
    `short_name` varchar(255),
    `parent_id` bigint unsigned
);
ALTER TABLE `regions` ADD FOREIGN KEY (`parent_id`) REFERENCES `regions` (`id`);