CREATE TABLE IF NOT EXISTS questions (
    id   BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    content_id BIGINT REFERENCES contents(id) NOT NULL
);