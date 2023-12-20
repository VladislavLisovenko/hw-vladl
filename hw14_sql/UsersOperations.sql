delete from public."Users";

INSERT INTO public."Users"(
	id, name, email, password)
	VALUES (1, 'Ivanov', 'ivanov@mail.ru', '012'),
	(2, 'Petrov', 'petrov@mail.ru', '123'),
	(3, 'Sidorov', 'sidorov@mail.ru', '234'),
	(4, 'Smirnov', 'smirnov@mail.ru', '345'),
	(5, 'Kulikov', 'kulikov@mail.ru', '456'),
	(6, 'Vetrov', 'vetrov@mail.ru', '567');
	
update public."Users" set email = replace(email, 'mail.ru', 'gmail.com');

delete from public."Users" where name in ('Petrov', 'Sidorov');

SELECT id, name, email, password
	FROM public."Users";