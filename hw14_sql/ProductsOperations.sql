delete from public."Products";

INSERT INTO public."Products"(
	id, name, price)
	VALUES (1, 'Computer', 100000),
	(2, 'Byke', 30000),
	(3, 'Keyboard', 2000),
	(4, 'Speakers', 5000),
	(5, 'Phone', 15000);
	
update public."Products" set price = price * 1.8 where price < 20000;

delete from public."Products" where price < 5000;

SELECT id, name, price
	FROM public."Products";