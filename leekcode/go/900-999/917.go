package _00_999



/*
917. 仅仅反转字母

给定一个字符串 S，返回 “反转后的” 字符串，其中不是字母的字符都保留在原地，而所有字母的位置发生反转。

示例 1：

输入："ab-cd"
输出："dc-ba"
示例 2：

输入："a-bC-dEf-ghIj"
输出："j-Ih-gfE-dCba"
示例 3：

输入："Test1ng-Leet=code-Q!"
输出："Qedo1ct-eeLg=ntse-T!"
*/

func reverseOnlyLetters(S string) string {

    var left, right int
    right = len(S) - 1
    f := []byte(strings.ToLower(S))
    Sbyte := []byte(S)
    for left < right {

        for left < right {
            if f[left] >= 'a' && f [left] <= 'z' {
                break
            }
            left++
        }

        for right >= 0 {
            if f[right] >= 'a' && f [right] <= 'z' {
                break
            }
            right--
        }
        if left > right {
            break
        }

        Sbyte[left], Sbyte[right] = Sbyte[right], Sbyte[left]

        left++
        right--
    }

    return string(Sbyte)
}
