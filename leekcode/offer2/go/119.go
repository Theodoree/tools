package _go

import (
	"sort"
)

/*
剑指 Offer II 119. 最长连续序列
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
*/
func longestConsecutive(nums []int) int {
	sort.Ints(nums)
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
