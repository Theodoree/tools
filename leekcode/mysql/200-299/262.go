package _00_299

/*
262. 行程和用户

SELECT
    Request_at AS 'Day',
	 round(COUNT(IF(t.Status != 'completed',1,NULL))/COUNT(1),2) AS 'Cancellation Rate'
FROM
	Users AS  u
LEFT JOIN
	Trips AS t ON (u.Users_Id = t.Client_id AND u.Banned = 'No')
WHERE Request_at BETWEEN '2013-10-01' AND  '2013-10-03'
GROUP BY Request_at




*/
