package _00_599

import "math/rand"

/*
528. 按权重随机选择
给你一个 下标从 0 开始 的正整数数组 w ，其中 w[i] 代表第 i 个下标的权重。

请你实现一个函数 pickIndex ，它可以 随机地 从范围 [0, w.length - 1] 内（含 0 和 w.length - 1）选出并返回一个下标。选取下标 i 的 概率 为 w[i] / sum(w) 。

例如，对于 w = [1, 3]，挑选下标 0 的概率为 1 / (1 + 3) = 0.25 （即，25%），而选取下标 1 的概率为 3 / (1 + 3) = 0.75（即，75%）。
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
