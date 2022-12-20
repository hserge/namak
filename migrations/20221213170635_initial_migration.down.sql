BEGIN;
DROP TABLE IF EXISTS transactions CASCADE;
-- DROP TRIGGER IF EXISTS clean_email ON transactions CASCADE;
DROP FUNCTION IF EXISTS clean_email() CASCADE;
DROP TYPE IF EXISTS transaction_kind;

DROP TABLE IF EXISTS unsubscribed;
DROP TABLE IF EXISTS campaigns_domains;
DROP TABLE IF EXISTS domains;
DROP TABLE IF EXISTS campaigns_templates;
DROP TABLE IF EXISTS templates;
DROP TABLE IF EXISTS campaigns_emails;
DROP TABLE IF EXISTS campaigns_emails;
DROP TABLE IF EXISTS emails;
DROP TABLE IF EXISTS campaigns;
COMMIT;
