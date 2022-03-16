package _100_1199

/*
WITH tmp_table AS (
SELECT
	event_type,
	avg(occurences) AS occurences
FROM
	Events
GROUP BY event_type
)

SELECT
	 e.business_id
FROM
	Events  AS e
INNER JOIN tmp_table AS t  ON (t.event_type = e.event_type AND t.occurences < e.occurences)
GROUP BY e.business_id
HAVING COUNT(1) > 1
*/
