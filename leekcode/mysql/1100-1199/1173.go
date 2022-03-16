package _100_1199

/*
SELECT
	IFNULL(
		cast(SUM(IF(order_date=customer_pref_delivery_date,1,0))/SUM(1)*100 AS DECIMAL(10,2))
		,0.0)  AS immediate_percentage
FROM
	Delivery
*/
