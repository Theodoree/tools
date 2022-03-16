package mysql

/*
with a as
(select
    customer_id,
    product_id,
    rank() over(partition by customer_id order by count(product_id) desc)r
from orders
group by customer_id, product_id
)
select customer_id, a.product_id, product_name from a
join products b on a.product_id = b.product_id
where r = 1
*/
