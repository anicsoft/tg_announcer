-- +goose Up
-- +goose StatementBegin
CREATE TABLE favorites (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id INT NOT NULL,
    company_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (company_id) REFERENCES companies(company_id) ON DELETE CASCADE,
    UNIQUE (user_id, company_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE favorites;
-- +goose StatementEnd
