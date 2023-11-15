CREATE OR REPLACE FUNCTION remove_related_contents()
    RETURNS TRIGGER AS
$$
BEGIN
    IF TG_OP = 'DELETE' THEN
        DELETE FROM contents WHERE id = OLD.content_id;
    END IF;

    IF TG_OP = 'UPDATE' AND OLD.content_id IS DISTINCT FROM NEW.content_id THEN
        DELETE FROM contents WHERE id = OLD.content_id;
    END IF;

    RETURN OLD;
END;
$$ LANGUAGE plpgsql;