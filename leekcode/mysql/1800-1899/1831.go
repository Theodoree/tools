package _800_1899

/*
with tmp_table AS (
SELECT
	DATE_FORMAT(day,'%Y-%m-%d')  AS d,
	RANK() OVER(partition by DATE_FORMAT(day,'%Y-%m-%d') ORDER BY amount DESC )  AS r,
	transaction_id
FROM
	Transactions
ORDER BY transaction_id
)
SELECT tmp_table.transaction_id FROM tmp_table WHERE tmp_table.r = 1;
*/
