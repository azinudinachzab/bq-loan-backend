-- +goose Up
ALTER TABLE users ADD is_leader INT NOT NULL DEFAULT 0;

-- +goose Down
ALTER TABLE users DROP COLUMN is_leader;
