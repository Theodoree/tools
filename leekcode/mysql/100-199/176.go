package _00_199

/*
176. 第二高的薪水


SELECT
	IFNULL((
SELECT Salary
	FROM
    Employee
    GROUP BY Salary
ORDER BY Salary DESC
LIMIT 1 OFFSET 1),NULL) AS SecondHighestSalary


*/
