package _00_699


/*
680. 验证回文字符串 Ⅱ

给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。

示例 1:

输入: "aba"
输出: True
示例 2:

输入: "abca"
输出: True
解释: 你可以删除c字符。
注意:

字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。
*/
func validPalindrome(s string) bool {

    var left, right int

    sByte := []byte(s)
    right = len(s) - 1

    for left < right {
        if sByte[left] != sByte[right] {
            return ValidPalindrome(sByte, left, right-1) || ValidPalindrome(sByte, left+1, right)
        }
        left++
        right--
    }

    return true

}

func ValidPalindrome(chars []byte, l, r int) bool {
    for l < r {
        if chars[l] != chars[r] {
            return false;
        }
        l++
        r--
    }
    return true;
}
