package _00_299

//简单栈类型
type MyStack struct {
    T     *Stack
    Count int
}
type Stack struct {
    Next *Stack
    Val  int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
    return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {

    stack := &Stack{Val: x}
    if this.T == nil {
        this.T = stack
    } else {
        stack.Next = this.T
        this.T = stack
    }
    this.Count++
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    if this.T == nil {
        return 0
    }

    s := this.T
    if s.Next != nil {
        this.T = s.Next
    }

    this.Count--
    return s.Val

}

/** Get the top element. */
func (this *MyStack) Top() int {
    if this.Empty() {
        return 0
    }

    return this.T.Val
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    return this.Count == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
