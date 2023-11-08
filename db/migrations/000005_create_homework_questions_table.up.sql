CREATE TABLE IF NOT EXISTS homework_questions (
     id   BIGSERIAL PRIMARY KEY,
     homework_id BIGINT REFERENCES homeworks(id) NOT NULL,
     question_id BIGINT REFERENCES questions(id) NOT NULL,
     "order" BIGINT NOT NULL,

     CONSTRAINT unique_homework_question_order UNIQUE (homework_id, question_id, "order")
);
