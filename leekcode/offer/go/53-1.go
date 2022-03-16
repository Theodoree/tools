package _go

/*
面试题53 - I. 在排序数组中查找数字 I

统计一个数字在排序数组中出现的次数。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: 2
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: 0

限制：

0 <= 数组长度 <= 50000


*/
func search(nums []int, target int) int {

    var cnt int

    for i:=0;i<len(nums);i++{
        if nums[i] == target {
            cnt++
        }
        if nums[i] > target {
            break
        }
    }
    return  cnt
}