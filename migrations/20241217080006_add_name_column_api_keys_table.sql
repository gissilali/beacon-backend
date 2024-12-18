-- +goose Up
-- +goose StatementBegin
ALTER TABLE api_keys ADD name varchar(255) not null default '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
