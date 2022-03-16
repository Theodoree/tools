package mysql

/*
SELECT
	c.customer_id,
	c.customer_name
FROM (
SELECT
    c.customer_id,
    c.customer_name,
		COUNT(DISTINCT IF(o.product_name='A' OR o.product_name='B',o.product_name,null)) AS cnt,
		COUNT(DISTINCT IF(o.product_name='C',o.product_name,null)) AS cnt1
FROM
    Customers AS c
INNER JOIN
    Orders AS o ON (o.customer_id = c.customer_id)
GROUP BY c.customer_id, c.customer_name
) AS c
WHERE c.cnt =2 AND cnt1 = 0
*/
