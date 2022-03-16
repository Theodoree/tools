package _300_1399

/*
1317. 将整数转换为两个无零整数的和
「无零整数」是十进制表示中 不含任何 0 的正整数。
给你一个整数 n，请你返回一个 由两个整数组成的列表 [A, B]，满足：
A 和 B 都是无零整数
A + B = n
题目数据保证至少有一个有效的解决方案。
如果存在多个有效解决方案，你可以返回其中任意一个。
*/

func getNoZeroIntegers(n int) []int {
	for i := 1; i < n; i++ {
		a := i
		b := n - i

		if !hasZero(a) && !hasZero(b) {
			return []int{a, b}
		}
	}
	return nil
}

func hasZero(i int) bool {
	for i > 0 {
		if i%10 == 0 {
			return true
		}
		i /= 10
	}
	return false
}
