package mysql

/*
select
    distinct e.student_id,s.student_name
from
    exam e left join student s on e.student_id=s.student_id
where
    e.student_id not in
    (select student_id from exam where (exam_id,score) in((select exam_id,max(score) from exam group by exam_id)union all(select exam_id,min(score) from exam group by exam_id)))
order by
    e.student_id
*/
