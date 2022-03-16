package _00_399

import "fmt"


/*
300. 最长上升子序列

给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
说明:

可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
你算法的时间复杂度应该为 O(n2) 。
进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?

在真实的面试中遇到过这道题？
*/

func lengthOfLIS(nums []int) int {

    if len(nums) == 0 {
        return 0
    }

    var dp = make([]int, len(nums), len(nums))

    dp[0] = 1 // 边界条件 子序列最短为一
    for i := 1; i < len(nums); i++ {

        dp[i] = 1 // b[i] 中储存从 0 -> i 最长的子序列长度

        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                dp[i] = GetMax(dp[i], dp[j]+1)
            }
        }
    }

    var max int
    for i := 0; i < len(dp); i++ {
        max = GetMax(max, dp[i])
    }

    return max

}

func lengthOfLISStack(nums []int) int {
    n := len(nums)
    if n < 2 {
        return n
    }
    var cell = make([]int, 0)
    cell = append(cell, nums[0])
    for i := 1; i < n; i++ {
        if nums[i] > cell[len(cell)-1] {
            cell = append(cell, nums[i])
            continue
        }
        l, r := 0, len(cell)-1
        for l < r {
            mid := (l + r) / 2
            if cell[mid] < nums[i] {
                l = mid + 1
            } else {
                r = mid
            }
        }
        cell[l] = nums[i]
    }
    return len(cell)
}
func GetMax(nums ...int) int {
    var Max int
    Max = nums[0]

    for _, v := range nums {
        if v > Max {
            Max = v
        }
    }
    return Max
}



func main(){

    fmt.Println(lengthOfLIS([]int{10,9,2,5,3,7,101,18}))

}