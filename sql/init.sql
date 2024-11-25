CREATE DATABASE IF NOT EXISTS mydb;

USE mydb;

CREATE TABLE IF NOT EXISTS `posts`
(
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `user_id` VARCHAR(36) NOT NULL,
    `text` TEXT NOT NULL,
    `image_path` VARCHAR(255),       -- 이미지 파일 경로
    `location` JSON,                 -- 위치 정보 (JSON으로 저장)
    `category` VARCHAR(50) DEFAULT NULL, -- AI가 분류한 카테고리
    `accuracy` FLOAT DEFAULT NULL,       -- AI 결과 신뢰도
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)Engine=InnoDB DEFAULT CHARSET=utf8mb4;

