package _00_999


/*
989. 数组形式的整数加法

对于非负整数 X 而言，X 的数组形式是每位数字按从左到右的顺序形成的数组。例如，如果 X = 1231，那么其数组形式为 [1,2,3,1]。

给定非负整数 X 的数组形式 A，返回整数 X+K 的数组形式。
*/

func addToArrayForm(A []int, K int) []int {

    right := len(A) - 1
    var cnt, result, y int
    for K > 0 || cnt == 1 {
        y = K % 10
        K /= 10
        result = A[right] + y + cnt
        cnt = 0
        if result >= 10 {
            cnt = 1
        }
        A[right] = result % 10
        if right > 0 {
            right--
        } else {
            if K > 0 || cnt == 1 {
                A = append([]int{0}, A...)
            }
        }
    }

    return A
}
