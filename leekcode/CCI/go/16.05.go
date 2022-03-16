package _go

/*
面试题 16.05. 阶乘尾数
设计一个算法，算出 n 阶乘有多少个尾随零。
*/
func trailingZeroes(n int) int {
	count := 0
	for n >= 5 {
		n /= 5
		count += n
	}
	return count
}
