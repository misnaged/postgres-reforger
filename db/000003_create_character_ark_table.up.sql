CREATE TABLE IF NOT EXISTS character_ark (
    char_uuid uuid NOT NULL PRIMARY KEY,
    character jsonb
);
