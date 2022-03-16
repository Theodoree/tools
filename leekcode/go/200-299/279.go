package _00_299





/*
279. 完全平方数

给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。

示例 1:

输入: n = 12
输出: 3
解释: 12 = 4 + 4 + 4.
示例 2:

输入: n = 13
输出: 2
解释: 13 = 4 + 9.
*/

func numSquares(n int) int {

    var f []int

    cnt := 1

    for cnt <= n {
        f = append(f, cnt*cnt)
        cnt++
    }

    var dp = make([]int, n+1)
    for i := range dp {
        dp[i] = math.MaxInt32
    }

    dp[0] = 0

    for _, v := range f {

        for i := v; i <= n; i++ {
            if dp[i-v] != math.MaxInt32 {
                dp[i] = min(dp[i], dp[i-v]+1) // 最优子解
            }
        }
    }

    if dp[n] == math.MaxInt32 {
        return -1
    }

    return dp[n]
}

func numSquares(n int) int {
    var d = make([]int, n+1) // d[i]表示i能够表示为完全平方数的最小个数

    // 动态规划
    for i := 1; i <= n; i++ {
        d[i] = i // 最坏情况是全由1组成
        for j := 1; i-j*j >= 0; j++ {
            if d[i] > d[i-j*j]+1 { // 在d[i-j*j]的基础上再加一个完全平方数j*j
                d[i] = d[i-j*j] + 1
            }
        }
    }
    return d[n]
}

func min(a, b int) int {

    if a > b {
        return b
    }

    return a
}