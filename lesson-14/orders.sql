DROP TABLE orders;
SELECT * FROM orders;
CREATE TABLE orders
(
	order_id SERIAL PRIMARY KEY,
	customer_id INT,
	product_id INT,
	order_date DATE,
	total_amount INT,
	CONSTRAINT FK_customers_customer_id FOREIGN KEY(customer_id) REFERENCES customers(customer_id),
	CONSTRAINT FK_products_product_id FOREIGN KEY(product_id) REFERENCES products(product_id)
);

INSERT INTO orders (customer_id, order_date, total_amount, product_id) VALUES
(1, '2024-11-02', 1, 1),
(1, '2024-11-02', 1, 2),
(1, '2024-11-02', 2, 3),
(2, '2024-12-02', 2, 9),
(3, '2024-12-04', 1, 4),
(3, '2024-12-04', 2, 8),
(3, '2024-12-05', 1, 6),
(4, '2024-12-06', 5, 5),
(5, '2024-12-07', 1, 8),
(5, '2024-12-07', 1, 7),
(6, '2024-11-28', 1, 10),
(6, '2024-11-28', 1, 1),
(7, '2024-11-20', 8, 4),
(8, '2024-11-11', 1, 3),
(8, '2024-11-11', 2, 8),
(8, '2024-11-10', 3, 2),
(9, '2024-10-10', 5, 7),
(10, '2024-09-10', 3, 1),
(10, '2024-09-10', 2, 9),
(10, '2024-09-10', 1, 10);

TRUNCATE orders;