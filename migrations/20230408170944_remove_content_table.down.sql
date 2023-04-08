CREATE TABLE IF NOT EXISTS contents (
    "id" SERIAL PRIMARY KEY,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    "deleted_at" timestamptz,
    "name" varchar NOT NULL,
    "description" text NOT NULL,
    "image_url" varchar NOT NULL,
    "author_id" int NOT NULL,
    FOREIGN KEY ("author_id")
      REFERENCES users ("id")
);