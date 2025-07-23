-- +goose Up
-- +goose StatementBegin
CREATE TABLE messenge (
    id PRIMARY KEY,
    author_uuid UUID NOT NULL,
    chat_uuid UUID NOT NULL,
    messenge TEXT,
    date TIMESTAMPTZ
);

CREATE INDEX chat_uuid_idx ON chat(uuid); 

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS chats;
DROP EXTENSION IF EXISTS pgcrypto;
-- +goose StatementEnd
