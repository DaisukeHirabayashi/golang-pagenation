\c go-pagenation;

CREATE TABLE IF NOT EXISTS shop (
    id SERIAL NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

INSERT INTO "public"."shop" ("name") VALUES ('Pretty');
INSERT INTO "public"."shop" ("name") VALUES ('Beauty-Beauty');
INSERT INTO "public"."shop" ("name") VALUES ('Beauty-Beauty');
INSERT INTO "public"."shop" ("name") VALUES ('Beauty-Beauty');
INSERT INTO "public"."shop" ("name") VALUES ('Beauty-Beauty');
INSERT INTO "public"."shop" ("name") VALUES ('Beauty-Beauty');
INSERT INTO "public"."shop" ("name") VALUES ('Beauty-Beauty');