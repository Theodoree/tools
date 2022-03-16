package mysql

/*
SELECT
	gender,
	day,
    SUM(score_points) over (partition by gender order by day ) AS  total
FROM
	Scores
ORDER BY gender,day
*/
