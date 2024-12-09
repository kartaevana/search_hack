-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table if not exists users (
    id serial not null primary key,
    email varchar unique,
    name varchar,
    sur_name varchar,
    hashed_password varchar,
    tg varchar
);

create table if not exists form (
    id serial not null primary key,
    id_user bigint,
    id_hack bigint,
    about varchar,
    photo varchar,
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
    id_hack bigint,
    id_kap bigint
);

create table if not exists team_user (
    team_id bigint,
    user_id bigint
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists users;
-- +goose StatementEnd
