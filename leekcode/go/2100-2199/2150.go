package _100_2199


/*
2150. 找出数组中的所有孤独数字
给你一个整数数组 nums 。如果数字 x 在数组中仅出现 一次 ，且没有 相邻 数字（即，x + 1 和 x - 1）出现在数组中，则认为数字 x 是 孤独数字 。
返回 nums 中的 所有 孤独数字。你可以按 任何顺序 返回答案。
*/
func findLonely(nums []int) []int {
	var m = make(map[int]int)
	for i:=0;i<len(nums);i++{
		m[nums[i]]++
	}

	var result []int
	for k,v:=range m{
		if v == 1 &&m[k-1] == 0 && m[k+1] == 0{
			result = append(result,k)
		}
	}

	return result
}
