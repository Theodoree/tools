package _go

/*
剑指 Offer 56 - II. 数组中数字出现的次数 II
在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。
*/

func singleNumber(nums []int) int {
	var a, b int
	for _, item := range nums {
		a = a ^ item & ^b
		b = b ^ item & ^a
	}
	return a
}
