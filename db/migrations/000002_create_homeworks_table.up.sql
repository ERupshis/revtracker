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

        /*
        WITH unused_questions AS (
            SELECT
                question_id AS id
            FROM homework_questions
        )
        DELETE FROM questions WHERE id NOT IN (SELECT id FROM unused_questions);
        */
    END IF;

    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER remove_homework_questions_trigger
    BEFORE DELETE
    ON homeworks
    FOR EACH ROW
EXECUTE FUNCTION remove_related_homework_questions();