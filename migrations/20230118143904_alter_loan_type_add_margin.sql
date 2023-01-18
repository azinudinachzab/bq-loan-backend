-- +goose Up
ALTER TABLE loan_types ADD margin decimal(20,3) NOT NULL DEFAULT 0 AFTER name;

-- +goose Down
ALTER TABLE loan_types DROP COLUMN margin;
