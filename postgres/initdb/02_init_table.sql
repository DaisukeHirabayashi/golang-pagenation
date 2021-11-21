\c go-pagenation;

CREATE TABLE IF NOT EXISTS shops (
    id SERIAL NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

INSERT INTO "public"."shops" ("name") VALUES ('Pretty');
INSERT INTO "public"."shops" ("name") VALUES ('Beauty-Beauty');
INSERT INTO "public"."shops" ("name") VALUES ('Beauty-Beauty');
INSERT INTO "public"."shops" ("name") VALUES ('Beauty-Beauty');
INSERT INTO "public"."shops" ("name") VALUES ('Beauty-Beauty');
INSERT INTO "public"."shops" ("name") VALUES ('Beauty-Beauty');
INSERT INTO "public"."shops" ("name") VALUES ('Beauty-Beauty');