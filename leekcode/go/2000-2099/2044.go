package _000_2099



/*
2044. 统计按位或能得到最大值的子集数目
给你一个整数数组 nums ，请你找出 nums 子集 按位或 可能得到的 最大值 ，并返回按位或能得到最大值的 不同非空子集的数目 。
如果数组 a 可以由数组 b 删除一些元素（或不删除）得到，则认为数组 a 是数组 b 的一个 子集 。如果选中的元素下标位置不一样，则认为两个子集 不同 。
对数组 a 执行 按位或 ，结果等于 a[0] OR a[1] OR ... OR a[a.length - 1]（下标从 0 开始）。
*/

func countMaxOrSubsets(nums []int) int {
	var res int
	var target int
	n := len(nums)
	for _, num := range nums {
		target |= num
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	var dfs func(idx, curr int)
	dfs = func(idx, curr int) {
		if curr == target {
			res += 1 << (n - idx)
			return
		}
		if idx == n {
			return
		}
		dfs(idx+1, curr|nums[idx])
		dfs(idx+1, curr)
	}
	dfs(0, 0)
	return res
}
