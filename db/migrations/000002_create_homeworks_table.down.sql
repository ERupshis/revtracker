DO
$$
    BEGIN
        IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'homeworks') THEN
            DROP TRIGGER IF EXISTS remove_homework_questions_trigger ON homeworks;
        END IF;
    END
$$;

DROP TABLE IF EXISTS homeworks;
DROP FUNCTION IF EXISTS remove_related_homework_questions;