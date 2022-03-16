package _go

/*
面试题 16.01. 交换数字
编写一个函数，不用临时变量，直接交换numbers = [a, b]中a与b的值。
示例：
输入: numbers = [1,2]
输出: [2,1]

建议使用异或处理方案
*/

func swapNumbers(numbers []int) []int {
	numbers[0] += numbers[1]
	numbers[1] = numbers[0] - numbers[1]
	numbers[0] -= numbers[1]
	return numbers
}
