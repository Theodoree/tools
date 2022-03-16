package __99

/*
53. 最大子序和
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
*/

func maxSubArray(nums []int) int {

	var dp = make([]int, len(nums)+1)
	max := func(j, i int) int {
		if j > i {
			return j
		}
		return i
	}
	count := nums[0]
	dp[0] = count
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > count {
			count = dp[i]
		}
	}

	return count
}
