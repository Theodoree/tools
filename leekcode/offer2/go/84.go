package _go

import "sort"

/*
剑指 Offer II 084. 含有重复元素集合的全排列
给定一个可包含重复数字的整数集合 nums ，按任意顺序 返回它所有不重复的全排列。
*/
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int
	Permute(nums, 0, len(nums)-1, &result)
	return result
}

func Permute(nums []int, left, right int, result *[][]int) {

	if left == right {
		*result = append(*result, nums)
		return
	}

	for i := left; i <= right; i++ {
		if i != left && nums[left] == nums[i] {
			continue
		}

		nums[left], nums[i] = nums[i], nums[left]
		buf := make([]int, len(nums))
		copy(buf[:], nums[:])
		Permute(buf, left+1, right, result)
	}
}
