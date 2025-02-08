-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id         uuid primary key not null default gen_random_uuid(),
    name       text,
    email      text,
    password   text,
    created_at   timestamp with time zone,
    updated_at   timestamp with time zone,
    deleted_at   timestamp with time zone
);

CREATE UNIQUE INDEX idx_users_email on users using btree (email);
CREATE INDEX idx_users_deleted_at on users using btree (deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd