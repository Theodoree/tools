package _800_1899

/*

SELECT order_id
FROM OrdersDetails
GROUP BY order_id
HAVING max(quantity) >  ALL (SELECT AVG(quantity) FROM OrdersDetails GROUP BY order_id)
*/
