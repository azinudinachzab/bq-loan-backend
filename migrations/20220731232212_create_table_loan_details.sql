-- +goose Up
CREATE TABLE IF NOT EXISTS `loan_details`(
    `id` INT UNSIGNED AUTO_INCREMENT,
    `loan_general_id` INT UNSIGNED NOT NULL,
    `amount` decimal(20,3) NOT NULL DEFAULT 0,
    `datetime` TIMESTAMP NOT NULL,
--     0 belum dilunasi, 1 sudah dilunasi
    `status` INT NOT NULL DEFAULT 0,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`),
    FOREIGN KEY (`loan_general_id`)
    REFERENCES `loan_generals` (`id`)
    ON UPDATE RESTRICT ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS `loan_details`;
