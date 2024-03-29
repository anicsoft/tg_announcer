-- +goose Up
create table companies (
    id text not null unique,
    name text not null unique,
    description text not null,
    create_at timestamp not null default now()
);

-- +goose Down
drop table companies;
