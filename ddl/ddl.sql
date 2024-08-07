CREATE TABLE `account` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL UNIQUE,
  `password_hash` varchar(255) NOT NULL,
  `display_name` varchar(255),
  `avatar` text,
  `header` text,
  `note` text,
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);

CREATE TABLE `status` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
  `account_id` BIGINT(20) NOT NULL,
  `content` VARCHAR(255),
  `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`account_id`) REFERENCES `account`(`id`) ON DELETE CASCADE
);

CREATE TABLE `attachment` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
  `type` VARCHAR(255) NOT NULL,
  `url` TEXT NOT NULL,
  `description` TEXT,
  PRIMARY KEY (`id`)
);

CREATE TABLE `relationship` (
  `follower_id` BIGINT(20) NOT NULL,
  `followee_id` BIGINT(20) NOT NULL,
  PRIMARY KEY (`follower_id`, `followee_id`),
  FOREIGN KEY (`follower_id`) REFERENCES `account`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`followee_id`) REFERENCES `account`(`id`) ON DELETE CASCADE
);
