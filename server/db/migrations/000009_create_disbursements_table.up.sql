CREATE TABLE IF NOT EXISTS disbursements (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    cause_id UUID NOT NULL REFERENCES causes(id) ON DELETE CASCADE,
    milestone_number INT NOT NULL CHECK (milestone_number BETWEEN 1 AND 4),
    amount NUMERIC(20, 2) NOT NULL,
    transaction_hash VARCHAR(66),
    disbursed_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(cause_id, milestone_number)
);

CREATE INDEX idx_disbursements_organization_id ON disbursements(organization_id);
CREATE INDEX idx_disbursements_cause_id ON disbursements(cause_id);
CREATE INDEX idx_disbursements_disbursed_at ON disbursements(disbursed_at DESC);

COMMENT ON TABLE disbursements IS 'Tracks milestone-based disbursements from escrow to organizations';
COMMENT ON COLUMN disbursements.milestone_number IS 'Milestone number (1=25%, 2=50%, 3=75%, 4=100%)';
COMMENT ON COLUMN disbursements.amount IS 'Amount disbursed at this milestone';
COMMENT ON COLUMN disbursements.transaction_hash IS 'Blockchain transaction hash for the disbursement';
