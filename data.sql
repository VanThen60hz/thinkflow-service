-- Cơ sở dữ liệu cho xác thực và thông tin người dùng
CREATE DATABASE IF NOT EXISTS `thinkflow-auth`;

-- Cơ sở dữ liệu cho tài khoản người dùng
CREATE DATABASE IF NOT EXISTS `thinkflow-users`;

-- Cơ sở dữ liệu cho ghi chú, cộng tác và tóm tắt
CREATE DATABASE IF NOT EXISTS `thinkflow-notes`;

-- Cơ sở dữ liệu cho media (ảnh và âm thanh)
CREATE DATABASE IF NOT EXISTS `thinkflow-media`;

-- Bảng images (Lưu trữ tất cả loại ảnh: avatar, logo, background)
CREATE TABLE IF NOT EXISTS `thinkflow-media`.`images` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `url` LONGTEXT NOT NULL,
    `width` BIGINT,
    `height` BIGINT,
    `extension` LONGTEXT,
    `folder` LONGTEXT,
    `cloud_name` LONGTEXT,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Bảng users (Quản lý tài khoản người dùng)
CREATE TABLE IF NOT EXISTS `thinkflow-users`.`users` (
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
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `idx_avatar_id` (`avatar_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Bảng auths (Quản lý đăng nhập)
CREATE TABLE IF NOT EXISTS `thinkflow-auth`.`auths` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT NOT NULL,
    `auth_type` ENUM('email_password', 'gmail', 'facebook') DEFAULT 'email_password',
    `email` VARCHAR(255) NOT NULL UNIQUE,
    `salt` VARCHAR(40) DEFAULT NULL,
    `password` VARCHAR(100) DEFAULT NULL,
    `facebook_id` VARCHAR(35) DEFAULT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_facebook_id` (`facebook_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Bảng notifications (Thông báo hệ thống)
CREATE TABLE IF NOT EXISTS `thinkflow-notes`.`notifications` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT NOT NULL,
    `message` TEXT NOT NULL,
    `status` ENUM('unread', 'read') DEFAULT 'unread',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Bảng notes (Quản lý ghi chú)
CREATE TABLE IF NOT EXISTS `thinkflow-notes`.`notes` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT NOT NULL,
    `title` VARCHAR(255),
    `content` TEXT,
    `image_id` BIGINT DEFAULT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_image_id` (`image_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Bảng collaborations (Chia sẻ ghi chú & cộng tác)
CREATE TABLE IF NOT EXISTS `thinkflow-notes`.`collaborations` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `note_id` INT NOT NULL,
    `collaborator_id` INT NOT NULL,
    `permission` ENUM('view', 'edit') NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `idx_note_id` (`note_id`),
    INDEX `idx_collaborator_id` (`collaborator_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Bảng audio_files (Lưu trữ file âm thanh)
CREATE TABLE IF NOT EXISTS `thinkflow-media`.`audio_files` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `url` LONGTEXT NOT NULL,
    `format` VARCHAR(50),
    `duration` BIGINT,
    `uploaded_at` DATETIME DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Bảng voice_notes (Lưu ghi âm & chuyển giọng nói thành văn bản)
CREATE TABLE IF NOT EXISTS `thinkflow-notes`.`voice_notes` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT NOT NULL,
    `note_id` INT DEFAULT NULL,
    `audio_id` BIGINT NOT NULL,
    `transcript` TEXT,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_note_id` (`note_id`),
    INDEX `idx_audio_id` (`audio_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Bảng summaries (Lưu kết quả tóm tắt bằng AI)
CREATE TABLE IF NOT EXISTS `thinkflow-notes`.`summaries` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `note_id` INT NOT NULL,
    `summary` TEXT NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `idx_note_id` (`note_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Bảng mindmaps (Lưu sơ đồ tư duy của ghi chú)
CREATE TABLE IF NOT EXISTS `thinkflow-notes`.`mindmaps` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `note_id` INT NOT NULL,
    `structure` JSON NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `idx_note_id` (`note_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;