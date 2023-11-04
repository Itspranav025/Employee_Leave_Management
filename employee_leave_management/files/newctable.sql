-- Table: public.leave_records
DROP TYPE IF EXISTS leave_type_enum;
DROP TYPE IF EXISTS leave_type;
-- DROP TABLE IF EXISTS public.leave_records;
CREATE TYPE leave_type AS ENUM (
    'Casual Leave',
    'Earned leave',
    'Sick Leave'
);
CREATE TYPE leave_type_enum AS ENUM (
    'Casual Leave',
    'Earned leave',
    'Sick Leave'
);


CREATE TABLE IF NOT EXISTS public.leave_records
(
    id bigint NOT NULL DEFAULT nextval('leave_records_id_seq'::regclass),
    full_name text COLLATE pg_catalog."default" NOT NULL,
    leave_type leave_type NOT NULL,
    from_date date NOT NULL,
    to_date date NOT NULL,
    team text COLLATE pg_catalog."default" NOT NULL,
    medical_certificate_url text COLLATE pg_catalog."default",
    reporter text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT leave_records_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.leave_records
    OWNER to postgres;