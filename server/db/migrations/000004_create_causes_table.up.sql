CREATE TABLE IF NOT EXISTS cause_domains (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name	VARCHAR(255) NOT NULL UNIQUE,
    description	TEXT,
    icon_url	TEXT
);

CREATE TABLE IF NOT EXISTS cause_aid_types (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name	VARCHAR(255) NOT NULL UNIQUE,
    description	TEXT,
    icon_url	TEXT
);

CREATE TABLE IF NOT EXISTS causes (
    id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id     UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    title	        VARCHAR(255) NOT NULL,	
    description	        TEXT,
    domain_id           UUID NOT NULL REFERENCES cause_domains(id) ON DELETE CASCADE,
    aid_type_id         UUID NOT NULL REFERENCES cause_aid_types(id) ON DELETE CASCADE,
    collected_amount    NUMERIC(12,2) NOT NULL DEFAULT 0,
    goal_amount	        NUMERIC(12,2),
    deadline	        TIMESTAMP WITH TIME ZONE,
    is_active	        BOOLEAN DEFAULT true,
    cover_image_url     TEXT,
    created_at          TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TYPE donation_status AS ENUM ('pending', 'paid', 'failed', 'refunded');

CREATE TABLE IF NOT EXISTS donations (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    cause_id        UUID NOT NULL REFERENCES causes(id) ON DELETE CASCADE,
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    amount	    NUMERIC(12,2) NOT NULL,
    status	    donation_status DEFAULT 'pending',
    payment_id	    VARCHAR(255),
    created_at	    TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
