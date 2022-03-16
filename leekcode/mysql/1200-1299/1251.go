package mysql

/*
SELECT
  p.product_id,
	IFNULL(CAST(SUM(p.price*u.units)/SUM(units)  AS DECIMAL(10,2)),0.0) AS average_price
FROM
	UnitsSold AS u
INNER JOIN
	Prices AS p ON (p.product_id = u.product_id AND u.purchase_date BETWEEN p.start_date AND p.end_date )

GROUP BY
	product_id
*/
