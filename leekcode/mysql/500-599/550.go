package _00_599

/*
SELECT ROUND(COUNT(t.player_id) / (SELECT COUNT(DISTINCT player_id) FROM Activity),2) AS fraction
FROM (
    SELECT DISTINCT player_id, MIN(event_date) AS first_date
    FROM Activity
    GROUP BY player_id
) AS t, Activity a
WHERE t.player_id=a.player_id AND DATEDIFF(a.event_date,t.first_date)=1
DATEDIFF函数等同于 ABS(date1 - date2)

*/
