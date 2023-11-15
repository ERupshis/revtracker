CREATE OR REPLACE FUNCTION remove_related_contents()
    RETURNS TRIGGER AS
$$
BEGIN
    IF TG_OP = 'DELETE' THEN
        UPDATE contents SET deleted = true WHERE id = OLD.content_id;
    END IF;

    IF TG_OP = 'UPDATE' AND (OLD.content_id IS DISTINCT FROM NEW.content_id OR NEW.deleted = true) THEN
        IF (OLD.content_id IS DISTINCT FROM NEW.content_id OR NEW.deleted = true) THEN
            UPDATE contents SET deleted = true WHERE id = OLD.content_id;
        ELSEIF (OLD.content_id = NEW.content_id AND OLD.deleted = true AND NEW.deleted = false) THEN
            UPDATE contents SET deleted = false WHERE id = OLD.content_id;
        END IF;
    END IF;

    RETURN OLD;
END;
$$ LANGUAGE plpgsql;