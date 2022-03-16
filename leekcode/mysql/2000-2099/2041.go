package mysql

/*
SELECT
	candidate_id
FROM
	Candidates AS c
INNER JOIN
	Rounds AS r ON (r.interview_id = c.interview_id)
WHERE
	c.years_of_exp >= 2
GROUP BY
	candidate_id
HAVING
	SUM(r.score) > 15
*/
