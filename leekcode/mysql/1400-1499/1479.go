package mysql

/*
select item_category category,
        sum(if(weekday(order_date)=0,quantity,0)) 'Monday',
        sum(if(weekday(order_date)=1,quantity,0)) 'Tuesday',
        sum(if(weekday(order_date)=2,quantity,0)) 'Wednesday',
        sum(if(weekday(order_date)=3,quantity,0)) 'Thursday',
        sum(if(weekday(order_date)=4,quantity,0)) 'Friday',
        sum(if(weekday(order_date)=5,quantity,0)) 'Saturday',
        sum(if(weekday(order_date)=6,quantity,0)) 'Sunday'
from orders o right join items i
on o.item_id = i.item_id
group by item_category
order by category
*/
