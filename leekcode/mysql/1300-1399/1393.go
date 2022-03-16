package mysql

/*
SELECT
	stock_name,
  SUM(IF(operation="Sell",price,0))-SUM(IF(operation="Buy",price,0)) AS capital_gain_loss
FROM
  Stocks

GROUP BY stock_name
*/
