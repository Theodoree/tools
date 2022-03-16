package _00_599

import "strings"

/*
500. 键盘行

给定一个单词列表，只返回可以使用在键盘同一行的字母打印出来的单词。键盘如下图所示。
示例：

输入: ["Hello", "Alaska", "Dad", "Peace"]
输出: ["Alaska", "Dad"]
*/

func findWords(words []string) []string {


    var result []string
    var flag bool
    var t1 int
    for i := 0; i < len(words); i++ {
        t1 = 0
        flag = false
        str:= strings.ToLower(words[i])
        for j := 0; j < len(str); j++ {
            if t1 == 0 {
                t1 = f[str[j]]
                continue
            }
            if f[str[j]] != t1 {
                flag = true
                break
            }
        }
        if !flag {
            result = append(result,words[i])
        }
    }




    return result

}

var f = map[byte]int{
    'q': 1,
    'w': 1,
    'e': 1,
    'r': 1,
    't': 1,
    'y': 1,
    'u': 1,
    'i': 1,
    'o': 1,
    'p': 1,
    'a': 2,
    's': 2,
    'd': 2,
    'f': 2,
    'g': 2,
    'h': 2,
    'j': 2,
    'k': 2,
    'l': 2,
    'z': 3,
    'x': 3,
    'c': 3,
    'v': 3,
    'b': 3,
    'n': 3,
    'm': 3,
}
