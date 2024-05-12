-- +goose Up
CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255) default '',
    username VARCHAR(255),
    latitude DOUBLE PRECISION default null,
    longitude DOUBLE PRECISION default null,
    language_code VARCHAR(10) default 'en',
    user_type VARCHAR(20) default 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS Users;
