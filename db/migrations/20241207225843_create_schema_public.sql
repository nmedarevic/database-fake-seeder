-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS "public"; 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA IF EXISTS "public";
-- +goose StatementEnd
