package _000_1099

/*
select user_id buyer_id, join_date,
sum(if(year(order_date)='2019',1,0)) orders_in_2019
from users left join orders
on user_id = buyer_id
group by user_id
*/
