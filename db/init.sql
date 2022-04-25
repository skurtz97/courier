CREATE TABLE "posts" (
    "id" bigserial PRIMARY KEY,
    "title" text NOT NULL,
    "body" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "name" text NOT NULL,
    "email" text NOT NULL,
    "password" text NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE INDEX ON "posts" ("user_id");

CREATE INDEX ON "users" ("id");

