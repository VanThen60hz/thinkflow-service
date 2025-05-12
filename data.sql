-- Create databases
CREATE DATABASE IF NOT EXISTS `thinkflow-auth`;
CREATE DATABASE IF NOT EXISTS `thinkflow-users`;
CREATE DATABASE IF NOT EXISTS `thinkflow-notes`;
CREATE DATABASE IF NOT EXISTS `thinkflow-collaborations`;
CREATE DATABASE IF NOT EXISTS `thinkflow-media`;
CREATE DATABASE IF NOT EXISTS `thinkflow-gen`;
CREATE DATABASE IF NOT EXISTS `thinkflow-notifications`;

-- thinkflow-auth (auths)
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

-- thinkflow-users (users)
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

-- thinkflow-notes (notes)
CREATE TABLE `thinkflow-notes`.`notes` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `title` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `archived` BOOLEAN DEFAULT FALSE,
    `summary_id` BIGINT DEFAULT NULL,
    `mindmap_id` BIGINT DEFAULT NULL,
    UNIQUE INDEX `idx_summary_id` (`summary_id`),
    UNIQUE INDEX `idx_mindmap_id` (`mindmap_id`),
    INDEX `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- thinkflow-notes (texts)
CREATE TABLE `thinkflow-notes`.`texts` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `note_id` BIGINT NOT NULL,
    `text_content` JSON NOT NULL,
    `text_string` TEXT NOT NULL,
    `summary_id` BIGINT DEFAULT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE INDEX `idx_summary_id` (`summary_id`),
    UNIQUE INDEX `idx_note_id` (`note_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- thinkflow-collaborations (collaborations)
CREATE TABLE `thinkflow-collaborations`.`collaborations` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `note_id` BIGINT NOT NULL,
    `user_id` INT NOT NULL,
    `permission` ENUM('read', 'write') NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE INDEX `idx_note_id_user_id` (`note_id`, `user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- thinkflow-collaborations (note_share_links)
CREATE TABLE `thinkflow-collaborations`.`note_share_links` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `note_id` BIGINT NOT NULL,
    `permission` ENUM('read', 'write') NOT NULL,
    `token` VARCHAR(512) NOT NULL UNIQUE,
    `expires_at` DATETIME DEFAULT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `created_by` INT NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- thinkflow-media (images)
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

-- thinkflow-media (audios)
CREATE TABLE `thinkflow-media`.`audios` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `note_id` BIGINT NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `file_url` VARCHAR(500) NOT NULL,
    `format` VARCHAR(10) NOT NULL,
    `transcript_id` BIGINT DEFAULT NULL,
    `summary_id` BIGINT DEFAULT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE INDEX `idx_transcript_id` (`transcript_id`),
    UNIQUE INDEX `idx_summary_id` (`summary_id`),
    INDEX `idx_note_id` (`note_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `thinkflow-media`.`attachments` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `note_id` BIGINT NOT NULL,
    `file_url` VARCHAR(1000) NOT NULL,
    `file_name` VARCHAR(255),
    `extension` VARCHAR(10),
    `size_bytes` BIGINT,
    `cloud_name` VARCHAR(100),
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `idx_note_id` (`note_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- thinkflow-gen (summaries)
CREATE TABLE `thinkflow-gen`.`summaries` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `summary_text` TEXT NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- thinkflow-gen (mindmaps)
CREATE TABLE `thinkflow-gen`.`mindmaps` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `mindmap_data` JSON NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- thinkflow-gen (transcripts)
CREATE TABLE `thinkflow-gen`.`transcripts` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `content` TEXT NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- thinkflow-notifications (notifications)
CREATE TABLE `thinkflow-notifications`.`notifications` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `noti_type` ENUM(
        'NOTE_CREATED',
        'TRANSCRIPT_GENERATED',
        'SUMMARY_GENERATED',
        'MINDMAP_GENERATED',
        'AUDIO_PROCESSED',
        'TEXT_PROCESSED',
        'REMINDER',
        'COLLAB_INVITE',
        'COMMENT'
    ) NOT NULL,
    `noti_sender_id` BIGINT NOT NULL,
    `noti_received_id` BIGINT NOT NULL,
    `noti_content` TEXT NOT NULL,
    `noti_options` JSON DEFAULT (JSON_OBJECT()),
    `is_read` BOOLEAN DEFAULT FALSE,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX `idx_type_content` (`noti_type`, `noti_content`(255))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- thinkflow-notifications (fcm_tokens)
CREATE TABLE `thinkflow-notifications`.`fcm_tokens` (
    `id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `user_id` VARCHAR(255) NOT NULL,
    `token` VARCHAR(255) NOT NULL,
    `device_id` VARCHAR(255) NOT NULL,
    `platform` VARCHAR(50) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,
    UNIQUE INDEX `idx_token` (`token`),
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_device_id` (`device_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;