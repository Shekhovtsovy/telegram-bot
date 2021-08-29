CREATE TABLE IF NOT EXISTS "messages" (
    "id" bigint PRIMARY KEY,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "user_id" bigint,
    "chat_id" bigint,
    "text" text
);
