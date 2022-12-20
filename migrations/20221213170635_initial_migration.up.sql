-- This script was generated by the ERD tool in pgAdmin 4.
-- Please log an issue at https://redmine.postgresql.org/projects/pgadmin4/issues/new if you find any bugs, including reproduction steps.
BEGIN;

CREATE TABLE IF NOT EXISTS public.campaigns
(
    id serial,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name text COLLATE pg_catalog."default" NOT NULL,
    description text COLLATE pg_catalog."default",
    CONSTRAINT campaigns_pkey PRIMARY KEY (id),
    CONSTRAINT campaigns_uidx_name UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS public.emails
(
    id serial,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_active boolean NOT NULL DEFAULT false,
    is_success boolean NOT NULL DEFAULT false,
    email text COLLATE pg_catalog."default" NOT NULL,
    first_name text COLLATE pg_catalog."default" NOT NULL DEFAULT 'Valued'::text,
    last_name text COLLATE pg_catalog."default" NOT NULL DEFAULT 'Customer'::text,
    container json NOT NULL DEFAULT '{}'::json,
    CONSTRAINT emails_pkey PRIMARY KEY (id),
    CONSTRAINT emails_uidx_email UNIQUE (email)
);

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

CREATE TABLE IF NOT EXISTS public.campaigns_emails
(
    campaign_id integer NOT NULL,
    email_id integer NOT NULL,
    CONSTRAINT campaigns_emails_pkey PRIMARY KEY (campaign_id, email_id),
    CONSTRAINT campaigns_emails_fk_campaign_id FOREIGN KEY (campaign_id)
        REFERENCES public.campaigns (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT campaigns_emails_fk_email_id FOREIGN KEY (email_id)
        REFERENCES public.emails (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS public.templates
(
    id serial,
    created_at time with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_active boolean NOT NULL DEFAULT false,
    name text NOT NULL,
    description text,
    CONSTRAINT templates_pkey PRIMARY KEY (id),
    CONSTRAINT templates_uidx_name UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS public.campaigns_templates
(
    campaign_id integer NOT NULL,
    template_id integer NOT NULL,
    CONSTRAINT campaigns_templates_pkey PRIMARY KEY (campaign_id, template_id),
    CONSTRAINT campaigns_templates_fk_campaign_id FOREIGN KEY (campaign_id)
        REFERENCES public.campaigns (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT campaigns_templates_fk_template_id FOREIGN KEY (template_id)
        REFERENCES public.templates (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS public.domains
(
    id serial,
    created_at time with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_active boolean NOT NULL DEFAULT false,
    name text NOT NULL,
    description text,
    CONSTRAINT domains_pkey PRIMARY KEY (id),
    CONSTRAINT domains_uidx_name UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS public.campaigns_domains
(
    campaign_id integer NOT NULL,
    domain_id integer NOT NULL,
    CONSTRAINT campaigns_domains_pkey PRIMARY KEY (campaign_id, domain_id),
    CONSTRAINT campaigns_domains_fk_campaign_id FOREIGN KEY (campaign_id)
        REFERENCES public.campaigns (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT campaigns_domains_fk_domain_id FOREIGN KEY (domain_id)
        REFERENCES public.domains (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
);


CREATE TYPE transaction_kind AS ENUM ('send', 'open', 'click','delayed');

CREATE TABLE IF NOT EXISTS public.transactions
(
    id serial,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_success boolean,
    campaign_id integer,
    domain_id integer,
    template_id integer,
    email_id integer,
    email text NOT NULL,
    kind transaction_kind NOT NULL,
    container json,
    CONSTRAINT transactions_pkey PRIMARY KEY (id),
    CONSTRAINT transactions_fk_campaign_id FOREIGN KEY (campaign_id)
        REFERENCES public.campaigns (id) MATCH SIMPLE
        ON UPDATE SET NULL
        ON DELETE SET NULL,
    CONSTRAINT transactions_fk_domain_id FOREIGN KEY (domain_id)
        REFERENCES public.domains (id) MATCH SIMPLE
        ON UPDATE SET NULL
        ON DELETE SET NULL,
    CONSTRAINT transactions_fk_template_id FOREIGN KEY (template_id)
        REFERENCES public.templates (id) MATCH SIMPLE
        ON UPDATE SET NULL
        ON DELETE SET NULL,
    CONSTRAINT transactions_fk_email_id FOREIGN KEY (email_id)
        REFERENCES public.emails (id) MATCH SIMPLE
        ON UPDATE SET NULL
        ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS public.unsubscribed
(
    id serial,
    created_at time with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    transaction_id integer,
    email text NOT NULL,
    CONSTRAINT unsubscribed_pkey PRIMARY KEY (id),
    CONSTRAINT unsubscribed_fk_transaction_id FOREIGN KEY (transaction_id)
        REFERENCES public.transactions (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

CREATE OR REPLACE FUNCTION unsubscribe() RETURNS trigger AS $unsubscribe$
BEGIN
    UPDATE emails SET is_active=false WHERE email=NEW.email;
    RETURN NEW;
END;
$unsubscribe$ LANGUAGE plpgsql VOLATILE;

COMMIT;