package _00_899

import (
    "fmt"
    "strings"
)

func mostCommonWord(paragraph string, banned []string) string {

    paragraph = strings.ToLower(paragraph)
    var tem string
    for _, v := range paragraph {
        if v >= 'a' && v <= 'z' {
            tem += string(v)
        } else {
            tem += ` `
        }
    }

    paragraph = tem
    strs := strings.Split(paragraph, ` `)
    resultMap := map[string]int{}

    for _, v := range banned {
        resultMap[v] = -1
    }

    var times, cnt int
    var word string

    for _, v := range strs {
        cnt = resultMap[v]
        if cnt == -1 || v == `` || v == ` `{
            continue
        }
        resultMap[v]++
        if times < cnt+1 {
            times = cnt + 1
            word = v
        }
    }

    return word

}

func main() {

    //fmt.Println(mostCommonWord("a, a, a, a, b,b,b,c, c", []string{"a"}))
    fmt.Println(mostCommonWord("Bob. hIt, baLl", []string{"bob","hit"}))
    //fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.",[]string{"hit"}))
}
