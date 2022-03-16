package _00_699

/*
SELECT
	t.id,
	t.num
FROM (
SELECT
	requester_id AS id,
	COUNT(accepter_id) AS num,
	ROW_NUMBER() OVER(ORDER BY COUNT(accepter_id) DESC ) AS rk
FROM (
SELECT
	requester_id,
	accepter_id
FROM
	RequestAccepted
UNION ALL
SELECT
	accepter_id,
	requester_id
FROM
	RequestAccepted
) AS t
GROUP BY requester_id
)  AS t
WHERE t.rk = 1
*/
