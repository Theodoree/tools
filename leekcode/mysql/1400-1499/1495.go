package mysql

/*
SELECT
	DISTINCT c.title
FROM
	TVProgram AS tv
INNER JOIN
	Content AS c ON (c.content_id = tv.content_id)
WHERE
	MONTH(tv.program_date) = 6 AND c.Kids_content = "Y" AND c.content_type = "Movies"

*/
