DROP TABLE products;
SELECT * FROM products;
CREATE TABLE products 
(
	product_id SERIAL PRIMARY KEY,
	name VARCHAR,
	category VARCHAR,
	price FLOAT
);

INSERT INTO products (name, category, price) VALUES
('Wireless Mouse', 'Electronics', 25.99),
('Bluetooth Keyboard', 'Electronics', 45.50),
('LED Desk Lamp', 'Home Appliances', 30.00),
('Ceramic Vase', 'Home Decor', 15.75),
('Stainless Steel Pan', 'Kitchenware', 40.99),
('Yoga Mat', 'Sports Equipment', 22.50),
('Noise Cancelling Headphones', 'Electronics', 99.99),
('Hardcover Journal', 'Stationery', 12.99),
('Travel Backpack', 'Accessories', 59.99),
('Running Shoes', 'Sports Equipment', 75.00);