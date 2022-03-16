package _900_1999

/*
with t1 as
(select 'Android' platform
 union
 select 'IOS' platform
 union
 select 'Web' platform),
t2 as
(select 'Reading' experiment_name
 union
 select 'Sports' experiment_name
 union
 select 'Programming' experiment_name)

select t1.platform, t2.experiment_name, count(experiment_id) num_experiments
from t1
join t2
left join experiments e on t1.platform = e.platform and t2.experiment_name = e.experiment_name
group by t1.platform, t2.experiment_name
order by 1
*/
