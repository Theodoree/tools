package _000_1099

/*
1021. 删除最外层的括号

有效括号字符串为空 ("")、"(" + A + ")" 或 A + B，其中 A 和 B 都是有效的括号字符串，+ 代表字符串的连接。例如，""，"()"，"(())()" 和 "(()(()))" 都是有效的括号字符串。

如果有效字符串 S 非空，且不存在将其拆分为 S = A+B 的方法，我们称其为原语（primitive），其中 A 和 B 都是非空有效括号字符串。

给出一个非空有效字符串 S，考虑将其进行原语化分解，使得：S = P_1 + P_2 + ... + P_k，其中 P_i 是有效括号字符串原语。

对 S 进行原语化分解，删除分解中每个原语字符串的最外层括号，返回 S 。
*/

type stack struct {
    Val  string
    Next *stack
}

//简单栈类型
type MyStack struct {
    T     *stack
    Count int
}

func (s *MyStack) Pop() string {
    if s.T == nil {
        return ""
    }
    val := s.T.Val
    s.T = s.T.Next
    s.Count--
    return val
}

func (s *MyStack) Push(str string) {
    newStack := &stack{Val: str}
    if s.T == nil {
        s.T = newStack
        s.Count++
        return
    }

    newStack.Next = s.T
    s.T = newStack
    s.Count++
}

func removeOuterParentheses(S string) string {
    ms := MyStack{}
    var result string
    for i := 0; i < len(S); i++ {
        //                  ( () )
        switch S[i] {
        case '(':
            ms.Push("(")
            if ms.Count > 1 {
                result += "("
            }
        case ')':
            ms.Pop()

            if ms.Count >= 1 {
                result += ")"
            }
        }
    }

    return result
}