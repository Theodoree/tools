package _00_299


/*
266. 回文排列

给定一个字符串，判断该字符串中是否可以通过重新排列组合，形成一个回文字符串。

示例 1：

输入: "code"
输出: false
示例 2：

输入: "aab"
输出: true
示例 3：

输入: "carerac"
输出: true
*/
func canPermutePalindrome(s string) bool {
    var m = make(map[int32]int)

    for _, v := range s {
        m[v]++
    }

    var isQiShu bool

    for _, v := range m {

        if v%2 == 1 {
            if isQiShu {
                return false
            }
            isQiShu = true
        }

    }

    return  true
}