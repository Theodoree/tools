package _go


/*
面试题 17.20. 连续中值
随机产生数字并传递给一个方法。你能否完成这个方法，在每次产生新值时，寻找当前所有值的中间值（中位数）并保存。

中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
*/
type IntHeap struct {
	arr   []int
	IsMin bool
}

func (h IntHeap) Len() int { return len(h.arr) }
func (h IntHeap) Less(i, j int) bool {
	if h.IsMin {
		return h.arr[i] < h.arr[j]
	}
	return h.arr[i] > h.arr[j]
}
func (h IntHeap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]

}
func (h *IntHeap) Push(x interface{}) { h.arr = append(h.arr, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	n := len(h.arr) - 1
	x := (h.arr)[n]
	h.arr = (h.arr)[:n]
	return x
}

type MedianFinder struct {
	maxHeap IntHeap
	minHeap IntHeap
	len     int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	var s MedianFinder
	s.minHeap.IsMin = true
	return s
}

func (this *MedianFinder) AddNum(num int) {
	this.len++
	// max 上半部分  min 下半部分

	if this.maxHeap.Len() == 0 || num <= this.maxHeap.arr[0] {
		heap.Push(&this.maxHeap, num)
	} else {
		heap.Push(&this.minHeap, num)
	}
	if this.len%2 == 0 { // 这时候判断最大堆的值是否
		for this.maxHeap.Len() > this.minHeap.Len() { // 确保平衡
			heap.Push(&this.minHeap, heap.Pop(&this.maxHeap))
		}
		for this.minHeap.Len() > this.maxHeap.Len() {
			heap.Push(&this.maxHeap, heap.Pop(&this.minHeap))
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.len%2 == 0 {
		return float64(this.maxHeap.arr[0]+this.minHeap.arr[0]) / 2
	}
	if this.maxHeap.Len() > this.minHeap.Len() {
		return float64(this.maxHeap.arr[0])
	}
	return float64(this.minHeap.arr[0])
}
