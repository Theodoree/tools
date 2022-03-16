package mysql

/*
# Write your MySQL query statement below
SELECT
    u.name,
    IFNULL(SUM(distance),0) AS travelled_distance
FROM
    Users AS u
LEFT JOIN
    Rides AS r  ON (r.user_id = u.id)
GROUP BY u.name
ORDER BY travelled_distance DESC,u.name
*/
