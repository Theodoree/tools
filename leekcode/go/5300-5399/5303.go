package _300_5399

import (
    "strconv"
)

/*

5303. 解码字母到整数映射

给你一个字符串 s，它由数字（'0' - '9'）和 '#' 组成。我们希望按下述规则将 s 映射为一些小写英文字符：

字符（'a' - 'i'）分别用（'1' - '9'）表示。
字符（'j' - 'z'）分别用（'10#' - '26#'）表示。
返回映射之后形成的新字符串。

题目数据保证映射始终唯一。



示例 1：

输入：s = "10#11#12"
输出："jkab"
解释："j" -> "10#" , "k" -> "11#" , "a" -> "1" , "b" -> "2".
示例 2：

输入：s = "1326#"
输出："acz"
示例 3：

输入：s = "25#"
输出："y"
示例 4：

输入：s = "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#"
输出："abcdefghijklmnopqrstuvwxyz"
*/

func freqAlphabets(s string) string {
    
    var (
        m      = make(map[string]byte)
        arr    []string
        result []byte
    )
    for i := 0; i < 26; i++ {
        
        key := strconv.Itoa(i + 1)
        if i >= 9 {
            key = "#" + key
        }
        m[key] = byte('a' + i)
        arr = append(arr, key)
    }
    
    for i := 0; i < len(s); i++ {
        if i+2 < len(s) && s[i+2] == '#' {
            result = append(result, m["#"+s[i:i+2]])
            i += 2
        } else {
            result = append(result, m[string(s[i])])
        }
        
    }
    
    return string(result)
}
