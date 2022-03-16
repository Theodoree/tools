package mysql

/*
WITH tmp_table AS (
SELECT
 username,
 activity,
 startDate,
 endDate,
 ROW_NUMBER() OVER(PARTITION BY username ORDER BY startDate DESC) AS rk
FROM
	UserActivity
)

SELECT
 DISTINCT u.username,
 IFNULL(t.activity,t1.activity) AS activity,
 IFNULL(t.startDate,t1.startDate) AS startDate,
 IFNULL(t.endDate,t1.endDate) AS endDate
FROM
	UserActivity AS u
LEFT JOIN
	tmp_table AS t ON (t.rk = 2 AND t.username = u.username)
LEFT JOIN
	tmp_table AS t1 ON (t1.rk = 1 AND t1.username = u.username)

*/
