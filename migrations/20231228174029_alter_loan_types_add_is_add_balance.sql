-- +goose Up
ALTER TABLE loan_types ADD is_add_balance int NOT NULL DEFAULT 0 COMMENT '1 no add balance\n2 add balance\n3 add vbalance';

-- +goose Down
ALTER TABLE loan_types DROP COLUMN is_add_balance;
