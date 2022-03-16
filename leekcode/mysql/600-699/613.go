package _00_699

/*
613. 直线上的最近距离

表 point 保存了一些点在 x 轴上的坐标，这些坐标都是整数。



写一个查询语句，找到这些点中最近两个点之间的距离。



| x   |
|-----|
| -1  |
| 0   |
| 2   |


最近距离显然是 '1' ，是点 '-1' 和 '0' 之间的距离。所以输出应该如下：



| shortest|
|---------|
| 1       |


注意：每个点都与其他点坐标不同，表 table 不会有重复坐标出现。


*/

/*
SELECT min(abs(p1.x - p2.x)) shortest
FROM  point p1 inner join point p2
ON p1.x != p2.x
*/
