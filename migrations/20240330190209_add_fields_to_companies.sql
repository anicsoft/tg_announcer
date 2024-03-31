-- +goose Up
alter table companies
    add column address text not null default '',
    add column latitude float8 default 0,
    add column longitude float8 default 0,
    add column who text not null default '';

ALTER TABLE companies
    DROP CONSTRAINT IF EXISTS companies_name_key;

-- +goose Down
alter table companies
    drop column address,
    drop column latitude,
    drop column longitude,
    drop column who;

alter table companies
    add constraint companies_name_key unique (name);
