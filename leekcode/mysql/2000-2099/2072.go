package mysql

/*
select
    case when c1 > c2 then 'New York University'
         when c2 > c1 then 'California University'
         else 'No Winner' end as Winner
from
    (select sum(score>=90) c1 from NewYork) t1
    , (select sum(score>=90) c2 from California) t2
*/
