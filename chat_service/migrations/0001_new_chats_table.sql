-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE TYPE chat_type AS ENUM ('group', 'personal');

CREATE TABLE chats (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    author_uuid UUID
    chat_name VARCHAR(255)
    type chat_type DEFAULT 'personal'
    
);

CREATE INDEX chat_uuid_idx ON chat(uuid); 

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS chats;
DROP EXTENSION IF EXISTS pgcrypto;
-- +goose StatementEnd
