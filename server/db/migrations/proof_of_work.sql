ALTER TABLE causes
ADD COLUMN execution_lat DOUBLE PRECISION,
ADD COLUMN execution_lng DOUBLE PRECISION,
ADD COLUMN execution_radius_meters INT DEFAULT 200,
ADD COLUMN execution_start_time TIMESTAMP,
ADD COLUMN execution_end_time TIMESTAMP,
ADD COLUMN funding_status VARCHAR(20) DEFAULT 'PENDING';


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
