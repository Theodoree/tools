package _go

/*
剑指 Offer 17. 打印从1到最大的n位数
输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
*/

func printNumbers(n int) []int {
	if n == 0 {
		return nil
	}
	var max int
	for n > 0 {
		max = max*10 + 9
		n--
	}

	var result = make([]int, max)

	for index := range result {
		result[index] = index + 1
	}
	return result
}
