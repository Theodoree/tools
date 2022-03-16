package _00_299

import "strings"

/*
290. 单词规律

给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。


*/
func wordPattern(pattern string, str string) bool {

    str_arr := strings.Fields(str)
    if len(str_arr) != len(pattern) {
        return false
    }
    hash := make(map[byte]string)
    hash2 := make(map[string]byte)
    for i := 0; i < len(pattern); i++ {
        v, ok := hash[pattern[i]]
        v2, ok2 := hash2[str_arr[i]]
        if ok && v != str_arr[i] || ok2 && v2 != pattern[i] {
            return false
        } else {
            hash[pattern[i]] = str_arr[i]
            hash2[str_arr[i]] = pattern[i]
        }
    }
    return true
}
