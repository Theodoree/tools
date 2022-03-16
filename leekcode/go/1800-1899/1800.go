package _800_1899

/*
1800. 最大升序子数组和
给你一个正整数组成的数组 nums ，返回 nums 中一个 升序 子数组的最大可能元素和。

子数组是数组中的一个连续数字序列。

已知子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，若对所有 i（l <= i < r），numsi < numsi+1 都成立，则称这一子数组为 升序 子数组。注意，大小为 1 的子数组也视作 升序 子数组。
*/

func maxAscendingSum(nums []int) int {

	var maxSum int
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			sum += nums[i]
		} else {
			if sum > maxSum {
				maxSum = sum
			}
			sum = nums[i]
		}
	}
	if sum > maxSum {
		maxSum = sum
	}

	return maxSum
}
