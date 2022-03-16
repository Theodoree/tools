package _700_1799

/*
1708. 长度为 K 的最大子数组
在数组 A 和数组 B 中，对于第一个满足 A[i] != B[i] 的索引 i ，当 A[i] > B[i] 时，数组 A 大于数组 B。
例如，对于索引从 0 开始的数组：
[1,3,2,4] > [1,2,2,4] ，因为在索引 1 上， 3 > 2。
[1,4,4,4] < [2,1,1,1] ，因为在索引 0 上， 1 < 2。
一个数组的子数组是原数组上的一个连续子序列。
给定一个包含不同整数的整数类型数组 nums ，返回 nums 中长度为 k 的最大子数组。
*/
func largestSubarray(nums []int, k int) []int {
	var max int
	var maxIdx int
	for i := 0; i <= len(nums)-k; i++ {
		if nums[i] > max {
			max = nums[i]
			maxIdx = i
		}

	}
	return nums[maxIdx : maxIdx+k]
}
