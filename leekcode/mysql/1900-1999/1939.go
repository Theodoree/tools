package _900_1999

/*
select
    a.user_id
from
    Confirmations a left join Confirmations b
    on TIMESTAMPDIFF(second,a.time_stamp,b.time_stamp) between 0 and 86400
    and a.user_id = b.user_id and a.time_stamp != b.time_stamp
group by
    1
having
    count(b.user_id) > 0
*/
