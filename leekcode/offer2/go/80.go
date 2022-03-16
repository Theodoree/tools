package _go

/*
剑指 Offer II 080. 含有 k 个元素的组合
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
*/
func combine(n int, k int) [][]int {
	l := make([]int, k)
	i := 0
	var ret [][]int
	for i >= 0 {
		l[i]++
		if l[i] > n-k+i+1 {
			i--
		} else if i == k-1 {
			t := make([]int, k)
			copy(t, l)
			ret = append(ret, t)
		} else {
			i++
			l[i] = l[i-1]
		}
	}
	return ret
}
