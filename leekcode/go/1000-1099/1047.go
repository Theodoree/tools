package _000_1099


/*
1047. 删除字符串中的所有相邻重复项


给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。

在 S 上反复执行重复项删除操作，直到无法继续删除。

在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。



示例：

输入："abbaca"
输出："ca"
解释：
例如，在 "abbaca" 中，我们可以删除 "bb" 由于两字母相邻且相同，这是此时唯一可以执行删除操作的重复项。之后我们得到字符串 "aaca"，其中又只有 "aa" 可以执行重复项删除操作，所以最后的字符串为 "ca"。


提示：

1 <= S.length <= 20000
S 仅由小写英文字母组成。
*/

type Stack struct {
    Next *Stack
    val  byte
}

func removeDuplicates(S string) string {

    f := []byte(S)

    var s *Stack
    for i := 0; i < len(f); i++ {
        stack := &Stack{val: f[i]}
        if s != nil && s.val == stack.val {
            s = s.Next
        } else {
            stack.Next = s
            s = stack
        }
    }
    var result []byte

    for s != nil {
        result = append([]byte{s.val}, result...)
        s = s.Next
    }

    return string(result)
}

// 原地算法
func removeDuplicates1(S string) string {

    var l int
    var k = []rune(S)

    for _, v := range k {
        if l == 0 || k[l-1] != v {
            k[l] = v
            l++
        } else {
            l--
        }
    }
    k = k[:l]
    return string(k)
}
