CREATE TABLE users(
    id         BIGSERIAL PRIMARY KEY,
    name       VARCHAR(255)  NOT NULL,
    email       VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP  NOT NULL DEFAULT now()
);

