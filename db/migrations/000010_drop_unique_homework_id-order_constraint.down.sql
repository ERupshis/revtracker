ALTER TABLE IF EXISTS homework_questions
    ADD CONSTRAINT unique_homework_order UNIQUE (homework_id, "order");

DROP TRIGGER IF EXISTS update_homework_questions_orders_trigger ON homework_questions;
DROP FUNCTION IF EXISTS remove_related_contents();

DROP TRIGGER IF EXISTS validate_question_id_trigger ON homework_questions;
DROP FUNCTION IF EXISTS validate_question_id_on_unique();

DROP TRIGGER IF EXISTS delete_homework_questions_on_question_delete_trigger ON questions;
DROP FUNCTION IF EXISTS delete_homework_questions_on_question_delete();
