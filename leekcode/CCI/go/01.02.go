package _go

/*
面试题 01.02. 判定是否互为字符重排
给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。

示例 1：

输入: s1 = "abc", s2 = "bca"
输出: true
示例 2：

输入: s1 = "abc", s2 = "bad"
输出: false
说明：

0 <= len(s1) <= 100
0 <= len(s2) <= 100
*/

func CheckPermutation(s1 string, s2 string) bool {

    var  arr = make([]int,26,26)
    for i:=0;i<len(s1);i++{
        arr[s1[i]-'a']++
    }

    for i:=0;i<len(s2);i++{
        arr[s2[i]-'a']--
        if  arr[s2[i]-'a'] < 0 {
            return  false
        }
    }

    return true
}
