package _00_399

import (
    "fmt"
)

/*
394. 字符串解码

给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

示例:

s = "3[a]2[bc]", 返回 "aaabcbc".
s = "3[a2[c]]", 返回 "accaccacc".
s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".
*/

type stack struct {
    Next  *stack
    times int
    Val   byte
    cur   []byte
}

func decodeString(s string) string {

    stack1 := &stack{times: -1}

    var result string
    var left int

    for left < len(s) {

        switch {
        case s[left] == ']':
            top := stack1
            stack1 = stack1.Next
            for i := 0; i < top.times; i++ {
                stack1.cur = append(stack1.cur,top.cur...)
            }
        case s[left] >= '0' && s[left] <= '9':
            var times int
            for s[left] >= '0' && s[left] <= '9' {
                times = times*10 + int(s[left]-'0')
                left++
            }
            s1 := &stack{
                times: times,
                Val:   s[left],
                Next:  stack1,
            }
            stack1 = s1

        default:
            stack1.cur = append(stack1.cur, s[left])
        }
        left++
    }

    result += string(stack1.cur)
    return result
}

func main() {

    fmt.Println(decodeString(`3[a]2[bc]`))

}
