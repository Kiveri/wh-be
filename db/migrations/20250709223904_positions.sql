-- +goose Up
-- +goose StatementBegin
create table positions
(
    id              bigserial primary key,
    barcode         bigint    not null default 0,
    name            text      not null default '',
    manufacturer    text      not null default '',
    price           bigint    not null default 0,
    position_type   smallint  not null default 0,
    production_date timestamp,
    expiration_date timestamp,
    is_has_order    boolean   not null default false,
    is_active       boolean   not null default false,
    created_at      timestamp not null default now(),
    updated_at      timestamp not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table positions;
-- +goose StatementEnd
