package mysql

/*
SELECT
	query_name,
	IFNULL(cast(SUM(rating/position)/COUNT(1) AS DECIMAL(10,2)),0.0) AS quality,
	IFNULL(cast(SUM(IF(rating < 3,1,0))/COUNT(1)*100 AS DECIMAL(10,2)),0.0) AS poor_query_percentage
FROM
	Queries
GROUP BY
	query_name
*/
