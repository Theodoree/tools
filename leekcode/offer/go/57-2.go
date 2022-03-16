package _go

/*
面试题57 - II. 和为s的连续正数序列

输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

示例 1：
输入：target = 9
输出：[[2,3,4],[4,5]]
示例 2：

输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]
*/
func findContinuousSequence(target int) [][]int {
    var ret [][]int
    for i := 1; i < target; i++ {
        t := []int{i}
        for j := i + 1; (i+j)*(j-i+1)/2 <= target; j++ {
            t = append(t, j)
            if (i+j)*(j-i+1)/2 == target {
                ret = append(ret, t)
            }
        }
    }
    return ret
}
