--ROLES IMPLEMENTATION
CREATE TABLE IF NOT EXISTS roles
(
    id   SMALLSERIAL PRIMARY KEY,
    value VARCHAR(10) NOT NULL UNIQUE
);

INSERT INTO roles(value)
VALUES ('USER'),
       ('ADMIN');

--USERS IMPLEMENTATION
CREATE TABLE IF NOT EXISTS users
(
    id   SERIAL PRIMARY KEY,
    login VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL UNIQUE,
    role_id SMALLINT REFERENCES roles(id)
);

INSERT INTO users(login, password, name, role_id)
VALUES ('admin', 'pwd', 'name', 1);
