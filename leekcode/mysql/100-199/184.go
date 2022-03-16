package _00_199

/* 184 部门工资最高的员工
SELECT
    d.name AS 'Department',
    e.name AS 'Employee',
    e.Salary
FROM
    Employee AS e
INNER JOIN
	Department AS d ON (e.DepartmentId = d.Id)
where
    (e.DepartmentId, e.Salary) IN
    (
        SELECT
            DepartmentId,  MAX(Salary)
        FROM
            Employee
        GROUP BY
			DepartmentId
    )
*/
