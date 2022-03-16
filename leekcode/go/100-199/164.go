package _00_199

import (
    "fmt"
    "sort"
)

/*
164. 最大间距

给定一个无序的数组，找出数组在排序之后，相邻元素之间最大的差值。

如果数组元素个数小于 2，则返回 0。
*/
func maximumGap(nums []int) int {
    if len(nums) < 2 {
        return 0
    }

    var max int
    sort.Ints(nums)
    for i := 1; i < len(nums); i++ {
        sum := nums[i] - nums[i-1]
        if sum > max {
            max = sum
        }
    }

    return max
}

func maximumGap1(nums []int) int {
    numsLen := len(nums)
    if numsLen < 2 {
        return 0
    }

    var maxInt =  1<<63 - 1
    var minInt = ^maxInt
    fmt.Println(maxInt, minInt)
    var min = maxInt
    var max = minInt

    for _, v := range nums {
        if v < min {
            min = v
        }
        if v > max {
            max = v
        }
    }
    fmt.Println(max, min)

    buckSize := getMax(1, (max-min)/(numsLen-1))

    type buc struct {
        min int
        max int
    }

    bucksNum := (max-min)/buckSize + 1
    bucks := make([]buc, bucksNum)
    bucksFlag := make([]int, bucksNum)
    for i := 0; i < bucksNum; i++ {
        bucks[i].min = maxInt
        bucks[i].max = minInt
    }

    for _, v := range nums {
        loc := (v - min) / buckSize

        bucksFlag[loc] = 1
        if v < bucks[loc].min {
            bucks[loc].min = v
        }
        if v > bucks[loc].max {
            bucks[loc].max = v
        }
    }

    maxGap := minInt

    var prev buc
    var prevFlag int
    for i, v := range bucks {
        if bucksFlag[i] == 0 {
            continue
        }
        maxGap = getMax(v.max-v.min, maxGap)

        if prevFlag != 0 {
            maxGap = getMax(v.min-prev.max, maxGap)
        }
        prev = v
        prevFlag = 1
    }

    return maxGap
}

func getMax(a int, b int) int {
    if a > b {
        return a
    }
    return b
}