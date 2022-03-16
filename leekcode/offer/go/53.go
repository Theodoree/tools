package _go


/*
面试题53 - II. 0～n-1中缺失的数字

一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。



示例 1:

输入: [0,1,3]
输出: 2
示例 2:

输入: [0,1,2,3,4,5,6,7,9]
输出: 8
*/
func missingNumber(nums []int) int {

    var cnt int
    var total int
    for i := 0; i < len(nums); i++ {
        cnt += nums[i]
        total += i + 1
    }

    return  total-cnt

}
