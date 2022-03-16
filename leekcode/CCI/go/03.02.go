package _go


/*
面试题 03.02. 栈的最小值
请设计一个栈，除了常规栈支持的pop与push函数以外，还支持min函数，该函数返回栈元素中的最小值。执行push、pop和min操作的时间复杂度必须为O(1)。


示例：

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
	stack []int
	min   []int
}

func Constructor() MinStack {
	return MinStack{}
}
func (this *MinStack) empty() bool {
	return len(this.stack) == 0
}

func (this *MinStack) Push(x int) {
	this.push(x)
}

func (this *MinStack) push(x int) {
	this.stack = append(this.stack, x)
	if len(this.stack) == 1 {
		this.min = append(this.min, x)
		return
	}

	min := this.min[len(this.min)-1]
	if min < x {
		this.min = append(this.min, min)
		return
	}
	this.min = append(this.min, x)

}

func (this *MinStack) Pop() {
	if this.empty() {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	if this.empty() {
		return -1
	}
	return this.stack[len(this.stack)-1]

}

func (this *MinStack) GetMin() int {
	if this.empty() {
		return -1
	}
	return this.min[len(this.min)-1]
}
