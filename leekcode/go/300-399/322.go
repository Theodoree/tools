package _00_399


/*
322. 零钱兑换

给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

示例 1:

输入: coins = [1, 2, 5], amount = 11
输出: 3
解释: 11 = 5 + 5 + 1
示例 2:

输入: coins = [2], amount = 3
输出: -1
说明:

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

    if dp[amount] ==  math.MaxInt32 {
        return  -1
    }

    return dp[amount]
}

func min(a, b int) int {

    if a > b {
        return b
    }

    return a
}
