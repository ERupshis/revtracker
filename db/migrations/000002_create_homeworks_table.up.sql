CREATE TABLE IF NOT EXISTS homeworks (
    id BIGSERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

CREATE OR REPLACE FUNCTION remove_related_homework_questions()
    RETURNS TRIGGER AS
$$
BEGIN
    IF TG_OP = 'DELETE' THEN
        DELETE FROM homework_questions WHERE OLD.id = homework_id;
    END IF;

    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER remove_homework_questions_trigger
    AFTER DELETE
    ON homeworks
    FOR EACH ROW
EXECUTE FUNCTION remove_related_homework_questions();