-- +goose Up
-- +goose StatementBegin
CREATE TABLE Image (
    id serial PRIMARY KEY UNIQUE NOT NULL,
    url text NOT NULL,
    announcement_id integer,
    company_id integer,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    FOREIGN KEY (announcement_id) REFERENCES Announcements(announcement_id) ON DELETE CASCADE,
    FOREIGN KEY (company_id) REFERENCES Companies(company_id)  ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE image;
-- +goose StatementEnd