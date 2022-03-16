package _00_199

import (
	"sort"
)

/*
128. 最长连续序列
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
*/
func longestConsecutive(nums []int) int {
	sort.Ints(nums) // 用map可以实现O(n),但是也太蠢了,开销太大了
	var max int
	cur := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			continue
		}
		if nums[i-1]+1 == nums[i] {
			cur++
		} else {
			if cur > max {
				max = cur
			}
			cur = 1
		}
	}

	if cur > max {
		max = cur
	}

	return max
}
