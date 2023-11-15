ALTER TABLE IF EXISTS homework_questions
    DROP CONSTRAINT IF EXISTS unique_homework_question;

ALTER TABLE IF EXISTS users
    DROP CONSTRAINT IF EXISTS users_name_key;