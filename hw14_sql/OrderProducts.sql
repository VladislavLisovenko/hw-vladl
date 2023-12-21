CREATE TABLE IF NOT EXISTS public."OrderProducts"
(
    id int,
    order_id int NOT NULL,
    product_id int NOT NULL,
    CONSTRAINT "OrderProducts_pkey" PRIMARY KEY (id),
    CONSTRAINT fk_order_products_orders FOREIGN KEY (order_id)
        REFERENCES public."Orders" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT fk_order_products_products FOREIGN KEY (product_id)
        REFERENCES public."Products" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."OrderProducts"
    OWNER to postgres;