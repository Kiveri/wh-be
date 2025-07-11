-- +goose Up
-- +goose StatementBegin
create table carts
(
    id            bigserial primary key,
    client_id     bigint    not null default 0,
    positions_ids bigint[]  not null default '{}',
    total_price   bigint    not null default 0,
    is_paid       boolean   not null default false,
    is_active     boolean   not null default false,
    created_at    timestamp not null default now(),
    updated_at    timestamp not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table carts
-- +goose StatementEnd
