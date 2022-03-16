package mysql

/*
SELECT '[0-5>' bin, COUNT(*) total
FROM Sessions
WHERE duration BETWEEN 0 AND 5*60
UNION ALL
SELECT '[5-10>' bin, COUNT(*) total
FROM Sessions
WHERE duration BETWEEN 5*60 AND 10*60
UNION ALL
SELECT '[10-15>' bin, COUNT(*) total
FROM Sessions
WHERE duration BETWEEN 10*60 AND 15*60
UNION ALL
SELECT '15 or more' bin, COUNT(*) total
FROM Sessions
WHERE duration >= 15*60
*/
