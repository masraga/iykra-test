DROP TABLE IF EXISTS "employees";
DROP SEQUENCE IF EXISTS employees_id_seq;
CREATE SEQUENCE employees_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."employees" (
    "id" integer DEFAULT nextval('employees_id_seq') NOT NULL,
    "name" character varying(100) NOT NULL,
    "position" character varying(70) NOT NULL,
    "salary" integer NOT NULL,
    CONSTRAINT "employees_pkey" PRIMARY KEY ("id")
)
WITH (oids = false);