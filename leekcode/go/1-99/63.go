package __99

/*
63. 不同路径 II
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。
现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
网格中的障碍物和空位置分别用 1 和 0 来表示。
*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	dp := make([][]int, len(obstacleGrid))
	var flag bool
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(obstacleGrid[0]))

		if obstacleGrid[i][0] == 1 {
			flag = true
		}
		if !flag {
			dp[i][0] = 1
		}
	}
	for i := 0; i < len(dp[0]); i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}

	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			if obstacleGrid[i-1][j] == 1 {
				dp[i][j] = max(dp[i][j], dp[i][j-1])
			} else if obstacleGrid[i][j-1] == 1 {
				dp[i][j] = max(dp[i][j], dp[i-1][j])
			} else {
				dp[i][j] = max(dp[i][j], dp[i-1][j]+dp[i][j-1])
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}
