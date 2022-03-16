package _go

/*
剑指 Offer II 072. 求平方根
给定一个非负整数 x ，计算并返回 x 的平方根，即实现 int sqrt(int x) 函数。

正数的平方根有两个，只输出其中的正数平方根。

如果平方根不是整数，输出只保留整数的部分，小数部分将被舍去。
*/
func mySqrt(x int) int {

	if x == 0 || x == 1 {
		return x
	}

	i := (x) / 2.0

	for i*i > x {
		i = (i + x/i) / 2
	}
	return i
}
