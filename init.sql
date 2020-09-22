create table url
(
    id         serial primary key not null,
    url        varchar            not null,
    name       varchar unique,
    created_at timestamptz default now()
);

create index on url(name);
