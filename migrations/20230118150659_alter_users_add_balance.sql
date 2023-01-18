-- +goose Up
ALTER TABLE users ADD balance decimal(20,3) NOT NULL DEFAULT 0;

-- +goose Down
ALTER TABLE users DROP COLUMN balance;
