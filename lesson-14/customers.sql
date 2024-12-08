CREATE TABLE customers
(
	customer_id SERIAL PRIMARY KEY,
	name VARCHAR,
	email VARCHAR,
	city VARCHAR,
	created_at TIMESTAMP
);

INSERT INTO customers(name, email, city) VALUES
('John', 'john@mail.ru', 'Frankfurt'),
('Anna', 'anna@mail.ru', 'Vienna'),
('Marton', 'm@mail.ru', 'Budapest'),
('Gabriel', 'g@mail.ru', 'Brasil'),
('Max', 'm@mail.ru', 'Moscow'),
('Sarah', 's@mail.ru', 'Colorado'),
('Michael', 'mich@mail.ru', 'New York'),
('Robert', 'rob@mail.ru', 'London'),
('Tiffany', 'tif@mail.ru', 'Stockholm'),
('Bob', 'bob@mail.ru', 'Warsaw');

SELECT * FROM customers;