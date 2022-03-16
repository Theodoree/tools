package _700_1799

/*
SELECT
	SUM(b.apple_count + IFNULL(c.apple_count,0)) AS apple_count,
	SUM(b.orange_count + IFNULL(c.orange_count,0)) AS orange_count
FROM
	Boxes  AS b
LEFT JOIN
	Chests AS c ON (c.chest_id = b.chest_id)

*/
