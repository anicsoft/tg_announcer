-- +goose Up
-- +goose StatementBegin
ALTER TABLE companies ADD COLUMN tel_number TEXT;
ALTER TABLE companies ADD COLUMN email TEXT;
ALTER TABLE companies ADD COLUMN website TEXT;
ALTER TABLE companies ADD COLUMN facebook TEXT;
ALTER TABLE companies ADD COLUMN instagram TEXT;
ALTER TABLE companies ADD COLUMN telegram TEXT;

CREATE TABLE business_hours (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    company_id UUID,
    day_of_week INT NOT NULL CHECK (day_of_week >= 0 AND day_of_week <= 6),
    open_time TIME NOT NULL,
    close_time TIME NOT NULL,
    FOREIGN KEY (company_id) REFERENCES companies(company_id) ON DELETE CASCADE
);
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
