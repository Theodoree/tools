package _100_1199

/*
SELECT
	e.student_id,
    e.course_id,
	e.grade
FROM (
SELECT
	student_id,
    course_id,
	grade,
	ROW_NUMBER()  OVER(PARTITION BY student_id ORDER BY grade DESC,course_id) AS rk
FROM
	Enrollments
	) AS e
WHERE e.rk = 1
*/
