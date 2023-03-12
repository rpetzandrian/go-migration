CREATE table if not exists users
(
    id          bigint not null
                constraint users_pk
                primary key,
    name       varchar,
    age        integer,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    deleted_at timestamp with time zone
);