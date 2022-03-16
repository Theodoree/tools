package _go

/*
面试题 05.06. 整数转换
整数转换。编写一个函数，确定需要改变几个位才能将整数A转成整数B。

示例1:

 输入：A = 29 （或者0b11101）, B = 15（或者0b01111）
 输出：2
示例2:

 输入：A = 1，B = 2
 输出：2
提示:

A，B范围在[-2147483648, 2147483647]之间
通过次数6,367提交次数12,399
*/


func convertInteger(A int, B int) int {

	var bit int

	f := A ^ B

	for i := 0; i < 32; i++ {
		if 1<<i&f > 0 {
			bit++
		}

	}

	return bit

}
