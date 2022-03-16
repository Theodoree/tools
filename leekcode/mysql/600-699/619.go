package _00_699

/*

619. 只出现一次的最大数字

表 my_numbers 的 num 字段包含很多数字，其中包括很多重复的数字。

你能写一个 SQL 查询语句，找到只出现过一次的数字中，最大的一个数字吗？

+---+
|num|
+---+
| 8 |
| 8 |
| 3 |
| 3 |
| 1 |
| 4 |
| 5 |
| 6 |
对于上面给出的样例数据，你的查询语句应该返回如下结果：

+---+
|num|
+---+
| 6 |


SELECT
IFNULL( (SELECT
	MAX(num)
FROM
	my_numbers  AS x
GROUP BY num
HAVING COUNT(1) = 1
ORDER BY num DESC
LIMIT  1
)  ,NULL) AS num
*/
