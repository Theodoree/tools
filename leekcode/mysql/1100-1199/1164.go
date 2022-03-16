package _100_1199

/*
WITH tmp_tabel AS (
SELECT
	product_id,
	new_price,
	change_date,
	ROW_NUMBER() OVER(PARTITION BY product_id ORDER BY change_date DESC) AS rt
FROM
	Products
WHERE
	change_date <= "2019-08-16"
)
SELECT
	DISTINCT p.product_id,
	IFNULL(t.new_price,10) AS price
FROM
	Products AS  p
LEFT JOIN
	tmp_tabel AS t ON (t.product_id = p.product_id AND t.rt = 1)



*/
