-- +goose Up
alter table companies
    add column address text not null default '',
    add column latitude text not null default '',
    add column longitude text not null default '',
    add column who text not null default '',
    drop constraint if exists uq_companies_name;

-- +goose Down
alter table companies
drop column address,
    drop column latitude,
    drop column longitude,
    drop column who,
    add constraint uq_companies_name unique (name);