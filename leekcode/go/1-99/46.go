package __99

/*
46. 全排列
给定一个没有重复数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/
func permute(nums []int) [][]int {
	var result [][]int

	var target int
	for i := 0; i < len(nums); i++ {
		target |= 1 << i
	}

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
