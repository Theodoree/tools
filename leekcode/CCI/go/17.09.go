package _go

/*面试题 17.09. 第 k 个数
有些数的素因子只有 3，5，7，请设计一个算法找出第 k 个数。注意，不是必须有这些素因子，而是必须不包含其他的素因子。例如，前几个数按顺序应该是 1，3，5，7，9，15，21。
*/
func getKthMagicNumber(k int) int {
	dp := make([]int, k)
	num3 := 0
	num5 := 0
	num7 := 0
	dp[0] = 1

	min := func(j, i int) int {
		if j < i {
			return j
		}
		return i
	}

	for i := 1; i < k; i++ {
		dp[i] = min(dp[num3]*3, min(dp[num5]*5, dp[num7]*7))
		if dp[i] == dp[num3]*3 {
			num3++
		}
		if dp[i] == dp[num5]*5 {
			num5++
		}
		if dp[i] == dp[num7]*7 {
			num7++
		}
	}
	return dp[k-1]
}
