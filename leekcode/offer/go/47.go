package _go

/*
剑指 Offer 47. 礼物的最大价值
在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
*/
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxValue(grid [][]int) int {
	dp := make([]int, len(grid[0]))

	dp[0] = grid[0][0]
	for i := 1; i < len(grid[0]); i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}
	for i := 1; i < len(grid); i++ {
		for j := 0; j < len(dp); j++ {
			if j == 0 {
				dp[j] = dp[j] + grid[i][j]
			} else {
				dp[j] = max(dp[j-1], dp[j]) + grid[i][j]
			}
		}
	}
	return dp[len(dp)-1]
}
