create table url
(
    id         serial primary key not null,
    long_url   varchar            not null,
    short_url  varchar,
    created_at timestamptz default now()
)