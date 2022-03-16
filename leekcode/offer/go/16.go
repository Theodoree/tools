package _go

import "math"

/*
剑指 Offer 16. 数值的整数次方
实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。
*/
func myPow(x float64, n int) float64 {

	return math.Pow(x, float64(n))
}