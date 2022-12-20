BEGIN;
-- Table: public.campaigns

-- DROP TABLE IF EXISTS public.campaigns;

CREATE TABLE IF NOT EXISTS public.campaigns
(
    id serial,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name text NOT NULL,
    description text,
    CONSTRAINT campaigns_pkey PRIMARY KEY (id),
    CONSTRAINT campaigns_uidx_name UNIQUE (name)
);

ALTER TABLE IF EXISTS public.campaigns
    OWNER to postgres;

-- Table: public.emails

-- CREATE EXTENSION citext;
-- CREATE DOMAIN email AS citext
--     -- html5 regex from https://html.spec.whatwg.org/multipage/input.html#e-mail-state-(type=email)
--     CHECK ( value ~ '^[a-zA-Z0-9.!#$%&''*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$' );

-- DROP TABLE IF EXISTS public.emails;

CREATE TABLE IF NOT EXISTS public.emails
(
    id serial,
    campaign_id integer,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    email text NOT NULL,
    first_name text NOT NULL DEFAULT 'Valued'::text,
    last_name text NOT NULL DEFAULT 'Customer'::text,
    container json NOT NULL DEFAULT '{}'::json,
    CONSTRAINT emails_pkey PRIMARY KEY (id),
    CONSTRAINT emails_uidx_campaign_id_email UNIQUE (campaign_id, email),
    CONSTRAINT emails_fk_campaign_id FOREIGN KEY (campaign_id)
        REFERENCES public.campaigns (id) MATCH SIMPLE
        ON UPDATE SET NULL
        ON DELETE SET NULL
);

ALTER TABLE IF EXISTS public.emails
    OWNER to postgres;

-- create email validation trigger

CREATE OR REPLACE FUNCTION clean_email() RETURNS trigger AS $clean_email$
BEGIN
    -- select REGEXP_REPLACE(' !"#$%&''()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz{|}~', '["`/:;,<>\\ ''^\[\]]+','','g')
    -- !#$%&()*+-.0123456789=?@ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz{|}~
    -- !#$%&()*+-.0123456789=?@abcdefghijklmnopqrstuvwxyz_abcdefghijklmnopqrstuvwxyz{|}~ (after trigger)
    NEW.email = LOWER(REGEXP_REPLACE(NEW.email, '["`/:;,<>\\ ''^\[\]]+','','g'));
    RETURN NEW;
END;
$clean_email$ LANGUAGE plpgsql;

CREATE TRIGGER clean_email BEFORE INSERT OR UPDATE ON public.emails
    FOR EACH ROW EXECUTE PROCEDURE clean_email();


-- Table: public.campaigns_emails

-- DROP TABLE IF EXISTS public.campaigns_emails;

CREATE TABLE IF NOT EXISTS public.campaigns_emails
(
    id serial,
    campaign_id integer,
    email_id integer,
    CONSTRAINT campaigns_emails_pkey PRIMARY KEY (id),
    CONSTRAINT campaigns_emails_fk_campaign_id FOREIGN KEY (campaign_id)
        REFERENCES public.campaigns (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT campaigns_emails_fk_email_id FOREIGN KEY (email_id)
        REFERENCES public.emails (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

ALTER TABLE IF EXISTS public.campaigns_emails
    OWNER to postgres;

-- Table: public.transactions

CREATE TYPE transaction_kind AS ENUM ('send', 'open', 'click','delayed');

-- DROP TABLE IF EXISTS public.transactions;

CREATE TABLE IF NOT EXISTS public.transactions
(
    id serial,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    campaign_id integer,
    email_id integer,
    kind transaction_kind NOT NULL,
    container json,
    is_success boolean,
    CONSTRAINT transactions_pkey PRIMARY KEY (id),
    CONSTRAINT transactions_fk_email_id FOREIGN KEY (email_id)
        REFERENCES public.emails (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT transactions_fk_campaign_id FOREIGN KEY (campaign_id)
        REFERENCES public.campaigns (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

ALTER TABLE IF EXISTS public.transactions
    OWNER to postgres;

COMMIT;