package _00_299



/*
216. 组合总和 III
找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
说明：
所有数字都是正整数。
解集不能包含重复的组合。
*/
func combinationSum3(k int, n int) [][]int {
	m := make(map[int]struct{})
	var result [][]int
	Backtracking(0, k, n, m, &result)
	return result
}

func Backtracking(flag int, k int, target int, m map[int]struct{}, result *[][]int) {
	if target < 0 || k < 0 {
		return
	}
	if target == 0 && k == 0 {
		if _, ok := m[flag]; !ok {
			var arr []int
			for i := 1; i <= 9; i++ {
				if flag&(1<<i) > 0 {
					arr = append(arr, i)
				}
			}
			m[flag] = struct{}{}
			*result = append(*result, arr)
		}
		return
	}
	for i := 1; i <= 9; i++ {
		if flag&(1<<i) > 0 {
			continue
		}
		Backtracking(flag|(1<<i), k-1, target-i, m, result)
	}
}
