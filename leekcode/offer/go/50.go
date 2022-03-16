package _go


/*
面试题50. 第一个只出现一次的字符

在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。

示例:

s = "abaccdeff"
返回 "b"

s = ""
返回 " "

*/
func firstUniqChar(s string) byte {
    hashMap := make(map[byte]int)
    for _, v := range []byte(s) {
        hashMap[v]++
    }

    for _,v:=range  []byte(s) {
        if hashMap[v] == 1 {
            return  v
        }
    }

    return ' '
}
