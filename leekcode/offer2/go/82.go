package _go

import "sort"

/*
剑指 Offer II 082. 含有重复元素集合的组合
给定一个可能有重复数字的整数数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的每个数字在每个组合中只能使用一次，解集不能包含重复的组合。
*/
func combinationSum2(candidates []int, target int) [][]int {

	sort.Ints(candidates)
	var result [][]int
	CombinationSum2(candidates, &result, []int{}, 0, target)
	return result
}

func CombinationSum2(nums []int, result *[][]int, cur []int, level int, target int) {

	if target == 0 {
		var c= make([]int, len(cur))
		copy(c, cur)
		*result = append(*result, c)
		return
	} else if target < 0 {
		return
	}

	for i := level; i < len(nums); i++ {

		if target-nums[i] >= 0 {
			if i > level && nums[i] == nums[i-1] {
				continue
			}
			cur = append(cur, nums[i])
			CombinationSum2(nums, result, cur, i+1, target-nums[i])
			cur = cur[0 : len(cur)-1]
		} else {
			return
		}
	}

}