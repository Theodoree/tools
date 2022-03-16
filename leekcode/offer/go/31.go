package _go





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
剑指 Offer 31. 栈的压入、弹出序列
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈的压栈序列，
序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。
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
