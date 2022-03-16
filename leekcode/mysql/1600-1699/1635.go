package _600_1699

/*
with recursive t(n) as (
    SELECT 1
    UNION ALL
    SELECT n+1
    FROM t WHERE n<12
)

SELECT d.month, active_drivers, ifnull(accepted_rides,0) as accepted_rides
FROM
    (SELECT n month, COUNT(driver_id) active_drivers
    FROM t
    LEFT JOIN Drivers ON 202000+n >=Date_format(join_date,'%Y%m')
    GROUP BY month) d
LEFT JOIN
    (SELECT month(requested_at) month, COUNT(distinct ride_id) accepted_rides
    FROM Rides
    WHERE ride_id IN (SELECT ride_id FROM AcceptedRides)
    AND requested_at BETWEEN '2020-01-01' AND '2020-12-31'
    GROUP BY 1) r
ON d.month = r.month
*/
