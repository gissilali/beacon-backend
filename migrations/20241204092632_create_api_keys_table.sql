-- +goose Up
-- +goose StatementBegin
CREATE TABLE api_keys
(
    id bigserial not null primary key,
    user_id bigint,
    key text not null,
    scopes jsonb null,
    created_at timestamp(0) without time zone default null,
    updated_at timestamp(0) without time zone default null,
    last_used_at timestamp(0) without time zone default null,
    revoked boolean default false
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE api_keys;
-- +goose StatementEnd
