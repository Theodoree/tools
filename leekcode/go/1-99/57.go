package __99

/*
57. 插入区间
给你一个 无重叠的 ，按照区间起始端点排序的区间列表。
在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
*/

func insert(intervals [][]int, newInterval []int) [][]int {

	intervals = append(intervals, newInterval)

	for j := len(intervals) - 1; j >= 1; j-- {
		if intervals[j-1][0] > intervals[j][0] {
			intervals[j-1], intervals[j] = intervals[j], intervals[j-1]
		}
	}

	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] >= intervals[i+1][0] {
			// 合并
			if intervals[i+1][1] > intervals[i][1] {
				intervals[i][1] = intervals[i+1][1]
			}
			copy(intervals[i+1:], intervals[i+2:])
			intervals = intervals[:len(intervals)-1]
			i--
		}
	}

	return intervals
}
