package mysql

/*
WITH tmp_table AS (
SELECT
	caller.caller_id,
	SUM(duration) AS duration,
	COUNT(1) AS cnt
FROM (
SELECT
	caller_id,
	duration
FROM
	Calls
UNION ALL
SELECT
	callee_id,
	duration
FROM
	Calls
) AS caller
GROUP BY caller.caller_id
)

SELECT
	c.name AS country
FROM
	Country AS c
INNER JOIN
	Person AS p ON (SUBSTR(p.phone_number,1,3) = c.country_code)
INNER JOIN
	tmp_table AS t ON (t.caller_id = p.id)
GROUP BY c.name
HAVING SUM(t.duration)/SUM(t.cnt) > (SELECT AVG(duration) FROM Calls)
*/
