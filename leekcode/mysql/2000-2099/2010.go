package mysql

/*
with s as(
    select employee_id, 70000 - sum(salary) over(order by salary) cum_sum
    from Candidates
    where experience = 'Senior'
),
     j as(
    select employee_id, ifnull((select min(cum_sum) from s where cum_sum >= 0),70000)
            - sum(salary) over(order by salary) cum_sum
    from Candidates
    where experience = 'Junior'
)
select employee_id from s where cum_sum >= 0
union all
select employee_id from j where cum_sum >= 0
*/
