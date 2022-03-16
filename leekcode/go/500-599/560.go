package _00_599

/*
560. 和为 K 的子数组
给你一个整数数组 nums 和一个整数 k ，请你统计并返回该数组中和为 k 的连续子数组的个数。
*/
func subarraySum(nums []int, k int) int {
	n := len(nums)
	preSum := make(map[int]int, n)
	preSum[0] = 1
	ans := 0
	sum0_i := 0
	for i := 0; i < n; i++ {
		sum0_i += nums[i]
		sum0_j := sum0_i - k
		if v, ok := preSum[sum0_j]; ok {
			ans += v
		}
		preSum[sum0_i] += 1
	}
	return ans
}
