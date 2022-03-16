package _00_899



/*
896. 单调数列

如果数组是单调递增或单调递减的，那么它是单调的。

如果对于所有 i <= j，A[i] <= A[j]，那么数组 A 是单调递增的。 如果对于所有 i <= j，A[i]> = A[j]，那么数组 A 是单调递减的。

当给定的数组 A 是单调数组时返回 true，否则返回 false。



示例 1：

输入：[1,2,2,3]
输出：true

示例 2：

输入：[6,5,4,4]
输出：true

示例 3：

输入：[1,3,2]
输出：false

示例 4：

输入：[1,2,4,5]
输出：true

示例 5：

输入：[1,1,1]
输出：true
*/

func isMonotonic(A []int) bool {

    switch {
    case len(A) <= 2:
        return true
    }
    var f int
    var IsDesc bool
    // i = 1 DESC i = 2 AES

    for i := 0; i < len(A)-1; i++ {

        if f == 0 {
            if A[i] > A[i+1] {
                if f == 0 || f == 1 {
                    f = 1
                }
            } else if A[i] < A[i+1]{
                if f == 0 || f == 2 {
                    f = 2
                }
            }

            if f != 0 {
                IsDesc = f == 1
            }
        }

        if IsDesc {
            if A[i] < A[i+1] {
                return false
            }
        } else {
            if A[i] > A[i+1] {
                return false
            }
        }

    }
    return true
}
