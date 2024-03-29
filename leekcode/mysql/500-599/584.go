package _00_599

/*
584. 寻找用户推荐人

给定表 customer ，里面保存了所有客户信息和他们的推荐人。

+------+------+-----------+
| id   | name | referee_id|
+------+------+-----------+
|    1 | Will |      NULL |
|    2 | Jane |      NULL |
|    3 | Alex |         2 |
|    4 | Bill |      NULL |
|    5 | Zack |         1 |
|    6 | Mark |         2 |
+------+------+-----------+
写一个查询语句，返回一个编号列表，列表中编号的推荐人的编号都 不是 2。

对于上面的示例数据，结果为：

+------+
| name |
+------+
| Will |
| Jane |
| Bill |
| Zack |
+------+

SELECT
    name
FROM
    customer
WHERE
    IFNULL(referee_id,0)!=2;

*/
