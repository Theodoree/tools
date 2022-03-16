package mysql

/*
SELECT
	s.id,
	s.name
FROM
	Students AS s
WHERE  NOT EXISTS
	(SELECT 1 FROM Departments AS d WHERE d.id = s.department_id)

*/
