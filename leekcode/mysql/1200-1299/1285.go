package mysql

/*
select start_id, min(end_id) end_id
from
(select log_id start_id from logs where log_id-1  not in (select * from logs))a,
(select log_id end_id from logs where log_id + 1 not in (select * from logs))b
where start_id <= end_id
group by start_id
*/
