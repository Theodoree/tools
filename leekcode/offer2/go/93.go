package _go


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
