package _600_1699

/*
SELECT
	r.contest_id,
	cast(COUNT(DISTINCT r.user_id)/(SELECT COUNT(DISTINCT user_id) FROM Users)*100 AS DECIMAL(10,2)) AS percentage
FROM
	Register AS r
GROUP BY r.contest_id
ORDER BY percentage DESC,contest_id
*/
