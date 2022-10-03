-- +goose Up
CREATE TABLE IF NOT EXISTS `income`(
    `id` INT UNSIGNED AUTO_INCREMENT,
    `title` varchar(255) NOT NULL,
    `amount` decimal(20,3) NOT NULL DEFAULT 0,
    `datetime` TIMESTAMP NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`)
);

-- +goose Down
DROP TABLE IF EXISTS `income`;
