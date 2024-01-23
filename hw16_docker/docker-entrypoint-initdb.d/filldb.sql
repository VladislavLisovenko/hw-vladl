-- Database: Trade

-- CREATE DATABASE "trade"
--     WITH
--     OWNER = postgres
--     ENCODING = 'UTF8'
--     LC_COLLATE = 'en_US.utf8'
--     LC_CTYPE = 'en_US.utf8'
--     LOCALE_PROVIDER = 'libc'
--     TABLESPACE = pg_default
--     CONNECTION LIMIT = -1
--     IS_TEMPLATE = False;

CREATE TABLE IF NOT EXISTS public."Users"
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    email character varying(100) COLLATE pg_catalog."default",
    password character varying(50) COLLATE pg_catalog."default",
    CONSTRAINT "Users_pkey" PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."Users"
    OWNER to postgres;


CREATE TABLE IF NOT EXISTS public."Products"
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    name character varying(100) COLLATE pg_catalog."default" NOT NULL DEFAULT ''::character varying,
    price numeric(15,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT "Products_pkey" PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."Products"
    OWNER to postgres;

CREATE TABLE IF NOT EXISTS public."Orders"
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    user_id integer NOT NULL,
    order_date date NOT NULL,
    total_amount numeric(15,2) NOT NULL DEFAULT 0,
    CONSTRAINT "Orders_pkey" PRIMARY KEY (id),
    CONSTRAINT fk_orders_users FOREIGN KEY (user_id)
        REFERENCES public."Users" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."Orders"
    OWNER to postgres;
-- Index: idx_user_id

-- DROP INDEX IF EXISTS public.idx_user_id;

CREATE INDEX IF NOT EXISTS idx_user_id
    ON public."Orders" USING btree
    (user_id ASC NULLS LAST)
    WITH (deduplicate_items=True)
    TABLESPACE pg_default;

CREATE TABLE IF NOT EXISTS public."OrderProducts"
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    order_id integer NOT NULL,
    product_id integer NOT NULL,
    CONSTRAINT "OrderProducts_pkey" PRIMARY KEY (id),
    CONSTRAINT fk_order_products_orders FOREIGN KEY (order_id)
        REFERENCES public."Orders" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_order_products_products FOREIGN KEY (product_id)
        REFERENCES public."Products" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."OrderProducts"
    OWNER to postgres;

INSERT INTO public."Users"(
	name, email, password)
	VALUES ('Dmitriy', 'dm@dm.dm', 'sec1'),
	('Andrey', 'an@an.an', 'sec2'),
	('Nikolay', 'ni@ni.ni', 'sec3');
	
INSERT INTO public."Products"(
	name, price)
	VALUES ('Computer', 10000.00),
	('Chair', 2000.00);
	
INSERT INTO public."Orders"(
	user_id, order_date, total_amount)
	VALUES (1, '2024-01-20', 12000),
	(2, '2024-01-21', 10000),
	(3, '2024-01-22', 2000);
	
INSERT INTO public."OrderProducts"(
	order_id, product_id)
	VALUES (1, 1),
	(1, 2),
	(2, 1),
	(3, 2);