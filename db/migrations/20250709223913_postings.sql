-- +goose Up
-- +goose StatementBegin
create table postings
(
    id             bigserial primary key,
    cart_id        bigint    not null default 0,
    positions_ids  bigint[]  not null default '{}',
    posting_status smallint  not null default 0,
    is_active      boolean   not null default false,
    created_at     timestamp not null default now(),
    updated_at     timestamp not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table postings;
-- +goose StatementEnd
