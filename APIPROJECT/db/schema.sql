CREATE TYPE status AS ENUM (
  'read',
  'reading',
  'to_read'
);

CREATE TABLE users (
  id   BIGSERIAL PRIMARY KEY,
  username text      NOT NULL,
  Password  text    NOT NULL
);

CREATE TABLE book {
    id  BIGSERIAL PRIMARY KEY,
    title   text    NOT NULL,
    status  status    NOT NULL,
    author  text,
    year    INTEGER,
    userid  INTEGER
}