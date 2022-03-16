package _400_1499

import "sort"

/*
1464. 数组中两元素的最大乘积
给你一个整数数组 nums，请你选择数组的两个不同下标 i 和 j，使 (nums[i]-1)*(nums[j]-1) 取得最大值。
请你计算并返回该式的最大值。
*/
func maxProduct(nums []int) int {
	sort.Ints(nums)

	return (nums[len(nums)-1] - 1) * (nums[len(nums)-2] - 1)
}
