package mysql

/*
select lower(replace(product_name, ' ', '')) as product_name, date_format(sale_date, '%Y-%m') as sale_date, count(*) as total
from sales
group by lower(replace(product_name, ' ', '')), date_format(sale_date, '%Y-%m')
order by 1, 2
*/
