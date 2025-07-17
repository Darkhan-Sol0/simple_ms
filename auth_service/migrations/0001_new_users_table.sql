-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE TYPE user_role AS ENUM ('admin', 'user');

CREATE TABLE users (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    login VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(255) UNIQUE CHECK (email ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'),
    phone VARCHAR(20) UNIQUE CHECK (phone ~* '^\+[0-9]{10,15}$'),
    password_hash TEXT NOT NULL,
    user_role user_role NOT NULL DEFAULT 'user',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

COMMENT ON COLUMN users.email IS 'Check format: user@domain.tld';
COMMENT ON COLUMN users.phone IS 'Check format: +7xxxxxxxxxx';

CREATE UNIQUE INDEX unique_email ON users(email) WHERE email IS NOT NULL;
CREATE UNIQUE INDEX unique_phone ON users(phone) WHERE phone IS NOT NULL;
CREATE INDEX users_login_idx ON users(login); 
CREATE INDEX idx_users_uuid ON users(uuid);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER users_update_timestamp
BEFORE UPDATE ON users
FOR EACH ROW EXECUTE FUNCTION update_timestamp();

INSERT INTO users (uuid, login, email, phone, password_hash, user_role) 
    VALUES (
        'c7eba5fd-f4d8-4e4c-8c38-b1470941e2cc'::UUID, 
        'admin', 
        'admin@example.com', 
        '+71234567899', 
        '$2a$10$8weqV5yZqdMTfYSmbeJuge1bJ1d66fixocYQfYlNOFubZgHofefx2', -- 123qwe
        'admin'
        );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS users_update_timestamp ON users;
DROP FUNCTION IF EXISTS update_timestamp;
DROP TABLE IF EXISTS users;
DROP TYPE IF EXISTS user_role;
DROP EXTENSION IF EXISTS pgcrypto;
-- +goose StatementEnd
