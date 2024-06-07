-- +goose Up
-- +goose StatementBegin
ALTER TABLE announcements ADD COLUMN content varchar(1000);
BEGIN;

SELECT * FROM announcements;

UPDATE announcements SET content = '<h1>This is my offer!<br></h1><p>I like when you buy! Please <strong>buy</strong>! <br>I need:</p><ul><li><p>money</p></li><li><p>money</p></li><li><p>money</p></li></ul><p></p>';

COMMIT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE announcements DROP COLUMN content;
-- +goose StatementEnd
