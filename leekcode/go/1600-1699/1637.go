package _600_1699



/*
1637. 两点之间不包含任何点的最宽垂直面积
给你 n 个二维平面上的点 points ，其中 points[i] = [xi, yi] ，请你返回两点之间内部不包含任何点的 最宽垂直面积 的宽度。
垂直面积 的定义是固定宽度，而 y 轴上无限延伸的一块区域（也就是高度为无穷大）。 最宽垂直面积 为宽度最大的一个垂直面积。
请注意，垂直区域 边上 的点 不在 区域内。

class Solution:
    def maxWidthOfVerticalArea(self, points: List[List[int]]) -> int:
        d = sorted({x for x, y in points})
        return max(d[i] - d[i-1] for i in range(1, len(d)))
*/

func maxWidthOfVerticalArea(points [][]int) int {

	arr:=make([]int,0,len(points))
	for i:=0;i<len(points);i++{
		arr = append(arr,points[i][0])
	}
	sort.Ints(arr)


	var max int
	for i:=1;i<len(arr);i++{
		if arr[i]  - arr[i-1] > max {
			max = arr[i]  - arr[i-1]
		}
	}

	return max

}

