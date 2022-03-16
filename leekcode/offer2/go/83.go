package _go

/*
剑指 Offer II 083. 没有重复元素集合的全排列
给定一个不含重复数字的整数数组 nums ，返回其 所有可能的全排列 。可以 按任意顺序 返回答案。
*/
func permute(nums []int) [][]int {
	var result [][]int

	var target int
	//for i := 0; i < len(nums); i++ {
	//	target |= 1 << i
	//}

	Permute(nums, 0, target, &result, nil)
	return result
}

func Permute(nums []int, flag int, target int, result *[][]int, cur []int) {

	if flag == target {
		*result = append(*result, cur)
		return
	}

	for i := 0; i < len(nums); i++ {
		if flag&(1<<i) > 0 {
			continue
		}
		Permute(nums, flag|1<<i, target, result, append(cur, nums[i]))
	}
}
