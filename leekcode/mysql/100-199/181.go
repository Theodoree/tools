package _00_199

/*
181. 超过经理收入的员工

SELECT
	e.Name AS Employee
FROM
	Employee AS e,
	Employee AS e1

WHERE  e.ManagerId = e1.id AND  e.Salary > e1.Salary
*/
