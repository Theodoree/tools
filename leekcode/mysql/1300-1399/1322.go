package mysql

/*
SELECT
	ad_id,
	cast(IFNULL(SUM(IF(action = "Clicked",1,0))/SUM(IF(action != "Ignored",1,0)),0.0)*100 AS DECIMAL(10,2)) AS 'ctr'
FROM
	Ads
GROUP BY
	ad_id
ORDER BY ctr DESC,ad_id
*/
