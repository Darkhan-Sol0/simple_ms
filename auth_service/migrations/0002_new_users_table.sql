-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(100),
    login VARCHAR(100) NOT NULL UNIQUE,
    email VARCHAR(100),
    phone VARCHAR(100),
    password VARCHAR(100) NOT NULL,
    id_role INT NOT NULL DEFAULT 2 REFERENCES roles(id)
);
CREATE UNIQUE INDEX unique_email ON users(email) WHERE email IS NOT NULL;
CREATE UNIQUE INDEX unique_phone ON users(phone) WHERE phone IS NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
