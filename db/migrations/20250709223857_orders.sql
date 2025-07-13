-- +goose Up
-- +goose StatementBegin
create table orders
(
    id               bigserial primary key,
    client_id        bigint    not null default 0,
    postings_ids     bigint[]  not null default '{}',
    status           smallint  not null default 0,
    delivery_type    smallint  not null default 0,
    delivery_address text      not null default '',
    is_active        boolean   not null default false,
    created_at       timestamp not null default now(),
    updated_at       timestamp not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table orders;
-- +goose StatementEnd
