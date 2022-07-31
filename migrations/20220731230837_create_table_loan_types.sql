-- +goose Up
CREATE TABLE IF NOT EXISTS `loan_types`(
    `id` INT UNSIGNED AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`)
);

-- +goose Down
DROP TABLE IF EXISTS `loan_types`;
