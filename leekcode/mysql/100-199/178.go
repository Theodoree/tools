package _00_199

/*
178. 分数排名


SELECT
  Score,		//这里的prev是一个局部变量,<>代表大于或小于,可以理解为永true,而在计算机中true转换成num则是1
  @rank := @rank + (@prev <> (@prev := Score)) AS  a
FROM
  Scores,
  (SELECT @rank := 0, @prev := -1) AS a
ORDER BY Score desc


*/
