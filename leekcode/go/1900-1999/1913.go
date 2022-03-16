package _900_1999

/*
1913. 两个数对之间的最大乘积差
两个数对 (a, b) 和 (c, d) 之间的 乘积差 定义为 (a * b) - (c * d) 。

例如，(5, 6) 和 (2, 7) 之间的乘积差是 (5 * 6) - (2 * 7) = 16 。
给你一个整数数组 nums ，选出四个 不同的 下标 w、x、y 和 z ，使数对 (nums[w], nums[x]) 和 (nums[y], nums[z]) 之间的 乘积差 取到 最大值 。
*/
func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	return (nums[l-1] * nums[l-2]) - (nums[0] * nums[1])
}
