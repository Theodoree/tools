package _00_899




/*
873. 最长的斐波那契子序列的长度
如果序列 X_1, X_2, ..., X_n 满足下列条件，就说它是 斐波那契式 的：
n >= 3
对于所有 i + 2 <= n，都有 X_i + X_{i+1} = X_{i+2}
给定一个严格递增的正整数数组形成序列 arr ，找到 arr 中最长的斐波那契式的子序列的长度。如果一个不存在，返回  0 。
（回想一下，子序列是从原序列 arr 中派生出来的，它从 arr 中删掉任意数量的元素（也可以不删），而不改变其余元素的顺序。例如， [3, 5, 8] 是 [3, 4, 5, 6, 7, 8] 的一个子序列）
*/
func lenLongestFibSubseq(arr []int) int {
	var m = make(map[int]int)
	for idx, v := range arr {
		m[v] = idx
	}

	var dp = make([][]int, len(arr))
	for idx := range dp {
		dp[idx] = make([]int, len(arr))
	}

	max := func(t1, t2 int) int {
		if t1 > t2 {
			return t1
		}
		return t2
	}
	var result int
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			subVal := arr[i] - arr[j]
			val, ok := m[subVal]
			if ok && val < j {
				dp[i][j] = dp[j][val] + 1
			} else {
				dp[i][j] = 2
			}
			result = max(result, dp[i][j])

		}
	}
	if result <= 2 {
		return 0
	}
	return result
}