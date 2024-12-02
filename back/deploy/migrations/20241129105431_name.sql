-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table if not exists users (
    id serial not null primary key,
    email varchar unique,
    name varchar,
    sur_name varchar,
    hashed_password varchar,
    tg varchar,
    about varchar,
    sphere varchar
);

create table if not exists hack (
    id serial not null primary key,
    name varchar,
    theame varchar,
    about varchar
);

create table if not exists team (
    id serial not null primary key,
    name varchar,
    id_hack varchar,
    id_kap varchar
);

create table if not exists team_user (
    team_id serial not null primary key,
    user_id varchar
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists users;
-- +goose StatementEnd
