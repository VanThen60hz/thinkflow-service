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
    `auth_type` ENUM('email_password', 'gmail', 'facebook') DEFAULT 'email_password',
    `email` VARCHAR(255) NOT NULL UNIQUE,
    `salt` VARCHAR(40) DEFAULT NULL,
    `password` VARCHAR(100) DEFAULT NULL,
    `facebook_id` VARCHAR(35) DEFAULT NULL,
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
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT NOT NULL,
    `title` VARCHAR(255),
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `thinkflow-notes`.`blocks` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `note_id` BIGINT NOT NULL,
    `type` ENUM('text', 'image', 'audio'),
    `content` TEXT,
    `position` INT NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX `idx_note_id` (`note_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `thinkflow-notes`.`collaborations` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `note_id` BIGINT NOT NULL,
    `user_id` INT NOT NULL,
    `permission` ENUM('read', 'write') NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX `idx_note_id` (`note_id`),
    INDEX `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Bảng trong thinkflow-media
CREATE TABLE `thinkflow-media`.`images` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `url` VARCHAR(2000) NOT NULL, -- Tối ưu từ LONGTEXT
    `width` BIGINT,
    `height` BIGINT,
    `extension` VARCHAR(10), -- Ví dụ: 'jpg', 'png', không cần LONGTEXT
    `folder` VARCHAR(255), -- Đường dẫn thư mục thường không quá dài
    `cloud_name` VARCHAR(100), -- Tên cloud (như 'cloudinary') không cần LONGTEXT
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `thinkflow-media`.`audio_files` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `block_id` BIGINT NOT NULL UNIQUE,
    `file_url` VARCHAR(500) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX `idx_block_id` (`block_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `thinkflow-media`.`transcripts` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `audio_id` BIGINT NOT NULL UNIQUE,
    `content` TEXT NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX `idx_audio_id` (`audio_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Bảng trong thinkflow-gen
CREATE TABLE `thinkflow-gen`.`summaries` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `block_id` BIGINT NOT NULL UNIQUE,
    `source_blocks` JSON NOT NULL,
    `summary_text` TEXT NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX `idx_block_id` (`block_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `thinkflow-gen`.`mindmaps` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `block_id` BIGINT NOT NULL UNIQUE,
    `source_blocks` JSON NOT NULL,
    `mindmap_data` JSON NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX `idx_block_id` (`block_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Bảng trong thinkflow-notifications
CREATE TABLE `thinkflow-notifications`.`notifications` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT NOT NULL,
    `message` TEXT NOT NULL,
    `is_read` BOOLEAN DEFAULT FALSE,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;