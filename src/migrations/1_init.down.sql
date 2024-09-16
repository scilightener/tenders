DROP INDEX IF EXISTS idx_bid_negotiation_bid_id;

DROP TABLE IF EXISTS bid_negotiation;

DROP INDEX IF EXISTS idx_bid_user_reviews_user_id;

DROP TABLE IF EXISTS bid_user_reviews;

DROP TABLE IF EXISTS bid_versions;

DROP INDEX IF EXISTS idx_bid_tender_id;

DROP TABLE IF EXISTS bid;

DROP TABLE IF EXISTS tender_versions;

DROP INDEX IF EXISTS idx_tender_creator_id;

DROP INDEX IF EXISTS idx_tender_service_type;

DROP TABLE IF EXISTS tender;

DROP INDEX IF EXISTS idx_organization_responsible_user_id;

DROP TYPE IF EXISTS bid_author_type;

DROP TYPE IF EXISTS bid_status;

DROP TYPE IF EXISTS bid_author_type;

DROP TYPE IF EXISTS tender_service_type;

DROP TYPE IF EXISTS tender_status;