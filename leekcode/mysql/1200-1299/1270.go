package mysql

/*
select a.employee_id
from Employees as a
left join
Employees as b
on a.manager_id = b.employee_id
left join
Employees as c
on b.manager_id = c.employee_id
where c.manager_id = 1 and a.employee_id <> 1

SELECT
	e.employee_id
FROM
	Employees AS e
INNER JOIN
	Employees AS e1 ON (e1.employee_id = e.manager_id)
LEFT  JOIN
	Employees AS e2 ON (e2.employee_id = e1.manager_id)
LEFT  JOIN
	Employees AS e3 ON (e3.employee_id = e2.manager_id)
WHERE
	 e.employee_id != 1 AND (e3.employee_id = 1 OR 	e1.employee_id = 1 OR 	e2.employee_id = 1)
*/
