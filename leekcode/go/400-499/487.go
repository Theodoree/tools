package _00_499



/*
487. 最大连续1的个数 II
给定一个二进制数组，你可以最多将 1 个 0 翻转为 1，找出其中最大连续 1 的个数。

示例 1：

输入：[1,0,1,1,0]
输出：4
解释：翻转第一个 0 可以得到最长的连续 1。
     当翻转以后，最大连续 1 的个数为 4。


注：

输入数组只包含 0 和 1.
输入数组的长度为正整数，且不超过 10,000


进阶：
如果输入的数字是作为 无限流 逐个输入如何处理？换句话说，内存不能存储下所有从流中输入的数字。您可以有效地解决吗？

1,0,1,1,0
*/

func findMaxConsecutiveOnes(nums []int) int {
    
    var (
        left     int
        count    int
        cur      int
        zeroIndex int
        flag     bool
    )
    
    for left < len(nums) {
        if nums[left] == 0 {
            if flag {
                count = max(count, cur)
                for zeroIndex+1 < len(nums) && nums[zeroIndex+1] == 0 {
                    zeroIndex++
                }
                if zeroIndex+1 >= len(nums) {
                    break
                }
                left = zeroIndex + 1
                cur = 0
                flag = false
            } else {
                zeroIndex = left
                flag = true
            }
        }
        left++
        cur++
        
    }
    
    if cur > count {
        count = cur
    }
    
    return count
    
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
