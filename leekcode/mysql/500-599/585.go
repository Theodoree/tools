package _00_599

/*
select round(sum(distinct i2.TIV_2016),2) tiv_2016
from insurance i1,insurance i2
where i2.PID!=i1.PID and i2.TIV_2015=i1.TIV_2015 and
(i2.LAT,i2.LON) not in
(select i.LAT,i.LON from insurance i where i.PID!=i2.PID)
*/
