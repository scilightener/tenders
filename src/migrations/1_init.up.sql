CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'organization_type') THEN
        CREATE TYPE organization_type AS ENUM (
            'IE',
            'LLC',
            'JSC'
            );
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tender_status') THEN
        CREATE TYPE tender_status AS ENUM ('CREATED', 'PUBLISHED', 'CLOSED');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tender_service_type') THEN
        CREATE TYPE tender_service_type AS ENUM ('CONSTRUCTION', 'DELIVERY', 'MANUFACTURE');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'bid_author_type') THEN
        CREATE TYPE bid_author_type AS ENUM ('ORGANIZATION', 'USER');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'bid_status') THEN
        CREATE TYPE bid_status AS ENUM ('CREATED', 'PUBLISHED', 'CANCELLED');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS employee (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(50) UNIQUE NOT NULL,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS organization (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    type organization_type,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS organization_responsible (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID REFERENCES organization(id) ON DELETE CASCADE,
    user_id UUID REFERENCES employee(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_organization_responsible_user_id ON organization_responsible(user_id);

CREATE TABLE IF NOT EXISTS tender (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500) NOT NULL,
    status tender_status,
    service_type tender_service_type,
    organization_id UUID REFERENCES organization(id) ON DELETE CASCADE,
    creator_id UUID REFERENCES organization_responsible(id) ON DELETE CASCADE,
    version INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_tender_service_type ON tender(service_type);
CREATE INDEX IF NOT EXISTS idx_tender_creator_id ON tender(creator_id);

CREATE TABLE IF NOT EXISTS tender_versions (
    tender_id UUID REFERENCES tender(id) ON DELETE CASCADE,
    version INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500) NOT NULL,
    status tender_status,
    service_type tender_service_type,
    organization_id UUID REFERENCES organization(id) ON DELETE CASCADE,
    creator_id UUID REFERENCES organization_responsible(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT tender_versions_pkey PRIMARY KEY (tender_id, version)
);

CREATE TABLE IF NOT EXISTS bid (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500) NOT NULL,
    status bid_status,
    author_type bid_author_type,
    organization_id UUID REFERENCES organization(id) ON DELETE CASCADE,
    user_id UUID REFERENCES employee(id) ON DELETE CASCADE,
    version INT NOT NULL,
    tender_id UUID REFERENCES tender(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_bid_tender_id ON bid(tender_id);

CREATE TABLE IF NOT EXISTS bid_versions (
    bid_id UUID REFERENCES bid(id) ON DELETE CASCADE,
    version INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500) NOT NULL,
    status bid_status,
    author_type bid_author_type,
    organization_id UUID REFERENCES organization(id) ON DELETE CASCADE,
    user_id UUID REFERENCES employee(id) ON DELETE CASCADE,
    tender_id UUID REFERENCES tender(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT bid_versions_pkey PRIMARY KEY (bid_id, version)
);

CREATE TABLE IF NOT EXISTS bid_user_reviews (
    bid_id UUID REFERENCES bid(id) ON DELETE CASCADE,
    reviewer_id UUID REFERENCES organization_responsible(id) ON DELETE CASCADE,
    description VARCHAR(1000) NOT NULL,
    user_id UUID REFERENCES employee(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT bid_user_reviews_pkey PRIMARY KEY (bid_id, reviewer_id)
);

CREATE INDEX IF NOT EXISTS idx_bid_user_reviews_user_id ON bid_user_reviews(user_id);

CREATE TABLE IF NOT EXISTS bid_negotiation (
    bid_id UUID REFERENCES bid(id) ON DELETE CASCADE,
    employee_id UUID REFERENCES organization_responsible(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT bid_negotiation_pkey PRIMARY KEY (bid_id, employee_id)
);

CREATE INDEX IF NOT EXISTS idx_bid_negotiation_bid_id ON bid_negotiation(bid_id);
