package _00_399

import "math"

/*
313. 超级丑数
超级丑数 是一个正整数，并满足其所有质因数都出现在质数数组 primes 中。
给你一个整数 n 和一个整数数组 primes ，返回第 n 个 超级丑数 。
题目数据保证第 n 个 超级丑数 在 32-bit 带符号整数范围内。

2 3 5
*/
func nthSuperUglyNumber(n int, primes []int) int {

	idxs := make([]int, len(primes))
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		min := math.MaxInt64
		for j := 0; j < len(primes); j++ {
			if primes[j]*dp[idxs[j]] < min {
				min = primes[j] * dp[idxs[j]]
			}
		}
		dp[i] = min
		for j := 0; j < len(primes); j++ {
			if primes[j]*dp[idxs[j]] == min {
				idxs[j]++
			}
		}
	}
	return dp[n-1]
}

func nthSuperUglyNumber1(n int, primes []int) int {
	dp := make([]int, n+1)
	m := len(primes)
	pointers := make([]int, m)
	nums := make([]int, m)
	for i := range nums {
		nums[i] = 1
	}
	for i := 1; i <= n; i++ {
		minNum := math.MaxInt64
		for j := range pointers {
			minNum = min(minNum, nums[j])
		}
		dp[i] = minNum
		for j := range nums {
			if nums[j] == minNum {
				pointers[j]++
				nums[j] = dp[pointers[j]] * primes[j]
			}
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
