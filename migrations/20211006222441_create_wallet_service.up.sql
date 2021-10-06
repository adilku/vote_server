CREATE TABLE wallets
(
    id        bigserial not null primary key,
    user_name varchar   not null unique,
    cur_balance   bigserial not null
);