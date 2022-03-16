package mysql

/*
WITH tmp_table AS (
SELECT
	c.name AS customer_name,
	o.customer_id,
	o.order_id,
	o.order_date,
	RANK() OVER(PARTITION BY customer_id ORDER BY order_date DESC) AS rk
FROM
	Orders AS o
INNER JOIN
	Customers AS c ON (c.customer_id = o.customer_id)
ORDER BY c.name,o.customer_id,o.order_date DESC

)
SELECT
	customer_name,
	customer_id,
	order_id,
	order_date
FROM
	tmp_table

WHERE
	rk <= 3
*/
