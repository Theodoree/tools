package _700_1799

/*
select user_id, max(windows) as biggest_window
from (select user_id,
      datediff(lead(visit_date, 1, "2021-01-01") over (partition       by user_id order by visit_date asc), visit_date) as      windows
from UserVisits) as temp
group by user_id;
*/
