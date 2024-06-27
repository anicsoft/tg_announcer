-- +goose Up
-- +goose StatementBegin
ALTER TABLE companies ADD COLUMN tel_number TEXT DEFAULT NULL;
ALTER TABLE companies ADD COLUMN email TEXT DEFAULT NULL;
ALTER TABLE companies ADD COLUMN website TEXT DEFAULT NULL;
ALTER TABLE companies ADD COLUMN facebook TEXT DEFAULT NULL;
ALTER TABLE companies ADD COLUMN instagram TEXT DEFAULT NULL;
ALTER TABLE companies ADD COLUMN telegram TEXT DEFAULT NULL;

CREATE TABLE business_hours (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    company_id UUID,
    day_of_week INT NOT NULL CHECK (day_of_week >= 0 AND day_of_week <= 6),
    open_time TIME NOT NULL,
    close_time TIME NOT NULL,
    FOREIGN KEY (company_id) REFERENCES companies(company_id) ON DELETE CASCADE
);

DO
$$
    DECLARE
        company RECORD;
    BEGIN
        FOR company IN SELECT company_id FROM companies LOOP
                INSERT INTO business_hours (company_id, day_of_week, open_time, close_time) VALUES
                    (company.company_id, 0, '09:00:00', '17:00:00'), -- Sunday
                    (company.company_id, 1, '09:00:00', '17:00:00'), -- Monday
                    (company.company_id, 2, '09:00:00', '17:00:00'), -- Tuesday
                    (company.company_id, 3, '09:00:00', '17:00:00'), -- Wednesday
                    (company.company_id, 4, '09:00:00', '17:00:00'), -- Thursday
                    (company.company_id, 5, '09:00:00', '17:00:00'), -- Friday
                    (company.company_id, 6, '10:00:00', '14:00:00'); -- Saturday
            END LOOP;
    END
$$;

UPDATE companies SET
    tel_number = '123-456-7890',
    email = 'info@dummycompany.com',
    website = 'https://www.dummycompany.com',
    facebook = 'https://www.facebook.com/dummycompany',
    instagram = 'https://www.instagram.com/dummycompany',
    telegram = '@dummycompany';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE companies DROP COLUMN tel_number;
ALTER TABLE companies DROP COLUMN email;
ALTER TABLE companies DROP COLUMN website;
ALTER TABLE companies DROP COLUMN facebook;
ALTER TABLE companies DROP COLUMN instagram;
ALTER TABLE companies DROP COLUMN telegram;

DROP TABLE business_hours;
-- +goose StatementEnd
