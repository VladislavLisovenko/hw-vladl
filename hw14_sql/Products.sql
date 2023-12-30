CREATE TABLE IF NOT EXISTS public."Products"
(
    id int,
    name character varying(100) COLLATE pg_catalog."default" NOT NULL DEFAULT ''::character varying,
    price numeric(15,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT "Products_pkey" PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."Products"
    OWNER to postgres;