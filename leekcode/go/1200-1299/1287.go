package _200_1299


/*
1287. 有序数组中出现次数超过25%的元素

给你一个非递减的 有序 整数数组，已知这个数组中恰好有一个整数，它的出现次数超过数组元素总数的 25%。

请你找到并返回这个整数



示例：

输入：arr = [1,2,2,6,6,6,6,7,10]
输出：6


提示：

1 <= arr.length <= 10^4
0 <= arr[i] <= 10^5
*/

func findSpecialInteger(arr []int) int {
    length := len(arr)
    count := 1
    for i := 1; i < length; i++ {
        if arr[i] == arr[i-1] {
            count++
            if count > length / 4 {
                return arr[i]
            }
        } else {
            count = 1
        }
    }
    return arr[0]
}
