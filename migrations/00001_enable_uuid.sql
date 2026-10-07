-- +goose Up
SELECT 'up SQL query';
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- +goose Down
SELECT 'down SQL query';
DROP EXTENSION IF EXISTS pgcrypto;