package _800_1899

/*
WITH tmp_table AS (
SELECT
	salary,
	RANK() OVER(ORDER BY salary) AS team_id
FROM
	Employees
GROUP BY salary
HAVING COUNT(1) > 1
)
SELECT e.employee_id,e.name,e.salary,t.team_id
FROM
	Employees  AS e
INNER JOIN
	tmp_table AS t ON (t.salary = e.salary)
ORDER BY team_id,employee_id
*/
