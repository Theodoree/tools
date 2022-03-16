package _go

import "sort"

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


/*
面试题 17.14. 最小K个数
设计一个算法，找出数组中最小的k个数。以任意顺序返回这k个数均可。
示例：
*/
func smallestK(arr []int, k int) []int {
	//if k == 0 {
	//	return nil
	//}
	//var h IntHeap
	//h =make([]int,0,k)
	//
	//for i:=0;i<k;i++{
	//	heap.Push(&h,arr[i])
	//}
	//
	//for i:=k;i<len(arr);i++{
	//	if h[0] > arr[i]{
	//		heap.Pop(&h)
	//		heap.Push(&h,arr[i])
	//	}
	//}
	//return h
	sort.Ints(arr)
	return arr[:k]
}

