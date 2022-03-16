package _00_599

/*
SELECT
	e1.name
FROM
    Employee  AS e
INNER JOIN
	Employee AS e1 ON (e1.id = e.managerId)
GROUP BY  e.managerId
HAVING COUNT(e1.id) >=5
*/
