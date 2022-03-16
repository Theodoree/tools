package _00_999


/*

905. 按奇偶排序数组

给定一个非负整数数组 A，返回一个数组，在该数组中， A 的所有偶数元素之后跟着所有奇数元素。

你可以返回满足此条件的任何数组作为答案。
*/
func sortArrayByParity(A []int) []int {
    var left, right int
    right = len(A) - 1
    for left < right {

        for right >= 0 && A[right]%2 == 1 {
            right--
        }

        for left < right && A[left]%2 == 0 {
            left++
        }

        if left <right{
            A[left],A[right] = A[right],A[left]
        }

    }
    return A
}