-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';


create table if not exists approve (
    id serial not null primary key,
    team_id bigint,
    form_id bigint
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
