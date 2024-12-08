-- +goose Up
-- +goose StatementBegin
alter role postgres set search_path = "public";
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
