package _00_299



/*
238. 除自身以外数组的乘积

给定长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

示例:

输入: [1,2,3,4]
输出: [24,12,8,6]
说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。

进阶：
你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）
*/

func productExceptSelf(nums []int) []int {

    left := 1
    right := 1
    var result = make([]int, len(nums))

    for i := 0; i < len(result); i++ {
        result[i] = 1
    }

    for i := 0; i < len(nums); i++ {
        result[i] *= left
        left *= nums[i]

        result[len(nums)-1-i] *= right
        right *= nums[len(nums)-1-i]
    }

    return result
}
