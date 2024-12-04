-- +goose Up
-- +goose StatementBegin
CREATE TABLE auth_tokens
(
    id            bigserial not null primary key,
    user_id       int references users (id) on delete cascade,
    refresh_token text      not null unique, -- Secure, opaque token
    created_at    timestamp default now(),
    expires_at    timestamp not null,
    revoked       boolean   default false    -- Tracks token revocation
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE auth_tokens
-- +goose StatementEnd
