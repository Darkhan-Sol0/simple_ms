-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  uuid UUID PRIMARY KEY,
  name VARCHAR(50),
  description TEXT,
  born_day TIMESTAMPTZ,
  city   VARCHAR(50),
  links  JSONB NOT NULL DEFAULT '{}'::jsonb
);

CREATE INDEX idx_users_uuid ON users(uuid);
CREATE INDEX idx_users_name ON users(name);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
