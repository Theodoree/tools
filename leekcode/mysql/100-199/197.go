package _00_199

/* 上升的温度


SELECT
	w.id
FROM
	 Weather  AS w
INNER JOIN Weather AS w1 ON (w1.RecordDate = DATE_SUB(w.RecordDate,INTERVAL 1 DAY ) )
WHERE w1.Temperature < w.Temperature

*/
