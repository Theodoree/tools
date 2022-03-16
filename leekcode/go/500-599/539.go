package _00_599

import (
    "sort"
    "strconv"
    "strings"
)

/*
539. 最小时间差

给定一个 24 小时制（小时:分钟）的时间列表，找出列表中任意两个时间的最小时间差并已分钟数表示。


示例 1：

输入: ["23:59","00:00"]
输出: 1

备注:

列表中时间数在 2~20000 之间。
每个时间取值在 00:00~23:59 之间。
*/

func findMinDifference(timePoints []string) int {

    minuArray := make([]int, len(timePoints))
    var hour, minu int
    for i, v := range timePoints {

        // spilt 比 fmt.Sscanf 快的多
        str:=strings.Split(v,":")
        hour,_ = strconv.Atoi(str[0])
        minu,_ = strconv.Atoi(str[1])
        minuArray[i] = hour*60 + minu

    }

    sort.Ints(minuArray)
    minPoint := 0x7fffffff // MaxInt32
    for i := 0; i < len(timePoints)-1; i++ {
        tmp := minuArray[i+1] - minuArray[i]
        if tmp == 0 {
            return 0
        }
        if tmp < minPoint {
            minPoint = tmp
        }
    }
    if 24*60-minuArray[len(minuArray)-1]+minuArray[0] < minPoint {
        return 24*60 - minuArray[len(minuArray)-1] + minuArray[0]
    }

    return minPoint
}

