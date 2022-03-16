package _go

/*
剑指 Offer II 070. 排序数组中只出现一次的数字
给定一个只包含整数的有序数组 nums ，每个元素都会出现两次，唯有一个数只会出现一次，请找出这个唯一的数字。
*/
func singleNonDuplicate(nums []int) int {

	var cnt int

	for _, v := range nums {
		cnt ^= v
	}

	return cnt
}
