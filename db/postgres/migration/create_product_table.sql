DROP TABLE IF EXISTS "public"."products";
CREATE SEQUENCE IF NOT EXISTS products_id_seq
CREATE TABLE "public"."products" (
    "id" int4 NOT NULL DEFAULT nextval('products_id_seq'::regclass),
    "name" varchar,
    "description" varchar,
    "price" float8,
    "quantity" int8,
    "created_at" timestamp,
    "updated_at" timestamp,
    "deleted_at" timestamp
);