package _800_1899

/*
1822. 数组元素积的符号
已知函数 signFunc(x) 将会根据 x 的正负返回特定值：

如果 x 是正数，返回 1 。
如果 x 是负数，返回 -1 。
如果 x 是等于 0 ，返回 0 。
给你一个整数数组 nums 。令 product 为数组 nums 中所有元素值的乘积。

返回 signFunc(product) 。
*/
func arraySign(nums []int) int {
	var sum int
	sum = 1
	for _, v := range nums {
		if v == 0 {
			return 0
		}
		sum *= v
		if sum < 0 {
			sum %= -0xffff
		} else {
			sum %= 0xffff
		}
	}

	if sum > 0 {
		return 1
	}
	return -1
}
