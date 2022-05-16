CREATE TABLE `region`
(
    `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `name` varchar(255),
    `short_name` varchar(255),
    `parent_id` bigint
);
ALTER TABLE `region` ADD FOREIGN KEY (`parent_id`) REFERENCES `region` (`id`);