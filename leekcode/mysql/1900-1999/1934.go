package _900_1999

/*
select a.user_id, round(avg(if(action='confirmed',1,0)),2) confirmation_rate
from Signups a left join Confirmations b on a.user_id = b.user_id
group by 1
*/
