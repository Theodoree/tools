package __99

/*
62. 不同路径

一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

问总共有多少条不同的路径？



例如，上图是一个7 x 3 的网格。有多少可能的路径？

说明：m 和 n 的值均不超过 100。

示例 1:

输入: m = 3, n = 2
输出: 3
解释:
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向右 -> 向下
2. 向右 -> 向下 -> 向右
3. 向下 -> 向右 -> 向右
示例 2:

输入: m = 7, n = 3
输出: 28

*/

// 取巧
func uniquePaths(m int, n int) int {

	dp := make([]int, n)
	for k, _ := range dp {
		dp[k] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}

	return dp[n-1]

}

// 递归
func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	nowMap := map[int]int{}

	return calc(m, n, nowMap)
}

func calc(m int, n int, nowMap map[int]int) int {
	if m == 1 || n == 1 {
		return 1
	}
	nowKey := m*100 + n
	if nowVal, ok := nowMap[nowKey]; ok {
		return nowVal
	} else {
		nowMap[nowKey] = calc(m-1, n, nowMap) + calc(m, n-1, nowMap)
		return nowMap[nowKey]
	}
}
