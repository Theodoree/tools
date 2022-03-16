package mysql

/*
SELECT
	a1.sell_date,
	COUNT(DISTINCT product) AS num_sold,
	group_concat(DISTINCT product ORDER BY product) AS products
FROM
	Activities  AS a1
GROUP BY a1.sell_date
*/
