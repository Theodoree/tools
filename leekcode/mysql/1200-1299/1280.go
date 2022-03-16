package mysql

/*
select stu.student_id, stu.student_name, sub.subject_name, count(e.student_id) as attended_exams
from Subjects sub join Students  stu left join Examinations e
on e.subject_name = sub.subject_name and e.student_id = stu.student_id
group by sub.subject_name, stu.student_id
order by stu.student_id, sub.subject_name;
*/
