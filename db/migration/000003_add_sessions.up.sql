CREATE TABLE "session" (
                            "id" bigserial PRIMARY KEY,
                            "username" varchar NOT NULL,
                            "refresh_token" varchar NOT NULL,
                            "user_agent" varchar NOT NULL,
                            "client_ip" varchar NOT NULL,
                            "is_blocked" bool NOT NULL,
                            "expires_at" timestamptz NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

ALTER TABLE "session" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");