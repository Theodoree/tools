package _go

/*
面试题 16.17. 连续数列
给定一个整数数组，找出总和最大的连续数列，并返回总和。
示例：
输入： [-2,1,-3,4,-1,2,1,-5,4]
输出： 6
解释： 连续子数组 [4,-1,2,1] 的和最大，为 6。
*/
func maxSubArray(nums []int) int {

	maxSum := math.MinInt64
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		if sum > maxSum {
			maxSum = sum
		}
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	return maxSum
}
