-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(100) NOT NULL UNIQUE,
    login VARCHAR(100) NOT NULL UNIQUE,
    email VARCHAR(100),
    phone VARCHAR(100),
    passwordhash VARCHAR(255) NOT NULL,
    id_role INT NOT NULL DEFAULT 2 REFERENCES roles(id)
);
CREATE UNIQUE INDEX unique_email ON users(email) WHERE email IS NOT NULL;
CREATE UNIQUE INDEX unique_phone ON users(phone) WHERE phone IS NOT NULL;

INSERT INTO users (uuid, login, email, phone, passwordhash, id_role) 
    VALUES ('c7eba5fd-f4d8-4e4c-8c38-b1470941e2cc', 'admin', 'test@test.com', '+77777777777', '$2a$10$8weqV5yZqdMTfYSmbeJuge1bJ1d66fixocYQfYlNOFubZgHofefx2', 1)
-- password 123qwe
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
