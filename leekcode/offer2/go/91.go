package _go

/*
剑指 Offer II 091. 粉刷房子
假如有一排房子，共 n 个，每个房子可以被粉刷成红色、蓝色或者绿色这三种颜色中的一种，你需要粉刷所有的房子并且使其相邻的两个房子颜色不能相同。
当然，因为市场上不同颜色油漆的价格不同，所以房子粉刷成不同颜色的花费成本也是不同的。每个房子粉刷成不同颜色的花费是以一个 n x 3 的正整数矩阵 costs 来表示的。
例如，costs[0][0] 表示第 0 号房子粉刷成红色的成本花费；costs[1][2] 表示第 1 号房子粉刷成绿色的花费，以此类推。
请计算出粉刷完所有房子最少的花费成本。
*/
func minCost(costs [][]int) int {

	if len(costs) == 0 {
		return 0
	}

	// 分析问题

	for i := 1; i < len(costs); i++ {
		costs[i][0] += Min(costs[i-1][1], costs[i-1][2])
		costs[i][1] += Min(costs[i-1][0], costs[i-1][2])
		costs[i][2] += Min(costs[i-1][0], costs[i-1][1])
	}

	var min int
	min = costs[len(costs)-1][0]
	for _, v := range costs[len(costs)-1] {
		min = Min(min, v)
	}

	return min

}

func Min(a, b int) int {
	if a > b {
		return b
	}

	return a

}
