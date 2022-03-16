package _600_1699

/*
select seller_name
from Seller
where seller_id not in (select seller_id from Orders where year(sale_date)=2020)
order by seller_name;
*/
