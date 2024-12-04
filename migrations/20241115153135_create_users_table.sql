-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "citext";
CREATE TABLE users (
    id                bigserial primary key,
    name              varchar(255) not null,
    email citext unique not null,
    email_verified_at timestamp(0) without time zone default null,
    password          bytea not null,
    created_at        timestamp(0) with time zone not null default now(),
    updated_at        timestamp(0) with time zone default null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
