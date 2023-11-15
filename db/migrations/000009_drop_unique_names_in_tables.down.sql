ALTER TABLE IF EXISTS homeworks
    ADD CONSTRAINT homeworks_name_key UNIQUE (name);

ALTER TABLE IF EXISTS questions
    ADD CONSTRAINT questions_name_key UNIQUE (name);
