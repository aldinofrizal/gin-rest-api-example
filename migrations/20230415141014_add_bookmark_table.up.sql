CREATE TABLE IF NOT EXISTS bookmarks (
  "id" SERIAL PRIMARY KEY,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz,
  "tmdb_id" int NOT NULL,
  "name" varchar NOT NULL,
  "overview" text NOT NULL,
  "poster_path" varchar NOT NULL,
  "user_id" int NOT NULL,
  FOREIGN KEY ("user_id")
    REFERENCES users ("id")
)