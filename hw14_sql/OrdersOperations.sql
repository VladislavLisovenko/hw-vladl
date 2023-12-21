delete from public."OrderProducts";
delete from public."Orders";

INSERT INTO public."Orders"(
	id, user_id, order_date, total_amount)
	VALUES (1, 1, '2023-12-01', 10000),
	(2, 4, '2023-12-04', 100000),
	(3, 5, '2023-12-05', 30000),
	(4, 6, '2023-12-06', 9000),
	(5, 1, now(), 30000),
	(6, 4, now(), 27000),
	(7, 5, now(), 30000),
	(8, 6, now(), 9000);

delete from public."Orders" where user_id = 1;

INSERT INTO public."OrderProducts"(
	id, order_id, product_id)
	VALUES (1, 2, 1),
	(2, 3, 2),
	(3, 4, 4),
	(4, 6, 5),
	(5, 7, 2),
	(6, 8, 4);

SELECT id, user_id, order_date, total_amount
	FROM public."Orders"
	WHERE user_id = 5;

SELECT u.name, COUNT(*) AS total_count, SUM(total_amount) AS total_amount
	FROM public."Orders" as o
	INNER JOIN public."Users" AS u
		ON u.id = o.user_id
	GROUP BY user_id, u.name;