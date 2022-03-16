package _go

/*
剑指 Offer II 003. 前 n 个数字二进制中 1 的个数
给定一个非负整数 n ，请计算 0 到 n 之间的每个数字的二进制表示中 1 的个数，并输出一个数组。
*/

func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		/*位运算更快
		  i>>1 == i/2
		  i&1 == i%2*/

		/*观察规律，如
		  5  : 101
		  10 : 1010
		  10 : 1011*/
		res[i] = res[i>>1] + i&1
	}
	return res
}
