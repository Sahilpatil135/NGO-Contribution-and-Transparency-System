CREATE TABLE IF NOT EXISTS cause_votes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    cause_id UUID NOT NULL REFERENCES causes(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    vote_value SMALLINT NOT NULL CHECK (vote_value IN (1, -1)),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE (cause_id, user_id)
);

CREATE INDEX IF NOT EXISTS idx_cause_votes_cause_id ON cause_votes(cause_id);
CREATE INDEX IF NOT EXISTS idx_cause_votes_user_id ON cause_votes(user_id);

