package _00_799


/*
796. 旋转字符串

给定两个字符串, A 和 B。

A 的旋转操作就是将 A 最左边的字符移动到最右边。 例如, 若 A = 'abcde'，在移动一次之后结果就是'bcdea' 。如果在若干次旋转操作之后，A 能变成B，那么返回True。

示例 1:
输入: A = 'abcde', B = 'cdeab'
输出: true

示例 2:
输入: A = 'abcde', B = 'abced'
输出: false

*/

func rotateString(A string, B string) bool {
    filter := make(map[int]byte)

    if len(A) != len(B) {
        return false
    }

    for k, v := range A {
        filter[k] = byte(v)
    }

    key := -1
    var IsFirst = true
    for i := 0; i < len(B)-1; i++ {

        if key == - 1 || !(B[i] == filter[key]) {
            for k, v := range filter {
                if v == B[i] {
                    if IsFirst {
                        key = k
                        IsFirst = false
                    } else {
                        if B[i+1] == filter[k+1] {
                            key = k
                        }
                    }
                }
            }
        }

        if !(B[i] == filter[key]) {
            return false
        }
        key++

    }

    return true

}