-- Table: public.leave_types

-- DROP TABLE IF EXISTS public.leave_types;


CREATE TABLE IF NOT EXISTS public.leave_types
(
    id bigint NOT NULL DEFAULT nextval('leave_types_id_seq'::regclass),
    name text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT leave_types_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.leave_types
    OWNER to postgres;