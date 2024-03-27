CREATE TABLE companies (
    id text not null unique,
    name text not null unique,
    description text not null,
    create_at datetime default CURRENT_TIMESTAMP
);