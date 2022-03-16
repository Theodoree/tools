package _go

/*
剑指 Offer II 081. 允许重复选择元素的组合
给定一个无重复元素的正整数数组 candidates 和一个正整数 target ，找出 candidates 中所有可以使数字和为目标数 target 的唯一组合。

candidates 中的数字可以无限制重复被选取。如果至少一个所选数字数量不同，则两种组合是唯一的。

对于给定的输入，保证和为 target 的唯一组合数少于 150 个。
*/

func combinationSum(candidates []int, target int) [][]int {
	comb := [][]int{}
	for i := range candidates {
		if candidates[i] == target {
			comb = append(comb, []int{candidates[i]})
		} else if candidates[i] < target {
			rtn := combinationSum(candidates[i:], target-candidates[i])
			for j := range rtn {
				if len(rtn[j]) == 0 {
					continue
				}
				comb = append(comb, append([]int{candidates[i]}, rtn[j]...))
			}
		}
	}
	return comb
}
