-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS blogs (
    id INTEGER PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
Drop table blogs;
-- +goose StatementEnd
