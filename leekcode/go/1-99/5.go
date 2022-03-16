package __99

import "fmt"

// Manacher算法 在字符串有序时,会退化成O^2

/*
5. 最长回文子串

给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
*/

var count int

func GetString(s string) string {
    var str string

    str += `$`
    str += `#`

    for i := 0; i < len(s); i++ {
        str += string(s[i])
        str += `#`
    }
    str += `\0`
    return str
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

func longestPalindrome1(s string) string {
    if s == "" {
        return ""
    }

    str := GetString(s)
    p := make([]int, len(str)*2)
    var str2 string
    var maxLenth int
    var id, mx int
    for i := 1; i < len(str); i++ {
        count++
        if i < mx {
            p[i] = min(p[2*id-1], mx-1)
        } else {
            p[i] = 1
        }
        if i+p[i] < len(str)-1 {
            for str[i-p[i]] == str[i+p[i]] {
                count++
                p[i]++
            }
        }
        if mx < i+p[i] {
            id = 1
            mx = i + p[i]
        }
        if maxLenth < p[i]-1 {
            maxLenth = p[i] - 1
            array := str[i-p[i]+1 : i+p[i]]
            var str1 string
            for i := 0; i < len(array); i++ {
                count++
                if string(array[i]) != `#` && string(array[i]) != `$` {
                    str1 += string(array[i])
                }
            }
            if len(str1) > len(str2) {
                str2 = str1
            }
        }
    }
    return str2
}

func longestPalindrome(s string) string {
    if len(s) < 2 {
        return s
    }
    if len(s) > 1000 {
        panic("超出1000字符")
    }
    begin, maxLen := 0, 1
    for i := 0; i < len(s); {
        count++
        if len(s)-i <= maxLen/2 { // 当
            break
        }

        b, e := i, i
        for e < len(s)-1 && s[e+1] == s[e] {
            count++
            e++
        }
        i = e + 1
        for e < len(s)-1 && b > 0 && s[e+1] == s[b-1] {
            count++
            e++
            b--
        }
        newLen := e + 1 - b
        if newLen > maxLen {
            begin = b
            maxLen = newLen
        }
    }
    return s[begin : begin+maxLen]
}

func main() {
    str := `sadlaskldajskldjlqw;jdlwqdjas;dklasl;kdkl;asl;kdkl;asdklas'f'l;adsfkl;'adskl';fkl';asd'`
    fmt.Println(longestPalindrome(str))
    fmt.Println(count)

    count = 0
    fmt.Println(longestPalindrome1(str))
    fmt.Println(count)

}
