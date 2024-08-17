CREATE TABLE "accounts" (
                            "id" bigserial PRIMARY KEY,
                            "owner" varchar NOT NULL,
                            "balance" bigint NOT NULL,
                            "currency" varchar NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE "entries" (
                           "id" bigserial PRIMARY KEY,
                           "account_id" bigint NOT NULL ,
                           "amount" bigint NOT NULL,
                           "created_at" timestamptz DEFAULT now(),
                           CONSTRAINT fk_account FOREIGN KEY ("account_id") REFERENCES "accounts" ("id")
);

CREATE TABLE "transfers" (
                             "id" bigserial PRIMARY KEY,
                             "from_account_id" bigint NOT NULL,
                             "to_account_id" bigint NOT NULL,
                             "amount" bigint NOT NULL,
                             "created_at" timestamptz NOT NULL DEFAULT now(),
                             CONSTRAINT fk_from_account FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id"),
                             CONSTRAINT fk_to_account FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id")
);

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "entries"."amount" IS 'it can be negative or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';

