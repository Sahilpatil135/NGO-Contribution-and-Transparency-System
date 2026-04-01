ALTER TABLE organizations
    DROP COLUMN IF EXISTS payout_wallet_address;

ALTER TABLE organizations
    ADD COLUMN IF NOT EXISTS amount NUMERIC(20, 2) DEFAULT 0.00;

COMMENT ON COLUMN organizations.amount IS 'Total amount released by escrow smart contract for this organization';
