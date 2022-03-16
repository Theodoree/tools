package mysql

/*
select customer_id,name
from
(select d.customer_id,d.name,c.month
from
(select a.customer_id,left(order_date,7) month ,sum(quantity*price)  cost
from Orders a join Product b on(a.product_id=b.product_id)
where left(order_date,7)='2020-06' || left(order_date,7)='2020-07'
group by customer_id,price  ,left(order_date,7)) c
join Customers d on(c.customer_id=d.customer_id)

group by d.customer_id,d.name,c.month
having sum(cost) >= 100) e
group by customer_id,name
having count(*) = 2
*/
