-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE users (
    id bigserial primary key,
    name varchar(255) not null,
    email varchar(255) not null,
    email_verified_at timestamp(0) without time zone default null,
    password varchar(255) not null,
    created_at timestamp(0) without time zone default null,
    updated_at timestamp(0) without time zone default null
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
