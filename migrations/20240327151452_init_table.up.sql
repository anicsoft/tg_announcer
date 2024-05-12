-- +goose Up
CREATE TABLE Companies (
    company_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description TEXT,
    address TEXT,
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION
);

CREATE TABLE BusinessCategories (
    category_id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE CompanyCategories (
    company_id INT REFERENCES Companies(company_id) ON DELETE CASCADE ,
    category_id INT REFERENCES BusinessCategories(category_id),
    PRIMARY KEY (company_id, category_id)
);

CREATE TABLE Announcements (
    announcement_id SERIAL PRIMARY KEY,
    company_id INT REFERENCES Companies(company_id),
    title TEXT,
    start_date DATE,
    end_date DATE,
    start_time TIME,
    end_time TIME,
    promo_code VARCHAR(50) DEFAULT NULL,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE OfferCategories (
    offer_category_id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE AnnouncementOffers (
    announcement_id INT REFERENCES Announcements(announcement_id) ON DELETE CASCADE ,
    offer_category_id INT REFERENCES OfferCategories(offer_category_id),
    PRIMARY KEY (announcement_id, offer_category_id)
);

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

INSERT INTO BusinessCategories (name) VALUES
    ('Food & Drinks'),
    ('Shops'),
    ('Beauty'),
    ('Entertainment');

INSERT INTO OfferCategories (name) VALUES
    ('Lunch'),
    ('Promo Code'),
    ('Sale'),
    ('Happy Hour'),
    ('Discount'),
    ('Special Offer');

-- +goose Down
DROP TABLE IF EXISTS AnnouncementOffers;
DROP TABLE IF EXISTS OfferCategories;
DROP TABLE IF EXISTS Announcements;
DROP TABLE IF EXISTS CompanyCategories;
DROP TABLE IF EXISTS BusinessCategories;
DROP TABLE IF EXISTS Companies;
DROP TABLE IF EXISTS Users;