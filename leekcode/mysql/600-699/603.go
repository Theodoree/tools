package _00_699

/*
603. 连续空余座位

几个朋友来到电影院的售票处，准备预约连续空余座位。

你能利用表 cinema ，帮他们写一个查询语句，获取所有空余座位，并将它们按照 seat_id 排序后返回吗？

| seat_id | free |
|---------|------|
| 1       | 1    |
| 2       | 0    |
| 3       | 1    |
| 4       | 1    |
| 5       | 1    |


对于如上样例，你的查询语句应该返回如下结果。



| seat_id |
|---------|
| 3       |
| 4       |
| 5       |
*/

/*
SELECT
	seat_id
FROM
	cinema  AS s
WHERE s.free = 1 AND EXISTS (SELECT 1 FROM cinema AS s1 WHERE s1.free = 1 AND (s1.seat_id = s.seat_id+1 OR  s1.seat_id = s.seat_id-1))
*/
