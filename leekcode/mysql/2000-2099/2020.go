package mysql

/*
select count(su.account_id) accounts_count
from subscriptions su
join streams st on  su.account_id=st.account_id
where year(end_date)=2021
and year(stream_date)<>2021
*/
