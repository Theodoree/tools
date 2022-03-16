package __99

import "sort"

/*
90. 子集 II
给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
*/
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	dfs(nums, 0, nil, &result)

	var result1 [][]int
	var m = make(map[string]struct{})
	var buf = make([]byte, 0, 10)
	for _, v := range result {
		for j := 0; j < len(v); j++ {
			if v[j] < 0 {
				buf = append(buf, byte(50+v[j]))
			} else {
				buf = append(buf, byte(v[j]))
			}
		}

		if _, ok := m[string(buf)]; !ok {
			m[string(buf)] = struct{}{}
			result1 = append(result1, v)

		}
		buf = buf[:0]
	}

	return result1

}
func dfs(n []int, i int, cur []int, result *[][]int) {
	if i == len(n) {
		*result = append(*result, cur)
		return
	}

	dfs(n, i+1, cur, result)
	buf := make([]int, len(cur))
	copy(buf, cur)
	buf = append(buf, n[i])
	dfs(n, i+1, buf, result)
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0, 1<<n)
	var dfs func(idx, lastChoiceIdx int, path []int)
	dfs = func(idx, lastChoiceIdx int, path []int) {
		if idx == n {
			a := make([]int, len(path))
			copy(a, path)
			ans = append(ans, a)
			return
		}
		//选择,不出现重复的条件:当前为第一个元素,或当前元素与上一个元素不同,或上一个元素被选择
		if idx == 0 || nums[idx] != nums[idx-1] || (lastChoiceIdx == idx-1) {
			path = append(path, nums[idx])
			dfs(idx+1, idx, path)
			path = path[:len(path)-1]
		}
		//不选择
		dfs(idx+1, lastChoiceIdx, path)
	}
	dfs(0, -1, make([]int, 0, n))
	return ans
}
