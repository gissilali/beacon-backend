-- +goose Up
-- +goose StatementBegin
CREATE TABLE servers (
    id bigserial not null primary key,
    name varchar(255) not null,
    workspace_id int not null,
    created_at timestamp(0) without time zone default null,
    updated_at timestamp(0) without time zone default null,
    foreign key (workspace_id) references workspaces(id) on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE servers;
-- +goose StatementEnd
