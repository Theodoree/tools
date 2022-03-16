package mysql

/*
select person_name
from
(
    select person_name,turn,sum(weight)over(order by turn) as addup_weight
    from Queue a
) t where addup_weight<=1000
order by turn desc
LIMIT 1
*/
