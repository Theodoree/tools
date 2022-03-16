package _600_1699

/*
WITH RECURSIVE t1 AS
(
SELECT 1 AS n
UNION ALL
SELECT
	n+1
from
	t1
where
	n < 100
)
SELECT
	n ids
FROM
	t1
WHERE
	n < (SELECT MAX(customer_id) FROM Customers)
AND
	n NOT IN (SELECT customer_id FROM Customers)
*/
