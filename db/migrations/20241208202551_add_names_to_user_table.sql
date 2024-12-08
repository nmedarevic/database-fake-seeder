-- +goose Up
-- +goose StatementBegin
ALTER TABLE IF EXISTS "public"."user" 
ADD COLUMN IF NOT EXISTS username varchar(255) NOT NULL,
ADD COLUMN IF NOT EXISTS full_name varchar(255) NOT NULL; 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS "public"."user" 
DROP COLUMN username,
DROP COLUMN full_name;
-- +goose StatementEnd
