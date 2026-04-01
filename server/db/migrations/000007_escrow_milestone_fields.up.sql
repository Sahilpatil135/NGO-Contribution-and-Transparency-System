ALTER TABLE organizations
    ADD COLUMN IF NOT EXISTS payout_wallet_address VARCHAR(42);

ALTER TABLE causes
    ADD COLUMN IF NOT EXISTS escrow_goal_wei TEXT;

COMMENT ON COLUMN causes.escrow_goal_wei IS 'On-chain funding goal in wei for CauseMilestoneEscrow; optional until org enables crypto escrow.';
