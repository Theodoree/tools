package _go

/*
剑指 Offer II 079. 所有子集
给定一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
*/
func subsets(nums []int) [][]int {
	var result [][]int
	var l uint = uint(len(nums))
	var n = 1 << l
	for i := 0; i < n; i++ {
		var cur []int
		for j := 0; j < int(l); j++ {
			if i&(1<<uint(j)) == 0 {
				cur = append(cur, nums[j])
			}
		}
		result = append(result, cur)
	}

	return result
}
