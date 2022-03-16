package _00_999

import (
    "fmt"
    "sort"
)

/*
949. 给定数字能组成的最大时间

给定一个由 4 位数字组成的数组，返回可以设置的符合 24 小时制的最大时间。

最小的 24 小时制时间是 00:00，而最大的是 23:59。从 00:00 （午夜）开始算起，过得越久，时间越大。

以长度为 5 的字符串返回答案。如果不能确定有效时间，则返回空字符串。



示例 1：

输入：[1,2,3,4]
输出："23:41"
示例 2：

输入：[5,5,5,5]
输出：""


提示：

A.length == 4
0 <= A[i] <= 9
*/

func largestTimeFromDigits(A []int) string {

    var result [][]int
    LargestTimeFromDigits(&result, A, []int{})

    sort.Slice(result, func(i, j int) bool {
        if result[i][0] == result[j][0] {
            return result[i][1] > result[j][1]
        }
        return result[i][0] > result[j][0]
    })
    if len(result) == 0 {
        return ""
    }

    r := result[0]
    t := fmt.Sprintf("%d", r[0])
    if r[0] < 10 {
        t = fmt.Sprintf("0%d", r[0])
    }

    if r[1] < 10 {
        t += fmt.Sprintf(":0%d", r[1])
    } else {
        t += fmt.Sprintf(":%d", r[1])
    }
    return t
}

func LargestTimeFromDigits(result *[][]int, A, cur []int) {
    if len(cur) >= 4 {
        if cur[2]*10+cur[3] >= 60 || cur[0]*10+cur[1] > 23 {
            return
        }
        *result = append(*result, []int{cur[0]*10 + cur[1], cur[2]*10 + cur[3]})
        return
    }

    for i := 0; i < len(A); i++ {
        c := make([]int, len(A))
        copy(c, A)
        c = append(c[:i], c[i+1:]...)
        LargestTimeFromDigits(result, c, append(cur, A[i]))
    }

}
