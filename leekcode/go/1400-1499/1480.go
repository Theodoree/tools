package _400_1499

/*
1480. 一维数组的动态和
给你一个数组 nums 。数组「动态和」的计算公式为：runningSum[i] = sum(nums[0]…nums[i]) 。

请返回 nums 的动态和。
*/
func runningSum(nums []int) []int {

	var count int
	for i := 0; i < len(nums); i++ {

		c := count
		count += nums[i]
		nums[i] += c

	}

	return nums
}
