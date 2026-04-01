ALTER TABLE organizations
    DROP COLUMN IF EXISTS amount;

ALTER TABLE organizations
    ADD COLUMN IF NOT EXISTS payout_wallet_address VARCHAR(42);
