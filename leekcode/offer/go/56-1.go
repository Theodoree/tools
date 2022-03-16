package _go

/*
剑指 Offer 56 - I. 数组中数字出现的次数
一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。

示例 1：
输入：nums = [4,1,4,6]
输出：[1,6] 或 [6,1]
示例 2：
输入：nums = [1,2,10,4,1,4,3,3]
输出：[2,10] 或 [10,2]


*/

func singleNumbers(nums []int) []int {
	var x int // 所有值异或后的结果
	for i := range nums {
		x ^= nums[i]
	}

	var d int = 1
	for {
		if (d & x) == 0 {
			d <<= 1
		} else {
			break
		}
	}

	var a = 0
	var b = 0

	for i := range nums {
		if (d & nums[i]) > 0 {
			a ^= nums[i]
		} else {
			b ^= nums[i]
		}
	}
	return []int{a, b}
}
