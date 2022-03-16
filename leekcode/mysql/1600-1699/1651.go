package _600_1699

/*
with recursive tmp(n) as (
    select 1
    union all
    select n+1 from tmp where n < 12
)
select
month,
round(avg(ride_distance) over(order by month rows between current row and 2 following), 2) average_ride_distance,
round(avg(ride_duration) over(order by month rows between current row and 2 following), 2) average_ride_duration
from
(
    select
    t.n month,
    ifnull(sum(ride_distance), 0) ride_distance,
    ifnull(sum(ride_duration), 0) ride_duration
    from
    acceptedrides a join rides r on a.ride_id = r.ride_id and year(requested_at) = '2020'
    right join tmp t on t.n = month(requested_at)
    group by t.n
) as t limit 10
*/
