package _00_599

/*
SELECT
	c.Name AS name
FROM
	Candidate  AS c
INNER JOIN
	Vote AS v ON (v.CandidateId = c.id)
GROUP BY
	c.Name
HAVING
	COUNT(1) >=  ALL(SELECT COUNT(*) FROM  Vote GROUP BY CandidateId)
*/
