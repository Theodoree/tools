package _800_1899

/*
select u.name,u.mail from contests c join users u on c.gold_medal = u.user_id
group by u.name,u.mail  having count(*) > 2
union
select name,mail
from
(
select
u.name,u.mail,
contest_id-row_number() over(partition by user_id order by contest_id) flag
from
users u join contests c
on user_id in (gold_medal, silver_medal, bronze_medal)
) t
group by name,mail, flag having count(*) > 2

*/
