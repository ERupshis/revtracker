CREATE TABLE IF NOT EXISTS questions
(
    id         BIGSERIAL PRIMARY KEY,
    name       TEXT UNIQUE                     NOT NULL,
    content_id BIGINT REFERENCES contents (id) NOT NULL
);

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

CREATE TRIGGER remove_contents_trigger
    AFTER DELETE OR UPDATE
    ON questions
    FOR EACH ROW
EXECUTE FUNCTION remove_related_contents();


CREATE OR REPLACE FUNCTION remove_related_homework_questions_from_questions()
    RETURNS TRIGGER AS
$$
BEGIN
    IF TG_OP = 'DELETE' THEN
        DELETE FROM homework_questions WHERE OLD.id = question_id;
    END IF;

    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER remove_homework_questions_from_questions_trigger
    AFTER DELETE
    ON questions
    FOR EACH ROW
EXECUTE FUNCTION remove_related_homework_questions_from_questions();