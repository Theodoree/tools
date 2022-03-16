package _800_1899

import "sort"

/*
1877. 数组中最大数对和的最小值
一个数对 (a,b) 的 数对和 等于 a + b 。最大数对和 是一个数对数组中最大的 数对和 。
比方说，如果我们有数对 (1,5) ，(2,3) 和 (4,4)，最大数对和 为 max(1+5, 2+3, 4+4) = max(6, 5, 8) = 8 。
给你一个长度为 偶数 n 的数组 nums ，请你将 nums 中的元素分成 n / 2 个数对，使得：
nums 中每个元素 恰好 在 一个 数对中，且
最大数对和 的值 最小 。
请你在最优数对划分的方案下，返回最小的 最大数对和 。
*/
func minPairSum(nums []int) int {
	sort.Ints(nums)
	var cnt int
	for i:=0;i<len(nums)/2;i++{
		result:=nums[i]+nums[len(nums)-1-i]
		if result > cnt {
			cnt = result
		}
	}

	return cnt
}
