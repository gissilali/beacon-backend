-- +goose Up
-- +goose StatementBegin
CREATE TABLE workspace_members (
    id           bigserial not null primary key,
    user_id      int       not null,
    workspace_id int       not null,
    joined_at    timestamp default current_timestamp,
    foreign key (user_id) references users (id) on delete cascade,
    foreign key (workspace_id) references workspaces (id) on delete cascade,
    unique (user_id, workspace_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE workspace_members;
-- +goose StatementEnd
