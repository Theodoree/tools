package _000_1099

import "sort"

/*
1051. 高度检查器
学校在拍年度纪念照时，一般要求学生按照 非递减 的高度顺序排列。

请你返回至少有多少个学生没有站在正确位置数量。该人数指的是：能让所有学生以 非递减 高度排列的必要移动人数。
*/

func heightChecker(heights []int) int {
    var cnt int
    f := make([]int,len(heights))
    copy(f,heights)
    sort.Ints(f)

    for i := 0; i < len(heights); i++ {
        if heights[i] != f[i] {
            cnt++
        }
    }
    return cnt
}
