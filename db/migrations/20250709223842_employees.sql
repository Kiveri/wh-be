-- +goose Up
-- +goose StatementBegin
create table employees
(
    id         bigserial primary key,
    first_name text        not null default '',
    last_name  text        not null default '',
    patronymic text                 default '',
    email      text unique not null,
    phone      text unique not null,
    role       smallint    not null default 0,
    is_active  boolean     not null default false,
    hire_date  timestamp   not null default now(),
    fire_date  timestamp,
    created_at timestamp   not null default now(),
    updated_at timestamp   not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table employees;
-- +goose StatementEnd
