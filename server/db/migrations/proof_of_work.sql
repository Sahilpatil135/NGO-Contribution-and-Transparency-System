ALTER TABLE causes
ADD COLUMN execution_lat DOUBLE PRECISION,
ADD COLUMN execution_lng DOUBLE PRECISION,
ADD COLUMN execution_radius_meters INT DEFAULT 200,
ADD COLUMN execution_start_time TIMESTAMP,
ADD COLUMN execution_end_time TIMESTAMP,
ADD COLUMN funding_status VARCHAR(20) DEFAULT 'PENDING';
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
