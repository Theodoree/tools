package _800_1899

/*
1828. 统计一个圆中点的数目
给你一个数组 points ，其中 points[i] = [xi, yi] ，表示第 i 个点在二维平面上的坐标。多个点可能会有 相同 的坐标。
同时给你一个数组 queries ，其中 queries[j] = [xj, yj, rj] ，表示一个圆心在 (xj, yj) 且半径为 rj 的圆。
对于每一个查询 queries[j] ，计算在第 j 个圆 内 点的数目。如果一个点在圆的 边界上 ，我们同样认为它在圆 内 。
请你返回一个数组 answer ，其中 answer[j]是第 j 个查询的答案。
*/
func countPoints(points [][]int, queries [][]int) []int {
	result := make([]int, 0, len(queries))
	for _, q := range queries {
		result = append(result, getResult(q[0], q[1], q[2], points))
	}
	return result
}

func getResult(x, y, r int, points [][]int) (result int) {
	for _, v := range points {
		if (v[0]-x)*(v[0]-x)+(v[1]-y)*(v[1]-y) <= r*r {
			result++
		}
	}
	return
}
