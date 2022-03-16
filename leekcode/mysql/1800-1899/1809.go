package _800_1899

/*
select
    session_id
from
    Playback p left join Ads a
on
    p.customer_id = a.customer_id
group by
    1
having
    sum(case when a.timestamp between p.start_time and p.end_time then 0 else 1 end) = count(1);
*/
