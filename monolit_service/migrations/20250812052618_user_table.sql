-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  uuid SERIAL PRIMARY KEY,
  login TEXT UNIQUE NOT NULL,
  email TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
