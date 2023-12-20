-- Table: public.Orders

-- DROP TABLE IF EXISTS public."Orders";

CREATE TABLE IF NOT EXISTS public."Orders"
(
    id integer NOT NULL,
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