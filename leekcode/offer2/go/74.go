package _go

import "sort"

/*
剑指 Offer II 074. 合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
*/
func merge(intervals [][]int) [][]int {

	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var result = make([][]int, 1)
	result[0] = intervals[0]
	for i := 1; i < len(intervals); i++ {

		if intervals[i][1] >= result[len(result)-1][1] && intervals[i][0] <= result[len(result)-1][1] {
			result[len(result)-1][1] = intervals[i][1]
		} else if intervals[i][1] > result[len(result)-1][1] && intervals[i][0] > result[len(result)-1][0] {
			result = append(result, intervals[i])
		}

	}

	return result
}
