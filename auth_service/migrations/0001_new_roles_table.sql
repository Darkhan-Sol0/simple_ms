-- +goose Up
-- +goose StatementBegin
CREATE TABLE roles (
  id SERIAL PRIMARY KEY,
  role VARCHAR(100) NOT NULL UNIQUE
);

INSERT INTO roles (role) VALUES ('admin'),('user');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE roles;
-- +goose StatementEnd
