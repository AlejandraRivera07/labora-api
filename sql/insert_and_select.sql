INSERT INTO items (customer_name, order_date, product, quantity, price)
VALUES ('Lucas Garcia', '2023-05-02', 'Manzana', 50, 2500),
       ('Javier Mendez', '2023-05-04', 'Tomate', 12, 800),
	   ('Carmenza Cortez', '2023-05-07', 'Aguacate', 23, 347),
	   ('Martin Mendoza', '2023-04-24', 'Banano', 2, 67),
	   ('Elizabeth Ramirez', '2023-04-12', 'Naranja', 1, 20),
	   ('Laura Diaz', '2023-05-08', 'fresa', 1, 35);


SELECT *
FROM items
WHERE quantity >2 
AND price >50;