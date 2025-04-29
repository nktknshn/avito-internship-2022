-- +goose Up

CREATE TABLE accounts (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL UNIQUE,
    balance_available BIGINT NOT NULL DEFAULT 0,
    balance_reserved BIGINT NOT NULL DEFAULT 0
);

CREATE INDEX idx_accounts_user_id ON accounts USING HASH (user_id);

CREATE TABLE transactions_deposit (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_id BIGINT REFERENCES accounts(id),
    user_id BIGINT NOT NULL,
    deposit_source VARCHAR NOT NULL,
    status VARCHAR(255) NOT NULL,
    amount BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_transactions_deposit_account_id ON transactions_deposit USING HASH (account_id);
CREATE INDEX idx_transactions_deposit_user_id ON transactions_deposit USING HASH (user_id);

CREATE TABLE transactions_spend (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_id BIGINT REFERENCES accounts(id),
    user_id BIGINT NOT NULL,
    order_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    product_title TEXT NOT NULL,
    status VARCHAR(255) NOT NULL,
    amount BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_transactions_spend_account_id ON transactions_spend USING HASH (account_id);
CREATE INDEX idx_transactions_spend_user_id ON transactions_spend USING HASH (user_id);
CREATE INDEX idx_transactions_spend_order_id ON transactions_spend USING HASH (order_id);
CREATE INDEX idx_transactions_spend_product_id ON transactions_spend USING HASH (product_id);

CREATE TABLE transactions_transfer (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    from_account_id BIGINT REFERENCES accounts(id),
    to_account_id BIGINT REFERENCES accounts(id),
    amount BIGINT NOT NULL,
    status VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_transactions_transfer_from_account_id ON transactions_transfer USING HASH (from_account_id);
CREATE INDEX idx_transactions_transfer_to_account_id ON transactions_transfer USING HASH (to_account_id);

-- +goose Down

DROP TABLE transactions_transfer;
DROP TABLE transactions_spend;
DROP TABLE transactions_deposit;
DROP TABLE accounts;
