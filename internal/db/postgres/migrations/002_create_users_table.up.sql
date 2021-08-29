CREATE TABLE IF NOT EXISTS "users" (
    "id" bigint PRIMARY KEY,
    "first_name" varchar,
    "last_name" varchar NULL,
    "user_name" varchar NULL,
    "language_code" varchar NULL,
    "is_bot" boolean NOT NULL DEFAULT false,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);
