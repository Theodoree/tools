package mysql

/*
WITH tmp_table AS (SELECT
	order_id,
	order_date,
	product_id,
	RANK() OVER(PARTITION BY product_id ORDER BY order_date DESC) AS rk
FROM
	Orders
)
SELECT
	p.product_name,
	p.product_id,
	t.order_id,
	t.order_date
FROM
	Products AS p
INNER JOIN
	tmp_table AS t  ON (p.product_id = t.product_id AND t.rk = 1)
ORDER BY product_name,order_id
*/
