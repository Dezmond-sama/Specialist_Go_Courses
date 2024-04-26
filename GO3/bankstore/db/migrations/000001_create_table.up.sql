CREATE TYPE Currency AS ENUM ('USD', 'EUR', 'RUB');

CREATE TABLE IF NOT EXISTS accounts (
    id BIGSERIAL PRIMARY KEY,
    owner VARCHAR(255) NOT NULL,
    balance BIGINT NOT NULL,
    currency Currency NOT NULL,
    created TIMESTAMP NOT NULL default now()
);

CREATE TABLE entries (
    id BIGSERIAL PRIMARY KEY,
    account_id BIGINT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    amount BIGINT NOT NULL,
    created TIMESTAMP NOT NULL default now()
);

CREATE TABLE transfers (
    id BIGSERIAL PRIMARY KEY,
    from_account_id BIGINT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    to_account_id BIGINT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    amount BIGINT NOT NULL,
    created TIMESTAMP NOT NULL default now()
);

CREATE INDEX ON accounts (owner);

CREATE INDEX ON entries (account_id);

CREATE INDEX ON transfers (from_account_id);

CREATE INDEX ON transfers (to_account_id);

CREATE INDEX ON transfers (from_account_id, to_account_id);