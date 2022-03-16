package _go


type MinStack struct {
    stack *stack
}

type stack struct {
    curMin int
    Next   *stack
    value  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{}
}

func NewStack(val int) *stack {
    return &stack{
        curMin: val,
        value:  val,
    }
}

func (this *MinStack) Push(x int) {
    stack := NewStack(x)
    if this.stack == nil {
        this.stack = stack
        return
    }

    if this.stack.curMin < stack.curMin {
        stack.curMin = this.stack.curMin
    }
    stack.Next = this.stack
    this.stack = stack
}

func (this *MinStack) Pop() {
    if this.stack == nil {
        return
    }
    this.stack = this.stack.Next
}

func (this *MinStack) Top() int {
    if this.stack == nil {
        return 0
    }
    return this.stack.value
}

func (this *MinStack) Min() int {
    if this.stack == nil {
        return 0
    }
    return this.stack.curMin
}

