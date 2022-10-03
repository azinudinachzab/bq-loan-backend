-- +goose Up
CREATE TABLE IF NOT EXISTS `users`(
    `id` INT UNSIGNED AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `email` varchar(100) NOT NULL,
    `password` varchar(255) NOT NULL,
    `role` INT NOT NULL DEFAULT 1,
    `is_active` INT NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`),
    CONSTRAINT `unique_user_name` UNIQUE (`name`),
    CONSTRAINT `unique_user_email` UNIQUE (`email`)
);

-- +goose Down
DROP TABLE IF EXISTS `users`;
