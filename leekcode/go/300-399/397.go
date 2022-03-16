package _00_399

/*
397. 整数替换
给定一个正整数 n ，你可以做如下操作：

如果 n 是偶数，则用 n / 2替换 n 。
如果 n 是奇数，则可以用 n + 1或n - 1替换 n 。
n 变为 1 所需的最小替换次数是多少？


*/
func integerReplacement(n int) int {
	//target := 0xffffffffffff
	var count int
	for n > 1 {
		if n%2 == 0 {
			n >>= 1
		} else {
			if n&1 == 1 && n&2 == 2 && n != 3 {
				n += 1
			} else {
				n -= 1
			}
		}
		count++
	}
	return count
}
