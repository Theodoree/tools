package _go

import "math/rand"

/*
剑指 Offer II 071. 按权重生成随机数
给定一个正整数数组 w ，其中 w[i] 代表下标 i 的权重（下标从 0 开始），请写一个函数 pickIndex ，它可以随机地获取下标 i，选取下标 i 的概率与 w[i] 成正比。
例如，对于 w = [1, 3]，挑选下标 0 的概率为 1 / (1 + 3) = 0.25 （即，25%），而选取下标 1 的概率为 3 / (1 + 3) = 0.75（即，75%）。
也就是说，选取下标 i 的概率为 w[i] / sum(w) 。
*/
type Solution struct {
	buf  []int
	size int
}

func Constructor(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] = w[i] + w[i-1]
	}
	size := w[len(w)-1]
	return Solution{buf: w, size: size}

}

func (s *Solution) PickIndex() int {
	r := rand.Intn(s.size)
	left := 0
	right := len(s.buf)

	for left < right {
		mid := (left + right) >> 1
		if s.buf[mid] > r {
			right = mid
		} else {
			left = mid+1
		}
	}


	return left
}

