-- Columns of table
CREATE TABLE employee
(
    id SERIAL PRIMARY KEY,
    name TEXT,
    department TEXT,
    salary FLOAT,
    hired_date TIMESTAMP
);

-- Jadvalni to'ldirish
INSERT INTO employee (name, department, salary, hired_date) VALUES
('Nina', 'Marketing', 92638.57, '2024-04-13 17:40:11'),
('Ian', 'Marketing', 74131.28, '2023-06-27 17:40:11'),
('Bob', 'HR', 73315.58, '2023-02-05 17:40:11'),
('Alice', 'Finance', 84161.96, '2024-02-16 17:40:11'),
('Bob', 'IT', 94653.14, '2024-03-19 17:40:11'),
('Bob', 'HR', 105624.72, '2023-02-26 17:40:11'),
('Hannah', 'HR', 75796.31, '2023-10-19 17:40:11'),
('Charlie', 'HR', 90061.51, '2024-06-26 17:40:11'),
('Kevin', 'IT', 79778.05, '2022-04-26 17:40:11');


-- Ish haqi $1000 dan kam bo‘lgan xodimlarning ro‘yxatini oling.
SELECT * FROM employee
WHERE salary < 1000;

-- Department bo‘yicha guruhlang va har bir bo‘limda 
-- qancha xodim borligini toping.
SELECT department, COUNT(id) FROM employee
GROUP BY department
ORDER BY COUNT(id) DESC;

-- Ishga qabul qilingan eng so‘nggi xodimning ma’lumotlarini ko‘rsating.
SELECT * FROM employee
ORDER BY hired_date DESC
LIMIT 1;