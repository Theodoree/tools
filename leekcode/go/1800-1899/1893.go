package _800_1899

import "sort"

/*
1893. 检查是否区域内所有整数都被覆盖
给你一个二维整数数组 ranges 和两个整数 left 和 right 。每个 ranges[i] = [starti, endi] 表示一个从 starti 到 endi 的 闭区间 。
如果闭区间 [left, right] 内每个整数都被 ranges 中 至少一个 区间覆盖，那么请你返回 true ，否则返回 false 。
已知区间 ranges[i] = [starti, endi] ，如果整数 x 满足 starti <= x <= endi ，那么我们称整数x 被覆盖了。
*/
func isCovered(ranges [][]int, left int, right int) bool {

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][1] < ranges[j][1]
	})

	leftNum := -1
	for i := 0; i < len(ranges); i++ {
		if ranges[i][0] <= left {
			leftNum = ranges[i][1]
			continue
		}

		if leftNum >= 0 && ranges[i][0] <= leftNum+1 && ranges[i][1] >= leftNum {
			leftNum = ranges[i][1]
			if leftNum >= right {
				return true
			}
		}
	}

	return leftNum >= right
}
