CREATE TABLE IF NOT EXISTS users (
    "id" SERIAL PRIMARY KEY,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    "deleted_at" timestamptz,
    "email" varchar NOT NULL,
    "password" varchar NOT NULL,
    "name" varchar NOT NULL
);