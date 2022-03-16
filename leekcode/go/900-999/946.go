package _00_999



type IntStack []int

func (h IntStack) Len() int            { return len(h) }
func (h *IntStack) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntStack) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/*
946. 验证栈序列
给定 pushed 和 popped 两个序列，每个序列中的 值都不重复，只有当它们可能是在最初空栈上进行的推入 push 和弹出 pop 操作序列的结果时，返回 true；否则，返回 false 。
*/
func validateStackSequences(pushed []int, popped []int) bool {
	var p IntStack
	for i:=0;i<len(pushed);i++{
		p.Push(pushed[i])
		for p.Len() >0 &&  popped[0] == p[p.Len()-1]{
			p.Pop()
			popped = popped[1:]
		}
	}

	return len(popped) == 0
}
