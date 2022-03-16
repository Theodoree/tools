package _go

/*
剑指 Offer II 090. 环形房屋偷盗
一个专业的小偷，计划偷窃一个环形街道上沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
给定一个代表每个房屋存放金额的非负整数数组 nums ，请计算 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。
*/
func rob(nums []int) int {
	/*
		dp[i][0] = max((不抢i)dp[i-1][0],(抢i)dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]	// 抢i
	*/
	max := func(i, j int) int {
		if j > i {
			return j
		}
		return i
	}
	dp0 := 0
	dp1 := nums[0]
	dp2 := 0
	dp3 := 0

	for i := 1; i < len(nums); i++ {
		//fmt.Println(dp0, dp1, i-1)
		if i < len(nums)-1 {
			tdp0 := max(dp0, dp1)
			tdp1 := dp0 + nums[i]

			dp0 = tdp0
			dp1 = tdp1
		}

		tdp2 := max(dp2, dp3)
		tdp3 := dp2 + nums[i]
		dp2 = tdp2
		dp3 = tdp3

	}

	return max(dp0, max(dp1, max(dp2, dp3)))

}
