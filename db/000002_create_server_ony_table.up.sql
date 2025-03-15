CREATE TABLE IF NOT EXISTS server_only (
    uuid uuid NOT NULL PRIMARY KEY,
    username text,
    first_name text,
    last_name text,
    identity text,
    steam_id text
);
