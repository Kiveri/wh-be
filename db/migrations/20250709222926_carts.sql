-- +goose Up
-- +goose StatementBegin
create table carts
(
    id            bigserial primary key,
    client_id     bigint    not null default 0,
    positions_ids bigint[]  not null default '{}',
    total_price   bigint    not null default 0,
    status        smallint  not null default 0,
    is_paid       boolean   not null default false,
    is_active     boolean   not null default false,
    created_at    timestamp not null default now(),
    updated_at    timestamp not null default now()
);

create unique index idx_carts_client_id_is_active on carts (client_id) where is_active = true;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table carts;
-- +goose StatementEnd