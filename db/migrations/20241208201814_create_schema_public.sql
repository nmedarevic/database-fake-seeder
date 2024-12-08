-- +goose Up
-- +goose StatementBegin
CREATE TABLE "public"."user" (
    "id" uuid NOT NULL,
    "email" varchar(255) NOT NULL,
    "password" varchar(255) NOT NULL,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "public"."user";
-- +goose StatementEnd
