CREATE TYPE new_type AS (new_options varchar , some_int bigint);
CREATE TABLE polls (
    id bigserial not null primary key,
    poll new_type[] not null
)