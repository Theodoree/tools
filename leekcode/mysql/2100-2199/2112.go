package _100_2199

/*
WITH t AS
(SELECT airport_id,
        DENSE_RANK()OVER(ORDER BY SUM(flights_count) DESC) AS rk
FROM
(SELECT departure_airport AS airport_id,
        flights_count
FROM Flights

UNION ALL

SELECT arrival_airport AS airport_id,
       flights_count
FROM Flights)temp

GROUP BY airport_id
)

SELECT airport_id
FROM t
WHERE rk = 1
*/
