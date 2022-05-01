-- migrate:up
CREATE TABLE users (
  id        BIGSERIAL PRIMARY KEY,
  username  text      NOT NULL
);
