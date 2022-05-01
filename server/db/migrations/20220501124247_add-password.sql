-- migrate:up
ALTER TABLE users ADD COLUMN password text NOT NULL;


-- migrate:down
ALTER TABLE users DROP COLUMN password;

