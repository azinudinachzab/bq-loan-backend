-- +goose Up
CREATE TABLE IF NOT EXISTS `social_funds`(
    `id` INT UNSIGNED AUTO_INCREMENT,
    `title` varchar(255) NOT NULL,
    `user_id` INT UNSIGNED NOT NULL,
    `fund_type` INT NOT NULL COMMENT '0 income 1 expense',
    `amount` decimal(20,3) NOT NULL DEFAULT 0,
    `status` INT NOT NULL DEFAULT 1,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`),
    FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`)
        ON UPDATE RESTRICT ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS `social_funds`;
