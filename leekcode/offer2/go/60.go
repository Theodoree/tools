package _go

import "sort"

/*
剑指 Offer II 060. 出现频率最高的 k 个数字
给定一个整数数组 nums 和一个整数 k ，请返回其中出现频率前 k 高的元素。可以按 任意顺序 返回答案
*/

func topKFrequent(nums []int, k int) []int {

	var buf []int
	var m = map[int]int{}
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			buf = append(buf, v)
		}
		m[v]++
	}

	sort.Slice(buf, func(i, j int) bool {
		return m[buf[i]] > m[buf[j]]
	})
	return buf[:k]
}
