-- migrate:up
ALTER TABLE users ADD UNIQUE (username);

-- migrate:down

