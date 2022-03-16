package _400_1499


/*
1492. n 的第 k 个因子
给你两个正整数 n 和 k 。
如果正整数 i 满足 n % i == 0 ，那么我们就说正整数 i 是整数 n 的因子。
考虑整数 n 的所有因子，将它们 升序排列 。请你返回第 k 个因子。如果 n 的因子数少于 k ，请你返回 -1 。
*/
func kthFactor(n int, k int) int {


	for i:=1;i<=n;i++{
		if i%n == 0 {
			k--
		}
		if k == 0 {
			return i
		}
	}

	return -1

}
