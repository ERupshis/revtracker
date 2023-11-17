DO
$$
    BEGIN
        IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'questions') THEN
            DROP TRIGGER IF EXISTS remove_contents_trigger ON questions;
        END IF;
    END
$$;

DROP TABLE IF EXISTS questions;
DROP FUNCTION IF EXISTS remove_related_contents();