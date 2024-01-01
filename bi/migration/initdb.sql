DROP DATABASE IF EXISTS e;

CREATE DATABASE e;

\connect e

CREATE TABLE example(
                        key TEXT NOT NULL,
                        value INTEGER NOT NULL
);

INSERT INTO example (key, value)
VALUES
    ('key_1', 10),
    ('key_2', 15),
    ('key_3', 20),
    ('key_4', 25),
    ('key_5', 30);
