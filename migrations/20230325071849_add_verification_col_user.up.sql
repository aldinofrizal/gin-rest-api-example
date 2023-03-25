ALTER TABLE users
ADD "verification_code" varchar,
ADD "is_active" boolean NOT NULL DEFAULT('false');