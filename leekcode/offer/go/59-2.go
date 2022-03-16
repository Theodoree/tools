package _go

import "container/heap"

/*
剑指 Offer 59 - II. 队列的最大值
请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。

若队列为空，pop_front 和 max_value 需要返回 -1

示例 1：

输入:
["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
[[],[1],[2],[],[],[]]
输出: [null,null,null,2,1,2]
示例 2：

输入:
["MaxQueue","pop_front","max_value"]
[[],[],[]]
输出: [null,-1,-1]
*/

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	n := len(*h) - 1
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}

type MaxQueue struct {
	h     IntHeap
	front *node
	tail  *node
	m     map[int]int
}

type node struct {
	val  int
	next *node
}

func Constructor() MaxQueue {

	sentinel := &node{}
	return MaxQueue{m: make(map[int]int), front: sentinel, tail: sentinel}
}

func (m *MaxQueue) Max_value() int {
	if m.h.Len() == 0 {
		return -1
	}

	for m.h.Len() > 0 && m.m[m.h[0]] == 0 {
		heap.Pop(&m.h)
	}
	if m.h.Len() == 0 {
		return -1
	}
	return m.h[0]
}

func (m *MaxQueue) Push_back(value int) {
	if val := m.m[value]; val == 0 {
		heap.Push(&m.h, value)
	}
	m.m[value]++
	n := &node{val: value}
	m.tail.next = n
	m.tail = n
}

func (m *MaxQueue) Pop_front() int {
	if m.front.next == nil {
		return -1
	}

	n := m.front.next
	m.front.next = n.next
	if n == m.tail {
		m.tail = m.front
	}

	m.m[n.val]--
	return n.val
}
