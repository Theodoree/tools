package _go

/*
面试题 03.01. 三合一
三合一。描述如何只用一个数组来实现三个栈。

你应该实现push(stackNum, value)、pop(stackNum)、isEmpty(stackNum)、peek(stackNum)方法。stackNum表示栈下标，value表示压入的值。

构造函数会传入一个stackSize参数，代表每个栈的大小。

示例1:

 输入：
["TripleInOne", "push", "push", "pop", "pop", "pop", "isEmpty"]
[[1], [0, 1], [0, 2], [0], [0], [0], [0]]
 输出：
[null, null, null, 1, -1, -1, true]
说明：当栈为空时`pop, peek`返回-1，当栈满时`push`不压入元素。
示例2:

 输入：
["TripleInOne", "push", "push", "push", "pop", "pop", "pop", "peek"]
[[2], [0, 1], [0, 2], [0, 3], [0], [0], [0], [0]]
 输出：
[null, null, null, null, 2, 1, -1, -1]


提示：

0 <= stackNum <= 2
*/
type TripleInOne struct {
	stacks [][]int
	size int
}

func Constructor(stackSize int) TripleInOne {
	var t = TripleInOne{}
	t.size = stackSize
	for i := 0; i < 3; i++ {
		t.stacks = append(t.stacks, make([]int, 0, stackSize))
	}
	return t
}

func (this *TripleInOne) Push(stackNum int, value int) {
	if len(this.stacks[stackNum]) == this.size {
		return
	}
	this.stacks[stackNum] = append(this.stacks[stackNum], value)
}

func (this *TripleInOne) Pop(stackNum int) int {
	if len(this.stacks[stackNum]) == 0 {
		return -1
	}
	val := this.stacks[stackNum][len(this.stacks[stackNum])-1]
	this.stacks[stackNum] = this.stacks[stackNum][:len(this.stacks[stackNum])-1]
	return val
}

func (this *TripleInOne) Peek(stackNum int) int {
	if len(this.stacks[stackNum]) == 0 {
		return -1
	}
	return this.stacks[stackNum][len(this.stacks[stackNum])-1]
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return len(this.stacks[stackNum]) == 0
}
