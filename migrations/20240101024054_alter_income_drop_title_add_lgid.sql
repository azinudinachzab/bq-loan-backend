-- +goose Up
ALTER TABLE income ADD loan_general_id int unsigned NOT NULL;
ALTER TABLE income DROP COLUMN title;

-- +goose Down
ALTER TABLE income ADD title varchar(255) NOT NULL;
ALTER TABLE income DROP COLUMN loan_general_id;
