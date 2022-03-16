package __99

/*
56. 合并区间

给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:

输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
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
        } else if intervals[i][1] > result[len(result)-1][1] && intervals[i][0] > result[len(result)-1][0]  {
            result = append(result, intervals[i])
        }

    }

    return result
}
