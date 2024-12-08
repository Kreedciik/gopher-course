-- Retrive all orders of specific customer

SELECT c.name, p.name, p.category, p.price, o.total_amount FROM orders o
INNER JOIN customers c ON c.customer_id = o.customer_id
INNER JOIN products p ON o.product_id = p.product_id
WHERE c.name = 'John'
ORDER BY p.price DESC;
