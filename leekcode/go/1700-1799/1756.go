package _700_1799




/*
1756. 设计最近使用（MRU）队列
设计一种类似队列的数据结构，该数据结构将最近使用的元素移到队列尾部。
实现 MRUQueue 类：
MRUQueue(int n)  使用 n 个元素： [1,2,3,...,n] 构造 MRUQueue 。
fetch(int k) 将第 k 个元素（从 1 开始索引）移到队尾，并返回该元素。
*/

type MRUQueue struct {
	arr []uint16
}


func Constructor(n int) MRUQueue {
	var m MRUQueue
	var i uint16
	for i=0;i<=uint16(n);i++{
		m .arr = append(m.arr,i)
	}
	return m
}


func (this *MRUQueue) Fetch(k int) int {
	n:=this.arr[k]
	if k != len(this.arr)-1 {
		copy(this.arr[k:], this.arr[k+1:])
		this.arr[len(this.arr)-1] = n
	}
	return int(n)
}

