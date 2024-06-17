-- +goose Up
-- +goose StatementBegin
ALTER TABLE companies ADD COLUMN deleted_at TIMESTAMP NULL DEFAULT NULL;
ALTER TABLE announcements ADD COLUMN active BOOLEAN DEFAULT TRUE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE companies DROP COLUMN deleted_at;
ALTER TABLE announcements DROP COLUMN active;
-- +goose StatementEnd
