CREATE TABLE IF NOT EXISTS cause_reviews (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    cause_id UUID NOT NULL REFERENCES causes(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    review_text TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE (cause_id, user_id)
);

CREATE INDEX IF NOT EXISTS idx_cause_reviews_cause_id ON cause_reviews(cause_id);
CREATE INDEX IF NOT EXISTS idx_cause_reviews_user_id ON cause_reviews(user_id);

