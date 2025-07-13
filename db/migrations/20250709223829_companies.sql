-- +goose Up
-- +goose StatementBegin
create table companies
(
    id            bigserial primary key,
    name          text unique   not null,
    inn           bigint unique not null,
    owners_ids    bigint[]      not null default '{}',
    legal_address text          not null default '',
    type          smallint      not null default 0,
    is_active     boolean       not null default false,
    created_at    timestamp     not null default now(),
    updated_at    timestamp     not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table companies;
-- +goose StatementEnd
