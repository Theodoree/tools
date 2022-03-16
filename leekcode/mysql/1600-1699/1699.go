package _600_1699

/*
SELECT
    IF(from_id<to_id,from_id,to_id) person1,
    IF(from_id>to_id,from_id,to_id) person2,
    COUNT(*) call_count,
    SUM(duration) total_duration
FROM
    Calls
GROUP BY
    IF(from_id<to_id,from_id,to_id),
    IF(from_id>to_id,from_id,to_id)
*/
