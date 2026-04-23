DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'ngo_verification_status') THEN
        CREATE TYPE ngo_verification_status AS ENUM ('pending', 'approved', 'rejected');
    END IF;
END
$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'dispute_status') THEN
        CREATE TYPE dispute_status AS ENUM ('open', 'in_review', 'resolved', 'dismissed');
    END IF;
END
$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'dispute_priority') THEN
        CREATE TYPE dispute_priority AS ENUM ('low', 'medium', 'high', 'critical');
    END IF;
END
$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'milestone_status') THEN
        CREATE TYPE milestone_status AS ENUM ('not_started', 'in_progress', 'completed', 'verified', 'blocked');
    END IF;
END
$$;

CREATE TABLE IF NOT EXISTS ngo_verification_requests (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    submitted_by UUID REFERENCES users(id) ON DELETE SET NULL,
    status ngo_verification_status NOT NULL DEFAULT 'pending',
    review_notes TEXT,
    reviewed_by UUID REFERENCES users(id) ON DELETE SET NULL,
    reviewed_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_ngo_verification_requests_status
    ON ngo_verification_requests(status);

CREATE INDEX IF NOT EXISTS idx_ngo_verification_requests_org
    ON ngo_verification_requests(organization_id);

CREATE TABLE IF NOT EXISTS disputes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    cause_id UUID REFERENCES causes(id) ON DELETE SET NULL,
    opened_by UUID REFERENCES users(id) ON DELETE SET NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    evidence_url TEXT,
    status dispute_status NOT NULL DEFAULT 'open',
    priority dispute_priority NOT NULL DEFAULT 'medium',
    resolution_notes TEXT,
    resolved_by UUID REFERENCES users(id) ON DELETE SET NULL,
    resolved_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_disputes_status
    ON disputes(status);

CREATE INDEX IF NOT EXISTS idx_disputes_org
    ON disputes(organization_id);

CREATE INDEX IF NOT EXISTS idx_disputes_cause
    ON disputes(cause_id);

CREATE TABLE IF NOT EXISTS cause_milestones (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    cause_id UUID NOT NULL REFERENCES causes(id) ON DELETE CASCADE,
    milestone_number INT NOT NULL CHECK (milestone_number BETWEEN 1 AND 10),
    title VARCHAR(255) NOT NULL,
    details TEXT,
    target_amount NUMERIC(12, 2),
    due_date TIMESTAMP WITH TIME ZONE,
    completion_percent NUMERIC(5, 2) NOT NULL DEFAULT 0 CHECK (completion_percent BETWEEN 0 AND 100),
    status milestone_status NOT NULL DEFAULT 'not_started',
    verified_by UUID REFERENCES users(id) ON DELETE SET NULL,
    verified_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    UNIQUE (cause_id, milestone_number)
);

CREATE INDEX IF NOT EXISTS idx_cause_milestones_cause
    ON cause_milestones(cause_id);

CREATE INDEX IF NOT EXISTS idx_cause_milestones_status
    ON cause_milestones(status);

CREATE TABLE IF NOT EXISTS ngo_trust_scores (
    organization_id UUID PRIMARY KEY REFERENCES organizations(id) ON DELETE CASCADE,
    verification_score NUMERIC(5, 2) NOT NULL DEFAULT 0 CHECK (verification_score BETWEEN 0 AND 100),
    donor_rating_score NUMERIC(5, 2) NOT NULL DEFAULT 0 CHECK (donor_rating_score BETWEEN 0 AND 100),
    milestone_completion_score NUMERIC(5, 2) NOT NULL DEFAULT 0 CHECK (milestone_completion_score BETWEEN 0 AND 100),
    overall_score NUMERIC(5, 2) NOT NULL DEFAULT 0 CHECK (overall_score BETWEEN 0 AND 100),
    last_calculated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_ngo_trust_scores_overall
    ON ngo_trust_scores(overall_score DESC);

CREATE TABLE IF NOT EXISTS admin_action_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    admin_user_id UUID REFERENCES users(id) ON DELETE SET NULL,
    action_type VARCHAR(100) NOT NULL,
    target_type VARCHAR(50) NOT NULL,
    target_id UUID,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_admin_action_logs_admin_user
    ON admin_action_logs(admin_user_id);

CREATE INDEX IF NOT EXISTS idx_admin_action_logs_target
    ON admin_action_logs(target_type, target_id);
