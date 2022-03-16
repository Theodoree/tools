package _go

import "math"

/*
剑指 Offer II 103. 最少的硬币数目
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
你可以认为每种硬币的数量是无限的。
*/
func coinChange(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}

	var dp = make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}

	dp[0] = 0 // 边界条件
	for _, coin := range coins {

		for i := coin; i <= amount; i++ {
			if dp[i-coin] != math.MaxInt32 {
				dp[i] = min(dp[i], dp[i-coin]+1) // 最优子解
			}
		}

	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}

func min(a, b int) int {

	if a > b {
		return b
	}

	return a
}
