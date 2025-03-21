-- Tạo các database trước
CREATE DATABASE IF NOT EXISTS `thinkflow-auth`;
CREATE DATABASE IF NOT EXISTS `thinkflow-users`;
CREATE DATABASE IF NOT EXISTS `thinkflow-notes`;
CREATE DATABASE IF NOT EXISTS `thinkflow-media`;
CREATE DATABASE IF NOT EXISTS `thinkflow-gen`;
CREATE DATABASE IF NOT EXISTS `thinkflow-notifications`;

-- Bảng trong thinkflow-auth
CREATE TABLE `thinkflow-auth`.`auths` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT NOT NULL,
    `auth_type` ENUM('email_password', 'google', 'facebook') DEFAULT 'email_password',
    `email` VARCHAR(255) NOT NULL UNIQUE,
    `salt` VARCHAR(40) DEFAULT NULL,
    `password` VARCHAR(100) DEFAULT NULL,
    `facebook_id` VARCHAR(35) DEFAULT NULL,
    `google_id` VARCHAR(35) DEFAULT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_facebook_id` (`facebook_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Bảng trong thinkflow-users
CREATE TABLE `thinkflow-users`.`users` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `first_name` VARCHAR(30) NOT NULL,
    `last_name` VARCHAR(30) NOT NULL,
    `email` VARCHAR(255) NOT NULL UNIQUE,
    `phone` VARCHAR(30) DEFAULT NULL,
    `avatar_id` BIGINT DEFAULT NULL,
    `gender` ENUM('male', 'female', 'unknown') DEFAULT 'unknown',
    `dob` DATE DEFAULT NULL,
    `system_role` ENUM('sadmin', 'admin', 'user') DEFAULT 'user',
    `status` ENUM('active', 'waiting_verify', 'banned') DEFAULT 'active',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `idx_avatar_id` (`avatar_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Bảng trong thinkflow-notes
CREATE TABLE `thinkflow-notes`.`notes` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `title` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `thinkflow-notes`.`blocks` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `note_id` BIGINT NOT NULL,
    `type` ENUM('text', 'image', 'audio', 'summary', 'mindmap'),
    `position` INT NOT NULL,
    `text_id` BIGINT DEFAULT NULL,
    `image_id` BIGINT DEFAULT NULL,
    `audio_id` BIGINT DEFAULT NULL,
    `summary_id` BIGINT DEFAULT NULL,
    `mindmap_id` BIGINT DEFAULT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE INDEX `idx_note_id_position` (`note_id`, `position`),
    UNIQUE INDEX `idx_text_id` (`text_id`),
    UNIQUE INDEX `idx_image_id` (`image_id`),
    UNIQUE INDEX `idx_audio_id` (`audio_id`),
    UNIQUE INDEX `idx_summary_id` (`summary_id`),
    UNIQUE INDEX `idx_mindmap_id` (`mindmap_id`),
    INDEX `idx_note_id` (`note_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `thinkflow-notes`.`collaborations` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `note_id` BIGINT NOT NULL,
    `user_id` INT NOT NULL,
    `permission` ENUM('read', 'write') NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE INDEX `idx_note_id_user_id` (`note_id`, `user_id`),
    INDEX `idx_note_id` (`note_id`),
    INDEX `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Bảng trong thinkflow-media
CREATE TABLE `thinkflow-media`.`images` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `url` VARCHAR(2000) NOT NULL,
    `width` INT,
    `height` INT,
    `extension` VARCHAR(10),
    `folder` VARCHAR(255),
    `cloud_name` VARCHAR(100),
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `thinkflow-media`.`audios` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `file_url` VARCHAR(500) NOT NULL,
    `transcript_id` BIGINT DEFAULT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `idx_transcript_id` (`transcript_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Bảng trong thinkflow-gen
CREATE TABLE `thinkflow-gen`.`transcripts` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `content` TEXT NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `thinkflow-gen`.`texts` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `content` TEXT NOT NULL,
    `format` ENUM('plain', 'markdown', 'html') DEFAULT 'plain',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `thinkflow-gen`.`summaries` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `summary_text` TEXT NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `thinkflow-gen`.`mindmaps` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `mindmap_data` JSON NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Bảng trong thinkflow-notifications
CREATE TABLE `thinkflow-notifications`.`notifications` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `message` TEXT NOT NULL,
    `is_read` BOOLEAN DEFAULT FALSE,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `idx_user_id_is_read` (`user_id`, `is_read`),
    INDEX `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;