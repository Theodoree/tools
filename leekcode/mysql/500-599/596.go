package _00_599

/*
596. 超过5名学生的课
SELECT
 	tmp.class
FROM
(
SELECT
	COUNT(DISTINCT student) AS cnt,
	class
FROM
	courses
GROUP BY class
HAVING cnt > 4
) AS tmp
*/
