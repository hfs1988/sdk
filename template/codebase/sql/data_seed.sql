CREATE TABLE users
(
    id serial PRIMARY KEY,
    name text,
    email text UNIQUE,
    age integer,
    status boolean DEFAULT true
);

INSERT INTO users (name, email, age) VALUES ('Husni Firmansyah', 'husni.firmansyah@gmail.com', 34);