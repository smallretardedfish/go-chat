CREATE TABLE users(
    id         BIGSERIAL PRIMARY KEY,
    name       TEXT UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);
