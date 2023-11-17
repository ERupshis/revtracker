ALTER TABLE IF EXISTS homework_questions
    DROP CONSTRAINT IF EXISTS unique_homework_order;



CREATE OR REPLACE FUNCTION update_homework_questions_orders()
    RETURNS TRIGGER
    LANGUAGE plpgsql
AS
$$
BEGIN
    IF TG_OP = 'UPDATE' AND NEW.deleted = true THEN
        UPDATE homework_questions
        SET "order" = "order" - 1
        WHERE homework_id = OLD.homework_id
          AND "order" > OLD."order";
    END IF;

    RETURN NEW;
END;
$$;

CREATE OR REPLACE TRIGGER update_homework_questions_orders_trigger
    AFTER UPDATE
    ON homework_questions
    FOR EACH ROW
EXECUTE FUNCTION update_homework_questions_orders();




CREATE OR REPLACE FUNCTION validate_question_id_on_unique()
    RETURNS TRIGGER
    LANGUAGE plpgsql
AS
$$
BEGIN
    IF TG_OP = 'INSERT' OR TG_OP = 'UPDATE' THEN
        IF EXISTS(SELECT true FROM homework_questions
                  WHERE homework_id = NEW.homework_id AND question_id = NEW.question_id AND deleted = false AND NEW.deleted = false) THEN
            RAISE EXCEPTION 'same question already has been added in homework';
        ELSEIF NOT EXISTS(SELECT true FROM questions WHERE id = NEW.question_id AND deleted = false) THEN
            RAISE EXCEPTION 'question is not found';
        END IF;
    END IF;

    RETURN NEW;
END;
$$;

CREATE OR REPLACE TRIGGER validate_question_id_trigger
    BEFORE INSERT OR UPDATE
    ON homework_questions
    FOR EACH ROW
EXECUTE FUNCTION validate_question_id_on_unique();




CREATE OR REPLACE FUNCTION delete_homework_questions_on_question_delete()
    RETURNS TRIGGER
    LANGUAGE plpgsql
AS
$$
BEGIN
    IF TG_OP = 'UPDATE' AND NEW.deleted = true THEN
        UPDATE homework_questions
        SET deleted = true
        WHERE question_id = OLD.id;
    END IF;

    RETURN NEW;
END;
$$;

CREATE OR REPLACE TRIGGER delete_homework_questions_on_question_delete_trigger
    AFTER UPDATE
    ON questions
    FOR EACH ROW
EXECUTE FUNCTION delete_homework_questions_on_question_delete();

