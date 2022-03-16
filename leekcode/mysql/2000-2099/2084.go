package mysql

/*
SELECT
	order_id,
	customer_id,
	order_type
FROM
	Orders AS o
WHERE
	IF(customer_id in (SELECT DISTINCT customer_id FROM Orders WHERE order_type = 0),order_type=0,1 = 1)

*/
