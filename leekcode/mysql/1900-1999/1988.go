package _900_1999

/*

WITH tmp_table AS ( SELECT
	s.school_id,
	e.score
FROM
	Schools AS s
LEFT JOIN
	Exam AS e on s.capacity >= e.student_count
)

SELECT
	s.school_id,
	IFNULL(MIN(t.score),-1) AS score
FROM
	Schools AS s
LEFT JOIN
	tmp_table AS t ON (t.school_id = s.school_id)
GROUP BY s.school_id
*/
