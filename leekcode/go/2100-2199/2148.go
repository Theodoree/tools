package _100_2199


/*
2148. 元素计数
给你一个整数数组 nums ，统计并返回在 nums 中同时至少具有一个严格较小元素和一个严格较大元素的元素数目。

*/
func countElements(nums []int) int {
	min, max, cntMin, cntMax := nums[0], nums[0], 1, 1
	for _, v := range nums[1:] {
		if v < min {
			min, cntMin = v, 1
		} else if v == min {
			cntMin++
		}
		if v > max {
			max, cntMax = v, 1
		} else if v == max {
			cntMax++
		}
	}
	if min == max {
		return 0
	}
	return len(nums) - cntMin - cntMax
}
