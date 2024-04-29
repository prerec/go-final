-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date VARCHAR(8),
    title TEXT,
    comment TEXT,
    repeat VARCHAR(128)
);

CREATE INDEX todo_date ON scheduler (date);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE scheduler;
-- +goose StatementEnd