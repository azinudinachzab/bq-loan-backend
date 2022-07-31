-- +goose Up
CREATE TABLE IF NOT EXISTS `loan_generals`(
    `id` INT UNSIGNED AUTO_INCREMENT,
    `user_id` INT UNSIGNED NOT NULL,
    `title` varchar(255) NOT NULL,
    `amount` decimal(20,3) NOT NULL DEFAULT 0,
    `datetime` TIMESTAMP NOT NULL,
    `tenor` INT NOT NULL,
--     0 diproses, 1 disetujui, 2 ditolak, 3 selesai
    `status` INT NOT NULL DEFAULT 0,
    `loan_type_id` INT UNSIGNED NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`),
    FOREIGN KEY (`user_id`)
        REFERENCES `users` (`id`)
        ON UPDATE RESTRICT ON DELETE CASCADE,
    FOREIGN KEY (`loan_type_id`)
        REFERENCES `loan_types` (`id`)
        ON UPDATE RESTRICT ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS `loan_generals`;
