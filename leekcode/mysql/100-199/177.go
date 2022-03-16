package _00_199

/*
177. 第N高的薪水


CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
SET N = N - 1;

  RETURN (

SELECT
	IFNULL((
SELECT Salary
	FROM
    Employee
    GROUP BY Salary
ORDER BY Salary DESC
LIMIT 1 OFFSET N),NULL) AS SecondHighestSalary
  );
END

*/
