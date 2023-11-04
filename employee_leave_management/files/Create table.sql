-- Create an ENUM for leave_type
CREATE TYPE leave_type AS ENUM (
    'Casual Leave',
    'Earned Leave',
    'Sick Leave'
);
-- Create the leave_records table
CREATE TABLE IF NOT EXISTS public.leave_records
(
    id serial PRIMARY KEY,
    full_name text NOT NULL,
    leave_type leave_type_enum NOT NULL,
    from_date date NOT NULL,
    to_date date NOT NULL,
    team text NOT NULL,
    medical_certificate_url text,
    reporter text NOT NULL
);

-- Create a trigger function for notification
CREATE OR REPLACE FUNCTION public.create_notification() RETURNS TRIGGER AS $$
BEGIN
    -- Your notification logic goes here
    -- You can use NEW to access the newly inserted/updated record
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create the leave_notification_trigger
CREATE TRIGGER leave_notification_trigger
AFTER INSERT OR UPDATE 
ON public.leave_records
FOR EACH ROW
EXECUTE FUNCTION public.create_notification();

-- Set the OWNER of the table
ALTER TABLE public.leave_records
OWNER TO postgres;

-- Set the default tablespace if needed
-- TABLESPACE pg_default;
