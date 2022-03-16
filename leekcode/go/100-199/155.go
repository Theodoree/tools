package _00_199



/*
155. 最小栈

设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) -- 将元素 x 推入栈中。
pop() -- 删除栈顶的元素。
top() -- 获取栈顶元素。
getMin() -- 检索栈中的最小元素。
示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
*/

type MinStack struct {
    top *Stack
    min int
}

type Stack struct {
    next *Stack
    Val  int
    min  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{}
}

func (this *MinStack) Push(x int) {
    if this.min > x || this.top == nil {
        this.min = x
    }
    stack := &Stack{Val: x, min: this.min}
    stack.next = this.top
    this.top = stack

}

func (this *MinStack) Pop() {
    if this.top == nil {
        return
    }
    this.top = this.top.next
    if this.top != nil {
        this.min = this.top.min
    }
}

func (this *MinStack) Top() int {
    if this.top == nil {
        return 0
    }
    return this.top.Val
}

func (this *MinStack) GetMin() int {
    if this.top == nil {
        return 0
    }

    return this.min

}
