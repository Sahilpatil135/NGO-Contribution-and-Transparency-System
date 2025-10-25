CREATE TABLE IF NOT EXISTS organizations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    organization_name VARCHAR(255) NOT NULL,
    registration_number VARCHAR(100),
    organization_type VARCHAR(31),
    about TEXT,
    website_url VARCHAR(255),
    is_approved BOOLEAN DEFAULT false,
    address TEXT
);


CREATE TABLE IF NOT EXISTS organization_documents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    document_type VARCHAR(50) NOT NULL,
    document_url TEXT NOT NULL,
    is_verified BOOLEAN DEFAULT false,
    uploaded_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS organization_contacts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    role VARCHAR(100),
    email VARCHAR(255),
    phone VARCHAR(20),
    is_primary BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE UNIQUE INDEX iF NOT EXISTS unique_primary_contact
ON organization_contacts (organization_id)
WHERE is_primary = true;
