CREATE TABLE "users" (
  "username" varchar UNIQUE NOT NULL,
  "hashed_password" varchar,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("username", "hashed_password")
);

ALTER TABLE "account" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");



