package _go

/*
剑指 Offer 65. 不用加减乘除做加法
写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。



示例:

输入: a = 1, b = 1
输出: 2


提示：

a, b 均可能是负数或 0
结果不会溢出 32 位整数

*/

func add(a int, b int) int {
	if b == 0 {
		return a
	}
	return add(a^b, int(uint(a&b))<<1)
}
