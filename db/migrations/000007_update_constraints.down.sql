ALTER TABLE IF EXISTS homework_questions
    ADD CONSTRAINT unique_homework_question UNIQUE (homework_id, question_id);

ALTER TABLE IF EXISTS users
    ADD CONSTRAINT users_name_key UNIQUE (name);