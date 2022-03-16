package _go

/*
剑指 Offer 09. 用两个栈实现队列

用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )



示例 1：

输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]
示例 2：

输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]
提示：

1 <= values <= 10000
最多会对 appendTail、deleteHead 进行 10000 次调用
*/

type CQueue struct {
	head *stack
	tail *stack
}

type stack struct {
	val  int
	Next *stack
}

func Constructor() CQueue {
	return CQueue{}

}
func NewStack(value int) *stack {
	return &stack{
		val:  value,
		Next: nil,
	}
}

func (this *CQueue) AppendTail(value int) {
	node := NewStack(value)
	if this.head == nil {
		this.head = node
		this.tail = node
		return
	}

	this.tail.Next = node
	this.tail = node
}

func (this *CQueue) DeleteHead() int {
	if this.head == nil {
		return -1
	}

	val := this.head.val
	if this.head == this.tail {
		this.head = nil
		this.tail = nil
		return val
	}

	this.head = this.head.Next
	return val
}
