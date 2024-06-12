-- +goose Up
-- +goose StatementBegin
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

INSERT INTO Companies (company_id, name, description, address, latitude, longitude) VALUES
    ('0e3df004-ca0c-45a3-aeee-fa21c4aa3e4d', 'Tartu Bakery', 'Traditional Estonian bakery', 'Tartu, Estonia', 58.3780, 26.7296),
    ('f8afcc36-45a1-4cae-bf60-308f3c405896', 'Tallinn Tech Store', 'Electronics store in Tallinn', 'Tallinn, Estonia', 59.4370, 24.7536),
    ('5d164261-5822-46cd-8976-fceb1c47387a', 'Parnu Spa Resort', 'Luxury spa resort in Parnu', 'Parnu, Estonia', 58.3854, 24.4971),
    ('edf51f88-236c-4708-8421-b2843d2834a3', 'Narva Nightclub', 'Popular nightclub in Narva', 'Narva, Estonia', 59.3772, 28.1900),
    ('81d5b127-8d68-4a81-afd3-a2e280ef0d8a', 'Haapsalu Hotel', 'Cozy hotel in Haapsalu', 'Haapsalu, Estonia', 58.9431, 23.5410);

INSERT INTO announcements (company_id, title, start_date_time, end_date_time, promo_code) VALUES
    ('0e3df004-ca0c-45a3-aeee-fa21c4aa3e4d', 'Grand Opening Sale', '2024-05-01 00:00:00+00', '2024-05-15 23:59:59+00', 'GRANDSALE'),
    ('0e3df004-ca0c-45a3-aeee-fa21c4aa3e4d', 'Specialty Bread Tasting Event', '2024-06-01 00:00:00+00', '2024-06-05 18:00:00+00', NULL),
    ('f8afcc36-45a1-4cae-bf60-308f3c405896', 'Tech Expo 2024', '2024-07-10 00:00:00+00', '2024-07-15 23:59:59+00', 'TECHEXPO'),
    ('f8afcc36-45a1-4cae-bf60-308f3c405896', 'Back to School Sale', '2024-08-20 00:00:00+00', '2024-09-01 23:59:59+00', 'SCHOOLSALE'),
    ('5d164261-5822-46cd-8976-fceb1c47387a', 'Spa Retreat Weekend', '2024-06-15 00:00:00+00', '2024-06-17 23:59:59+00', 'SPAWEEKEND'),
    ('5d164261-5822-46cd-8976-fceb1c47387a', 'Wellness Workshop Series', '2024-07-05 00:00:00+00', '2024-07-25 23:59:59+00', NULL),
    ('edf51f88-236c-4708-8421-b2843d2834a3', 'Summer Nights Dance Party', '2024-07-20 00:00:00+00', '2024-07-21 02:00:00+00', 'SUMMERPARTY'),
    ('edf51f88-236c-4708-8421-b2843d2834a3', 'Live Music Fridays', '2024-08-06 00:00:00+00', '2024-08-30 01:00:00+00', NULL),
    ('81d5b127-8d68-4a81-afd3-a2e280ef0d8a', 'Weekend Getaway Special', '2024-09-01 14:00:00+00', '2024-09-03 12:00:00+00', 'GETAWAY2024'),
    ('81d5b127-8d68-4a81-afd3-a2e280ef0d8a', 'Autumn Relaxation Package', '2024-10-01 14:00:00+00', '2024-10-10 12:00:00+00', 'AUTUMNRELAX'),
    ('81d5b127-8d68-4a81-afd3-a2e280ef0d8a', 'Winter Wonderland Stay', '2024-12-01 14:00:00+00', '2024-12-31 12:00:00+00', 'WINTERSTAY'),
    ('81d5b127-8d68-4a81-afd3-a2e280ef0d8a', 'Spring Blossom Retreat', '2024-04-01 14:00:00+00', '2024-04-10 12:00:00+00', 'SPRINGRETREAT');

-- Insert company categories
INSERT INTO companycategories (company_id, category_id) VALUES
    ('0e3df004-ca0c-45a3-aeee-fa21c4aa3e4d', (SELECT category_id FROM businesscategories WHERE name = 'Food & Drinks')),
    ('f8afcc36-45a1-4cae-bf60-308f3c405896', (SELECT category_id FROM businesscategories WHERE name = 'Shops')),
    ('5d164261-5822-46cd-8976-fceb1c47387a', (SELECT category_id FROM businesscategories WHERE name = 'Beauty')),
    ('edf51f88-236c-4708-8421-b2843d2834a3', (SELECT category_id FROM businesscategories WHERE name = 'Entertainment')),
    ('81d5b127-8d68-4a81-afd3-a2e280ef0d8a', (SELECT category_id FROM businesscategories WHERE name = 'Entertainment'));

-- Insert announcement offers
INSERT INTO announcementoffers (announcement_id, offer_category_id) VALUES
    ((SELECT announcement_id FROM announcements WHERE title = 'Grand Opening Sale'), (SELECT offer_category_id FROM offercategories WHERE name = 'Sale')),
    ((SELECT announcement_id FROM announcements WHERE title = 'Grand Opening Sale'), (SELECT offer_category_id FROM offercategories WHERE name = 'Special Offer')),
    ((SELECT announcement_id FROM announcements WHERE title = 'Specialty Bread Tasting Event'), (SELECT offer_category_id FROM offercategories WHERE name = 'Special Offer')),
    ((SELECT announcement_id FROM announcements WHERE title = 'Tech Expo 2024'), (SELECT offer_category_id FROM offercategories WHERE name = 'Promo Code')),
    ((SELECT announcement_id FROM announcements WHERE title = 'Back to School Sale'), (SELECT offer_category_id FROM offercategories WHERE name = 'Sale')),
    ((SELECT announcement_id FROM announcements WHERE title = 'Back to School Sale'), (SELECT offer_category_id FROM offercategories WHERE name = 'Special Offer')),
    ((SELECT announcement_id FROM announcements WHERE title = 'Spa Retreat Weekend'), (SELECT offer_category_id FROM offercategories WHERE name = 'Special Offer')),
    ((SELECT announcement_id FROM announcements WHERE title = 'Wellness Workshop Series'), (SELECT offer_category_id FROM offercategories WHERE name = 'Lunch')),
    ((SELECT announcement_id FROM announcements WHERE title = 'Wellness Workshop Series'), (SELECT offer_category_id FROM offercategories WHERE name = 'Happy Hour')),
    ((SELECT announcement_id FROM announcements WHERE title = 'Summer Nights Dance Party'), (SELECT offer_category_id FROM offercategories WHERE name = 'Happy Hour')),
    ((SELECT announcement_id FROM announcements WHERE title = 'Live Music Fridays'), (SELECT offer_category_id FROM offercategories WHERE name = 'Discount')),
    ((SELECT announcement_id FROM announcements WHERE title = 'Weekend Getaway Special'), (SELECT offer_category_id FROM offercategories WHERE name = 'Special Offer')),
    ((SELECT announcement_id FROM announcements WHERE title = 'Autumn Relaxation Package'), (SELECT offer_category_id FROM offercategories WHERE name = 'Special Offer')),
    ((SELECT announcement_id FROM announcements WHERE title = 'Winter Wonderland Stay'), (SELECT offer_category_id FROM offercategories WHERE name = 'Special Offer')),
    ((SELECT announcement_id FROM announcements WHERE title = 'Spring Blossom Retreat'), (SELECT offer_category_id FROM offercategories WHERE name = 'Special Offer'));


UPDATE announcements SET content = '<h1>This is my offer!<br></h1><p>I like when you buy! Please <strong>buy</strong>! <br>I need:</p><ul><li><p>money</p></li><li><p>money</p></li><li><p>money</p></li></ul><p></p>';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM BusinessCategories;
DELETE FROM OfferCategories;
DELETE FROM Companies;
DELETE FROM Announcements;
DELETE FROM CompanyCategories;
DELETE FROM AnnouncementOffers;
-- +goose StatementEnd
