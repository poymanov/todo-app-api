-- +goose Up
-- +goose StatementBegin
CREATE TABLE tasks
(
    id           uuid primary key not null default gen_random_uuid(),
    user_id      uuid,
    description  text,
    created_at   timestamp with time zone,
    updated_at   timestamp with time zone,
    deleted_at   timestamp with time zone,
    is_completed boolean                   default false,
    foreign key (user_id) references public.users (id)
        match simple on update cascade on delete set null
);
CREATE INDEX idx_tasks_deleted_at ON tasks USING btree (deleted_at);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;
-- +goose StatementEnd