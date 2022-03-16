package _go

import "math"

/*
剑指 Offer 43. 1～n整数中1出现的次数
输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。

例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。

示例 1：

输入：n = 12
输出：5
示例 2：

输入：n = 13
输出：6
*/

/*
1  1次

10 11 21 31 41 51 61 71 81 91 11次

100  101 111 121   11次*9 99次
1000   99次*10 999次
*/

func countDigitOne(n int) int {
	num := n
	i := 1
	s := 0

	for num > 0 {
		if num%10 == 0 {
			s += (num / 10) * i
		}
		if num%10 == 1 {
			s += (num/10)*i + (n % i) + 1
		}
		if num%10 > 1 {
			s += int(math.Ceil(float64(num)/10.0)) * i
		}
		num = num / 10
		i = i * 10

	}
	return s
}
