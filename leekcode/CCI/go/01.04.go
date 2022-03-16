package _go


/*
面试题 01.04. 回文排列
给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。

回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。

回文串不一定是字典当中的单词。



示例1：

输入："tactcoa"
输出：true（排列有"tacocat"、"atcocta"，等等）
*/

func canPermutePalindrome(s string) bool {
    var arr = make([]int, 1 << 8, 1 << 8)
    isdouble := len(s)%2 == 0
    /* 回文串,字母的数量必须为二倍,或者s为单数,辣么应该有一个单数的字母*/

    for i := 0; i < len(s); i++ {
        arr[s[i]-'a']++
    }

    var flag bool
    for _, v := range arr {
        if isdouble && v%2 != 0 {
            return false
        }
        if !isdouble && v%2 != 0 {
            if flag {
                return false
            }
            flag = true
        }
    }

    return true
}
