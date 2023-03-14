CREATE table if not exists roles
(
    id          bigint not null
                constraint roles_pk
                primary key,
    roles       varchar,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    deleted_at timestamp with time zone
);