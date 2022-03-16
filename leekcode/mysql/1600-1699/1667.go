package _600_1699

/*
select
    user_id,
    concat(upper(left(name,1)),lower(substr(name,2))) as name
from
    Users
order by
    user_id
*/
