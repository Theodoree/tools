package _100_1199

import (
    "container/heap"
    "fmt"
)

/*
1172. 餐盘栈
我们把无限数量 ∞ 的栈排成一行，按从左到右的次序从 0 开始编号。每个栈的的最大容量 capacity 都相同。

实现一个叫「餐盘」的类 DinnerPlates：

DinnerPlates(int capacity) - 给出栈的最大容量 capacity。
void push(int val) - 将给出的正整数 val 推入 从左往右第一个 没有满的栈。
int pop() - 返回 从右往左第一个 非空栈顶部的值，并将其从栈中删除；如果所有的栈都是空的，请返回 -1。
int popAtStack(int index) - 返回编号 index 的栈顶部的值，并将其从栈中删除；如果编号 index 的栈是空的，请返回 -1。

*/

type DinnerPlates struct {
    cap     int
    stack   []*stack
    minHeap *Heap
    maxHeap *Heap
}

type Heap struct {
    IsMin bool  // true 小顶堆   false 大顶堆
    arr   []int // 里面存的是下标
    m     map[int]struct{}
}

func (h *Heap) Less(i, j int) bool {
    if h.IsMin {
        return h.arr[i] < h.arr[j]
    }
    return h.arr[i] > h.arr[j]
}

func (h *Heap) first() int {
    return h.arr[0]
}

func (h *Heap) Swap(i, j int) {
    h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *Heap) Len() int {
    return len(h.arr)
}

func (h *Heap) Push(x interface{}) {
    h.m[x.(int)] = struct{}{}
    h.arr = append(h.arr, x.(int))
}

func (h *Heap) Pop() interface{} {
    el := h.arr[len(h.arr)-1]
    h.arr = h.arr[:len(h.arr)-1]
    delete(h.m, el)
    return el
}

func (h *Heap) exist(idx int) bool {
    _, ok := h.m[idx]
    return ok
}

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

func Constructor(capacity int) DinnerPlates {
    minHeap := &Heap{
        m:     make(map[int]struct{}),
        IsMin: true,
    }
    maxHeap := &Heap{
        m: make(map[int]struct{}),
    }

    heap.Init(minHeap)
    heap.Init(maxHeap)

    return DinnerPlates{
        cap:     capacity,
        stack:   nil,
        minHeap: minHeap,
        maxHeap: maxHeap,
    }

}

func (this *DinnerPlates) Push(val int) {

    if this.minHeap.Len() > 0 {
        idx := this.minHeap.first()
        stack := this.stack[idx]
        stack.push(val)

        // 如果已满,那么就出栈
        if stack.len() == this.cap {
            heap.Pop(this.minHeap)
        }

        if !this.maxHeap.exist(idx) {
            heap.Push(this.maxHeap, idx)
        }
        return
    }

    stack := NewStack(this.cap)
    stack.push(val)
    this.stack = append(this.stack, stack)
    l := len(this.stack) - 1

    // 入栈 最大堆
    heap.Push(this.maxHeap, l)
    if stack.len() != this.cap {
        // 入栈 最小堆
        heap.Push(this.minHeap, l)
    }

}

func (this *DinnerPlates) print() {
    for i := 0; i < len(this.stack); i++ {
        fmt.Println(this.stack[i])
    }
    fmt.Println()
}

func (this *DinnerPlates) Pop() int {

    for this.maxHeap.Len() > 0 && this.stack[this.maxHeap.first()].len() == 0 {
        heap.Pop(this.maxHeap)
    }

    if this.maxHeap.Len() > 0 {
        idx := this.maxHeap.first()
        stack := this.stack[idx]
        val := stack.pop()

        if !this.minHeap.exist(idx) {
            heap.Push(this.minHeap, idx)
        }

        return val

    }

    return -1

}

func (this *DinnerPlates) PopAtStack(index int) int {
    if index >= len(this.stack) {
        return -1
    }

    stack := this.stack[index]
    val := stack.pop()

    if !this.minHeap.exist(index) {
        heap.Push(this.minHeap, index)
    }

    return val
}
