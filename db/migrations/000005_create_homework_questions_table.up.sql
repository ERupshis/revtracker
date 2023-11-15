CREATE TABLE IF NOT EXISTS homework_questions (
     id   BIGSERIAL PRIMARY KEY,
     homework_id BIGINT REFERENCES homeworks(id) NOT NULL,
     question_id BIGINT REFERENCES questions(id) NOT NULL,
     "order" BIGINT NOT NULL,

     CONSTRAINT unique_homework_question UNIQUE (homework_id, question_id),
     CONSTRAINT unique_homework_order UNIQUE (homework_id, "order")
);

--TODO: need to add trigger to validate homework_id - question_id unique.
