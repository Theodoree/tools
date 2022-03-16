package _go

import "math"

/*
剑指 Offer II 008. 和大于等于 target 的最短子数组
给定一个含有 n 个正整数的数组和一个正整数 target 。
找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
*/
func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func minSubArrayLen(s int, nums []int) int {

	ans := math.MaxInt64
	var left, sum int

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= s {
			ans = min(ans, i+1-left)
			sum -= nums[left]
			left++
		}
	}
	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}
