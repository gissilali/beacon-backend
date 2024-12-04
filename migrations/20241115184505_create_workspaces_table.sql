-- +goose Up
-- +goose StatementBegin
CREATE TABLE workspaces (
    id         bigserial    not null primary key,
    name       varchar(255) not null,
    owner_id   int          not null,
    created_at timestamp(0) without time zone default null,
    updated_at timestamp(0) without time zone default null,
    foreign key (owner_id) references users (id) on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE workspaces;
-- +goose StatementEnd
