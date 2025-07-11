-- +goose Up
-- +goose StatementBegin
create table companies
(
    id            bigserial primary key,
    name          text unique   not null,
    inn           bigint unique not null,
    employees_ids bigint[]      not null default '{}',
    legal_address text          not null default '',
    is_active     boolean       not null default false,
    created_at    timestamp     not null default now(),
    updated_at    timestamp     not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table companies;
-- +goose StatementEnd
