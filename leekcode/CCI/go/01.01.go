package _go


/*
面试题 01.01. 判定字符是否唯一
实现一个算法，确定一个字符串 s 的所有字符是否全都不同。

示例 1：

输入: s = "leetcode"
输出: false
示例 2：

输入: s = "abc"
输出: true


按位异 不同为一  相同为
按位与 相同为1 不同为0
按位零 相同为0 不同为1

*/

func isUnique(astr string) bool {

    arr := make([]int, 26, 26)

    for i:=0;i<len(astr);i++{
        arr[astr[i]-'a']++
        if arr[astr[i]-'a'] > 1 {
            return  false
        }
    }

    return true

}
