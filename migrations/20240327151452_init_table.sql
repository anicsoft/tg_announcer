-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE announcements (
    announcement_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    company_id UUID,
    title TEXT,
    content VARCHAR(1000),
    promo_code VARCHAR(50),
    picture_url TEXT,
    start_date_time TIMESTAMPTZ,
    end_date_time TIMESTAMPTZ,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE companies (
    company_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255),
    description TEXT,
    address TEXT,
    logo_url TEXT,
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,
    updated_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255) DEFAULT '',
    username VARCHAR(255),
    language_code VARCHAR(10) DEFAULT 'en',
    user_type VARCHAR(20) DEFAULT 'user',
    company_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE offercategories (
    offer_category_id serial PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE businesscategories (
    category_id serial PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE announcementoffers (
    announcement_id UUID,
    offer_category_id SERIAL,
    PRIMARY KEY (announcement_id, offer_category_id)
);

CREATE TABLE companycategories (
    company_id UUID,
    category_id SERIAL,
    PRIMARY KEY (company_id, category_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS AnnouncementOffers;
DROP TABLE IF EXISTS OfferCategories;
DROP TABLE IF EXISTS Announcements;
DROP TABLE IF EXISTS CompanyCategories;
DROP TABLE IF EXISTS BusinessCategories;
DROP TABLE IF EXISTS Users;
DROP TABLE IF EXISTS Companies;
-- +goose StatementEnd
