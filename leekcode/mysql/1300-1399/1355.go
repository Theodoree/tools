package mysql

/*
SELECT activity
FROM Friends
GROUP BY activity
HAVING
COUNT(*)>SOME(SELECT COUNT(*) FROM Friends GROUP BY activity) // SOME = ANY = IN,必须要大于那么不是最后一条
AND
COUNT(*)<SOME(SELECT COUNT(*) FROM Friends GROUP BY activity) // 必须要小于,那么不是最大的一条
*/
