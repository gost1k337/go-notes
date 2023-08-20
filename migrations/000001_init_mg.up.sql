CREATE TABLE IF NOT EXISTS notes (
    id serial not null primary key,
    title varchar not null,
    description varchar,
    reminds_at timestamptz[] null,
    deleted bool not null default false,
    created_at timestamp not null default now()
);
