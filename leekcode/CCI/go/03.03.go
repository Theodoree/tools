package _go


/*
面试题 03.03. 堆盘子
堆盘子。设想有一堆盘子，堆太高可能会倒下来。因此，在现实生活中，盘子堆到一定高度时，我们就会另外堆一堆盘子。请实现数据结构SetOfStacks，模拟这种行为。SetOfStacks应该由多个栈组成，并且在前一个栈填满时新建一个栈。此外，SetOfStacks.push()和SetOfStacks.pop()应该与普通栈的操作方法相同（也就是说，pop()返回的值，应该跟只有一个栈时的情况一样）。 进阶：实现一个popAt(int index)方法，根据指定的子栈，执行pop操作。

当某个栈为空时，应当删除该栈。当栈中没有元素或不存在该栈时，pop，popAt 应返回 -1.

示例1:

 输入：
["StackOfPlates", "push", "push", "popAt", "pop", "pop"]
[[1], [1], [2], [1], [], []]
 输出：
[null, null, null, 2, 1, -1]
示例2:

 输入：
["StackOfPlates", "push", "push", "push", "popAt", "popAt", "popAt"]
[[2], [1], [2], [3], [0], [0], [0]]
 输出：
[null, null, null, null, 2, 1, 3]
通过次数8,050提交次数20,873
*/

type stack []int

func (s *stack) len() int {
	return len(*s)
}

func (s *stack) pop() int {
	l := s.len()
	if l == 0 {
		return -1
	}
	el := (*s)[l-1]
	*s = (*s)[:l-1]
	return el
}

func (s *stack) push(i int) {
	*s = append(*s, i)
}

func NewStack(cap int) *stack {
	s := stack{}
	s = make([]int, 0, cap)
	return &s
}

type StackOfPlates struct {
	stack []*stack
	cap   int
}

func Constructor(cap int) StackOfPlates {
	return StackOfPlates{
		cap: cap,
	}
}

func (plates *StackOfPlates) Push(val int) {
	if plates.cap == 0 {
		return
	}
	if len(plates.stack) > 0 && plates.stack[len(plates.stack)-1].len() < plates.cap {
		plates.stack[len(plates.stack)-1].push(val)
		return
	}
	s := NewStack(plates.cap)
	plates.stack = append(plates.stack, s)
	s.push(val)
}

func (plates *StackOfPlates) Pop() int {
	return plates.PopAt(len(plates.stack) - 1)
}

func (plates *StackOfPlates) PopAt(index int) int {
	if plates.cap == 0 {
		return - 1
	}
	if len(plates.stack) > 0 && index < len(plates.stack) {
		s := plates.stack[index]
		val := s.pop()
		if s.len() == 0 {
			if index == len(plates.stack)-1 {
				plates.stack = plates.stack[:index]
			} else {
				plates.stack = append(plates.stack[:index], plates.stack[index+1:]...)
			}
		}
		return val
	}

	return -1
}

