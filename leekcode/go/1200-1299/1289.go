package _200_1299


func minFallingPathSum(grid [][]int) int {
	// dp[i]p[j] = grid[i][j] + min(dp[i-1][col] col != j)
	n := len(grid)
	dp := make([]int, n)
	for i := range grid{
		tmp := make([]int, n)
		for j := range grid[i]{
			tmp[j] = grid[i][j]
			// t := min(append(dp[:j], dp[j+1:]...)...)  slice 底层数组 浅copy
			t := min(j, dp...)
			if t != math.MaxInt32{
				tmp[j] += t
			}
		}
		dp = tmp
	}
	return min(-1, dp...)
}
// 排除except 下标的 值
func min(except int, nums ...int) int{
	m := math.MaxInt32
	for i := range nums{
		if m > nums[i] && i != except{
			m = nums[i]
		}
	}
	return m
}
