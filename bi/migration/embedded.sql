DROP DATABASE IF EXISTS embedded;

CREATE DATABASE embedded;

\connect embedded

CREATE TABLE source_typ(
                           source_typ_id SERIAL PRIMARY KEY,
                           name TEXT NOT NULL
);

CREATE TABLE source(
                       source_id SERIAL PRIMARY KEY,
                       dsn TEXT NOT NULL,
                       source_typ_id INTEGER NOT NULL
                           REFERENCES source_typ (source_typ_id)
                               ON DELETE CASCADE
);

INSERT INTO source_typ (name)
VALUES ('postgres');
