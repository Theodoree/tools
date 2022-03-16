package _00_599


/*
541. 反转字符串 II

给定一个字符串和一个整数 k，你需要对从字符串开头算起的每个 2k 个字符的前k个字符进行反转。如果剩余少于 k 个字符，则将剩余的所有全部反转。如果有小于 2k 但大于或等于 k 个字符，则反转前 k 个字符，并将剩余的字符保持原样。

示例:

输入: s = "abcdefg", k = 2
输出: "bacdfeg"
要求:

该字符串只包含小写的英文字母。
给定字符串的长度和 k 在[1, 10000]范围内。
*/

func reverseStr(s string, k int) string {

    var l, cnt, left, right int
    b := []byte(s)
    for l < len(s)-1 {

        left = l
        right = l + (k -1)
        cnt = k -1
        l +=k+k
        if right > len(s)-1 {
            right = len(s) - 1
        }


        for cnt > 0 && left < right{

            b[left], b[right] = b[right], b[left]
            left++
            right--
            cnt--
        }

    }

    return string(b)
}