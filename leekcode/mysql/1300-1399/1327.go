package mysql

/*
SELECT
	p.product_name,
	SUM(o.unit) AS 'unit'
FROM
	Orders AS o
INNER JOIN
	Products AS p ON  p.product_id = o.product_id
WHERE
	o.order_date >= '2020-02-01' AND o.order_date < '2020-03-01'
GROUP BY
	o.product_id
HAVING
	unit >= 100
*/
