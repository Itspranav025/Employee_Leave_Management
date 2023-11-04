--Create the Notifications Table:
CREATE TABLE IF NOT EXISTS public.Notifications
(
    id serial PRIMARY KEY,
    Reporting_Manager text COLLATE pg_catalog."default" NOT NULL,
    Leave_ID serial NOT NULL,
    approved boolean DEFAULT false
);

--Create a Trigger Function:
CREATE OR REPLACE FUNCTION notify_reporting_manager()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO public.Notifications (Reporting_Manager, Leave_ID)
    VALUES (NEW.reporter, NEW.id);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

--Create the trigger
CREATE TRIGGER leave_record_notification
AFTER INSERT OR UPDATE ON public.leave_records
FOR EACH ROW
EXECUTE FUNCTION notify_reporting_manager();
