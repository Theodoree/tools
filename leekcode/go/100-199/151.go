package _00_199

import (
    "fmt"
    "strings"
)

func reverseWords(s string) string {
    var strs []string
    var s1 string
    s = strings.TrimSpace(s)
    for k, v := range s {
        s1 += string(v)
        if v == ' ' || k == len(s)-1 {
            s1 = strings.TrimSpace(s1)
            if len(s1) > 0 && s1[0] != ' '{
                strs = append(strs, s1)
            }
            s1 = ``
        }
    }

    var first, last int
    last = len(strs) - 1

    for first < last {
        strs[first], strs[last] = strs[last], strs[first]
        first++
        last--
    }
    return strings.Join(strs, ` `)

}

func main() {
    fmt.Println(reverseWords("a good   example"))
}
