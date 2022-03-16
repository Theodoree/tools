package _00_199

/*
183. 从不订购的客户

SELECT
	c.Name
FROM
	Customers AS c
WHERE NOT EXISTS (SELECT 1 FROM Orders AS o WHERE o.CustomerId = c.id)
*/
