package _00_199

/*
. 连续出现的数字


三表对比

select distinct L1.Num ConsecutiveNums from Logs L1,Logs L2,Logs L3
where L1.id=L2.id+1 and L2.id+1=L3.id+2 and L1.num=L2.num and L2.num=L3.num;

用户变量
select distinct Num as ConsecutiveNums from
(select
   Num,  //这里使用的还是一个用户变量,按照主键排序,如果当前Num等于上一个num cnt++
	@cnt:=if(@pre=num,@cnt:=@cnt+1,@cnt:=1) as n,
 @pre:=Num from Logs,
 (select @pre := -1,@cnt := 0)as init
)as t
where t.n>=3;

*/
