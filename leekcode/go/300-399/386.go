package _00_399


/*
386. 字典序排数
给你一个整数 n ，按字典序返回范围 [1, n] 内所有整数。

你必须设计一个时间复杂度为 O(n) 且使用 O(1) 额外空间的算法。



示例 1：

输入：n = 13
输出：[1,10,11,12,13,2,3,4,5,6,7,8,9]
示例 2：

输入：n = 2
输出：[1,2]


提示：
1 <= n <= 5 * 104
*/
func lexicalOrder(n int) []int {
	var res = make([]int, n, n)
	cur := 1

	for i := 0; i < n; i++ {
		res[i] = cur
		if cur*10 <= n {
			cur *= 10
		} else {
			if cur >= n {
				cur /= 10
			}
			cur+=1
			for cur%10 == 0 {
				cur /= 10
			}
		}
	}
	return res
}

