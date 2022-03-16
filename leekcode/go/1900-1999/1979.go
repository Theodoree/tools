package _900_1999

/*
1979. 找出数组的最大公约数
给你一个整数数组 nums ，返回数组中最大数和最小数的 最大公约数 。

两个数的 最大公约数 是能够被两个数整除的最大正整数。
*/
func findGCD(nums []int) int {
	max := 0
	min := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}

	for i := min; i >= 0; i-- {
		if max%i == 0 && min%i == 0 {
			return i
		}
	}
	return 1
}
