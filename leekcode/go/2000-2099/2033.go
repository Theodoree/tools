package _000_2099

import "sort"

/*
2033. 获取单值网格的最小操作数
给你一个大小为 m x n 的二维整数网格 grid 和一个整数 x 。每一次操作，你可以对 grid 中的任一元素 加 x 或 减 x 。
单值网格 是全部元素都相等的网格。
返回使网格化为单值网格所需的 最小 操作数。如果不能，返回 -1 。
*/
func minOperations(grid [][]int, x int) (ans int) {
	n := len(grid) * len(grid[0])
	a := make([]int, 0, n)
	for _, row := range grid {
		for _, v := range row {
			if (v-grid[0][0])%x != 0 { // 以其中一元素为基准，若所有元素与它的差均为 x 的倍数，则任意两元素之差为 x 的倍数
				return -1
			}
		}
		a = append(a, row...)
	}
	sort.Ints(a) // 除了排序，也可以用求第 k 大算法来找中位数
	for _, v := range a {
		ans += abs(v-a[n/2]) / x
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
