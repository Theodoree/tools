package mysql

/*
WITH tmp_table AS (
SELECT
	company_id,
  CASE
		when  max(salary) > 10000 then 1-0.49
		when  max(salary) > 1000 then 1-0.24
		else 1
	end AS val

FROM
	Salaries
GROUP BY company_id
)


SELECT
	s.company_id,
	s.employee_id,
	s.employee_name,
	CAST((s.salary *t.val) AS DECIMAL(10,0)) AS salary
FROM
	Salaries AS s
INNER JOIN
	tmp_table AS t ON (t.company_id =s.company_id)

*/
