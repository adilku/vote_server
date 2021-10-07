CREATE TABLE IF NOT EXISTS wallets
(
    id        bigserial not null primary key,
    cur_balance   bigserial not null CHECK ( cur_balance >= 0 )
);