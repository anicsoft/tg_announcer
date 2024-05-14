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
    company_id INT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (company_id) REFERENCES Companies(company_id)
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

INSERT INTO Companies (name, description, address, latitude, longitude) VALUES
    ('Tartu Bakery', 'Traditional Estonian bakery', 'Tartu, Estonia', 58.3780, 26.7296),
    ('Tallinn Tech Store', 'Electronics store in Tallinn', 'Tallinn, Estonia', 59.4370, 24.7536),
    ('Parnu Spa Resort', 'Luxury spa resort in Parnu', 'Parnu, Estonia', 58.3854, 24.4971),
    ('Narva Nightclub', 'Popular nightclub in Narva', 'Narva, Estonia', 59.3772, 28.1900),
    ('Haapsalu Hotel', 'Cozy hotel in Haapsalu', 'Haapsalu, Estonia', 58.9431, 23.5410);

INSERT INTO CompanyCategories (company_id, category_id) VALUES
    (1, 1),
    (2, 2),
    (3, 3),
    (4, 4),
    (5, 1);

-- Add dummy announcements for each company
INSERT INTO Announcements (company_id, title, start_date, end_date, start_time, end_time, promo_code) VALUES
    (1, 'Grand Opening Sale', '2024-05-01', '2024-05-15', '12:00:00', '22:00:00', 'GRANDSALE'),
    (1, 'Specialty Bread Tasting Event', '2024-06-01', '2024-06-05', '15:00:00', '18:00:00', NULL),
    (2, 'Tech Expo 2024', '2024-07-10', '2024-07-15', '09:00:00', '17:00:00', 'TECHEXPO'),
    (2, 'Back to School Sale', '2024-08-20', '2024-09-01', '10:00:00', '20:00:00', 'SCHOOLSALE'),
    (3, 'Spa Retreat Weekend', '2024-06-15', '2024-06-17', '14:00:00', '22:00:00', 'SPAWEEKEND'),
    (3, 'Wellness Workshop Series', '2024-07-05', '2024-07-25', '13:00:00', '18:00:00', NULL),
    (4, 'Summer Nights Dance Party', '2024-07-20', '2024-07-21', '21:00:00', '02:00:00', 'SUMMERPARTY'),
    (4, 'Live Music Fridays', '2024-08-06', '2024-08-30', '20:00:00', '01:00:00', NULL);

-- Add categories for announcements
INSERT INTO AnnouncementOffers (announcement_id, offer_category_id) VALUES
    (1, 3),  -- Grand Opening Sale: Sale
    (1, 6),  -- Grand Opening Sale: Special Offer
    (2, 6),  -- Specialty Bread Tasting Event: Special Offer
    (3, 2),  -- Tech Expo 2024: Promo Code
    (4, 3),  -- Back to School Sale: Sale
    (4, 6),  -- Back to School Sale: Special Offer
    (5, 6),  -- Spa Retreat Weekend: Special Offer
    (6, 1),  -- Wellness Workshop Series: Lunch
    (6, 4),  -- Wellness Workshop Series: Happy Hour
    (7, 4),  -- Summer Nights Dance Party: Happy Hour
    (8, 5);  -- Live Music Fridays: Discount

-- +goose Down
DROP TABLE IF EXISTS AnnouncementOffers;
DROP TABLE IF EXISTS OfferCategories;
DROP TABLE IF EXISTS Announcements;
DROP TABLE IF EXISTS CompanyCategories;
DROP TABLE IF EXISTS BusinessCategories;
DROP TABLE IF EXISTS Users;
DROP TABLE IF EXISTS Companies;
