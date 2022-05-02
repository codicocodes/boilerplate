CREATE TABLE users (
  id        BIGSERIAL PRIMARY KEY,
  username  text      NOT NULL
);

ALTER TABLE users ADD COLUMN password text NOT NULL;

ALTER TABLE users ADD UNIQUE (username);


