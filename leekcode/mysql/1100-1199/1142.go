package _100_1199

/*
SELECT
	IFNULL(cast(COUNT(DISTINCT session_id)/COUNT(DISTINCT user_id) as decimal(10,2)),0.0)  AS average_sessions_per_user
FROM
	Activity
WHERE
	activity_date > DATE_SUB("2019-07-27",INTERVAL 30 DAY) AND activity_date <= "2019-07-27"

*/
