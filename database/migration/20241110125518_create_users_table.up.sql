CREATE TABLE IF NOT EXISTS `users` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `first_name` VARCHAR(255) NOT NULL,
    `last_name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `phone` VARCHAR(255),
    `gender` TINYINT DEFAULT 1,
    `status` INT DEFAULT 1,
    `profile_picture` VARCHAR(255),
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NULL,
    `deleted_at` DATETIME NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `UQ_user_email` UNIQUE (`email`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
