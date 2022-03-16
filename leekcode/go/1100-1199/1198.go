package _100_1199

import "sort"

/*
1198. 找出所有行中最小公共元素
给你一个矩阵 mat，其中每一行的元素都已经按 严格递增 顺序排好了。请你帮忙找出在所有这些行中 最小的公共元素。
如果矩阵中没有这样的公共元素，就请返回 -1。
*/
func smallestCommonElement(arr [][]int) (res int) {
	for j := 0; j < len(arr[0]); j++ {
		ok := true
		for i := 0; i < len(arr); i++ {
			t := sort.SearchInts(arr[i], arr[0][j])
			if arr[i][t] != arr[0][j] {
				ok = false
				break
			}
		}
		if ok {
			return arr[0][j]
		}
	}
	return -1
}
