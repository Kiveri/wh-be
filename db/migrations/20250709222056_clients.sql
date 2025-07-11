-- +goose Up
-- +goose StatementBegin
create table clients
(
    id           bigserial primary key,
    first_name   text        not null default '',
    last_name    text        not null default '',
    patronymic   text,
    email        text unique not null,
    phone        text unique not null,
    home_address text        not null default '',
    company_id   bigint,
    is_active    boolean     not null default false,
    created_at   timestamp   not null default now(),
    updated_at   timestamp   not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table clients;
-- +goose StatementEnd
