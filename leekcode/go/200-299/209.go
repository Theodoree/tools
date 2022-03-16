package _00_299

/*
209. 长度最小的子数组

给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。如果不存在符合条件的连续子数组，返回 0。

示例:

输入: s = 7, nums = [2,3,1,2,4,3]
输出: 2
解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
进阶:

如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
*/

func min(a, b int) int {
    if a > b {
        return b
    }

    return a
}

func minSubArrayLen(s int, nums []int) int {

    ans := math.MaxInt64
    var left, sum int

    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        for sum >= s {
            ans = min(ans, i+1-left)
            sum -= nums[left]
            left++
        }
    }
    if ans == math.MaxInt64 {
        return 0
    }
    return ans
}
