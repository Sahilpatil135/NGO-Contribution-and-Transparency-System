ALTER TABLE causes
ADD COLUMN execution_lat DOUBLE PRECISION,
ADD COLUMN execution_lng DOUBLE PRECISION,
ADD COLUMN execution_radius_meters INT DEFAULT 200,
ADD COLUMN execution_start_time TIMESTAMP,
ADD COLUMN execution_end_time TIMESTAMP,
ADD COLUMN funding_status VARCHAR(20) DEFAULT 'PENDING',
ADD COLUMN beneficiaries_count INTEGER NOT NULL DEFAULT 0,
ADD COLUMN execution_location TEXT,
ADD COLUMN impact_goal TEXT,
ADD COLUMN problem_statement TEXT,
ADD COLUMN execution_plan TEXT,
ADD COLUMN donor_count INTEGER DEFAULT 0,
ADD COLUMN updated_at TIMESTAMP DEFAULT NOW();

-- Cause Products Table
CREATE TABLE cause_products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    cause_id UUID REFERENCES causes(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    price_per_unit NUMERIC(12,2) NOT NULL CHECK (price_per_unit > 0),
    quantity_needed INTEGER NOT NULL CHECK (quantity_needed > 0),
    quantity_funded INTEGER DEFAULT 0 CHECK (quantity_funded >= 0),
    image_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);


CREATE TYPE update_type_enum AS ENUM (
    'Engagement',      -- During fundraising
    'Milestone',       -- 25%, 50%, 75%
    'Execution',       -- After funds released
    'Completion'       -- Final report
);

-- Cause Updates Table
CREATE TABLE cause_updates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    cause_id UUID REFERENCES causes(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    update_type update_type_enum NOT NULL,
    funding_percentage INTEGER CHECK (funding_percentage >= 0 AND funding_percentage <= 100),
    is_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW()
);


-- Cause Update Media Table
CREATE TABLE update_media (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    update_id UUID REFERENCES cause_updates(id) ON DELETE CASCADE,
    media_type TEXT CHECK (media_type IN ('image','receipt','pdf')),
    media_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Donations Table
ALTER TABLE donations
ADD COLUMN product_id UUID REFERENCES cause_products(id);


CREATE TYPE funding_status_enum AS ENUM (
    'Not Started',
    'Active',
    'Fully Funded',
    'Closed'
);

ALTER TABLE causes
ALTER COLUMN funding_status DROP DEFAULT;

UPDATE causes
SET funding_status = 'Not Started'
WHERE funding_status = 'PENDING';

ALTER TABLE causes
ALTER COLUMN funding_status TYPE funding_status_enum
USING funding_status::funding_status_enum;

ALTER TABLE causes
ALTER COLUMN funding_status SET DEFAULT 'Not Started';


CREATE TABLE proof_sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID REFERENCES organizations(id),
    cause_id UUID REFERENCES causes(id),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW()
);


CREATE TABLE proof_images (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    session_id UUID REFERENCES proof_sessions(id),
    image_hash TEXT NOT NULL,
    ipfs_cid TEXT,
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,
    timestamp TIMESTAMP,
    metadata_score INT,
    created_at TIMESTAMP DEFAULT NOW()
);

--  Added on 16/3/2026
ALTER TABLE cause_updates
ADD COLUMN proof_session_id UUID REFERENCES proof_sessions(id);

ALTER TABLE proof_images 
ADD COLUMN final_score NUMERIC(5,2), 
ADD COLUMN verification_score TEXT;

ALTER TABLE proof_sessions
ADD COLUMN total_images INTEGER DEFAULT 0,
ADD COLUMN verified_images INTEGER DEFAULT 0,
ADD COLUMN session_score NUMERIC(5,2);

-- Added on 18/3/2026
CREATE TABLE disbursements (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    cause_id UUID REFERENCES causes(id) ON DELETE CASCADE,

    amount NUMERIC(12,2) NOT NULL,
    milestone_percentage INTEGER, -- e.g. 20, 40, etc.

    tx_hash TEXT, -- blockchain transaction reference

    released_at TIMESTAMP DEFAULT NOW()
);

ALTER TABLE cause_updates DROP COLUMN is_verified;

ALTER TABLE cause_updates
ADD COLUMN claimed_amount NUMERIC(12,2),
ADD COLUMN verification_score NUMERIC(5,2),
ADD COLUMN verification_status TEXT DEFAULT 'pending';

ALTER TABLE proof_images
RENAME COLUMN verification_score TO verification_status;

-- new 
-- Added on 20/3/2026
-- Async receipt verification jobs for Execution updates.
CREATE TABLE receipt_verification_jobs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    receipt_path TEXT NOT NULL,
    claimed_amount NUMERIC(12,2) NOT NULL,
    status TEXT NOT NULL DEFAULT 'pending', -- pending | verified | review | rejected | error
    receipt_score NUMERIC(5,2),
    error_message TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX receipt_verification_jobs_org_status_idx
ON receipt_verification_jobs (organization_id, status);