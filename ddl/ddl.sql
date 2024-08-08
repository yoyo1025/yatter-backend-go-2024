-- テーブルスキーマの定義
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

CREATE TABLE `attachment_binding` (
  `attachment_id` BIGINT(20) NOT NULL,
  `status_id` BIGINT(20) NOT NULL,
  PRIMARY KEY (`attachment_id`, `status_id`),
  FOREIGN KEY (`attachment_id`) REFERENCES `attachment`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`status_id`) REFERENCES `status`(`id`) ON DELETE CASCADE
);

-- 初期データの挿入
INSERT INTO `account` (`username`, `password_hash`, `display_name`, `avatar`, `header`, `note`)
VALUES 
('user1', 'hashed_password_1', 'User One', 'avatar1.png', 'header1.png', 'This is user 1'),
('user2', 'hashed_password_2', 'User Two', 'avatar2.png', 'header2.png', 'This is user 2'),
('user3', 'hashed_password_3', 'User Three', 'avatar3.png', 'header3.png', 'This is user 3');

INSERT INTO `status` (`account_id`, `content`)
VALUES 
(1, 'This is the first status of user 1'),
(2, 'This is the first status of user 2'),
(3, 'This is the first status of user 3');

INSERT INTO `attachment` (`type`, `url`, `description`)
VALUES 
('image', 'https://example.com/image1.png', 'First image attachment'),
('video', 'https://example.com/video1.mp4', 'First video attachment'),
('image', 'https://example.com/image2.png', 'Second image attachment');

INSERT INTO `relationship` (`follower_id`, `followee_id`)
VALUES 
(1, 2),
(2, 3),
(3, 1);

INSERT INTO `attachment_binding` (`attachment_id`, `status_id`)
VALUES 
(1, 1),
(2, 2),
(3, 3);
