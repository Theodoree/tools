package _00_699



/*
674. 最长连续递增序列
给定一个未经排序的整数数组，找到最长且连续的的递增序列。

示例 1:

输入: [1,3,5,4,7]
输出: 3
解释: 最长连续递增序列是 [1,3,5], 长度为3。
尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为5和7在原数组里被4隔开。
*/

func findLengthOfLCIS(nums []int) int {

    switch {
    case len(nums) == 0:
        return 0

    }

    var max int
    cnt := 1
    for i := 0; i < len(nums)-1; i++ {
        if nums[i] < nums[i+1] {
            cnt++
        } else {
            if cnt > max {
                max = cnt
            }
            cnt = 1

        }
    }

    if cnt > max {
        max = cnt
    }

    return max
}

