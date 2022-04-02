package _000_2099


/*
2034. 股票价格波动
给你一支股票价格的数据流。数据流中每一条记录包含一个 时间戳 和该时间点股票对应的 价格 。
不巧的是，由于股票市场内在的波动性，股票价格记录可能不是按时间顺序到来的。某些情况下，有的记录可能是错的。如果两个有相同时间戳的记录出现在数据流中，前一条记录视为错误记录，后出现的记录 更正 前一条错误的记录。
*/
type IntHeap struct {
	arr   []*[3]int
	IsMin bool
}

func (h IntHeap) Len() int { return len(h.arr) }
func (h IntHeap) Less(i, j int) bool {
	if h.IsMin {
		return h.arr[i][0] < h.arr[j][0]
	}
	return h.arr[i][0] > h.arr[j][0]
}
func (h IntHeap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]

	if !h.IsMin {
		h.arr[i][1] = i
		h.arr[j][1] = j
	} else {
		h.arr[i][2] = i
		h.arr[j][2] = j
	}
}
func (h *IntHeap) Push(x interface{}) { h.arr = append(h.arr, x.(*[3]int)) }
func (h *IntHeap) Pop() interface{} {
	n := len(h.arr) - 1
	x := (h.arr)[n]
	h.arr = (h.arr)[:n]
	return x
}

type StockPrice struct {
	maxHeap IntHeap
	minHeap IntHeap
	m       map[int]*[3]int
	cur     int
	curTs   int
}

func Constructor() StockPrice {
	s := StockPrice{m: make(map[int]*[3]int)}
	s.minHeap.IsMin = true
	return s
}

func (this *StockPrice) Update(timestamp int, price int) {
	if timestamp >= this.curTs {
		this.cur = price
		this.curTs = timestamp
	}

	if val, ok := this.m[timestamp]; ok {
		val[0] = price
		heap.Fix(&this.maxHeap, val[1])
		heap.Fix(&this.minHeap, val[2])
		return
	} else {
		n := [3]int{price, this.maxHeap.Len(), this.minHeap.Len()}
		heap.Push(&this.maxHeap, &n)
		heap.Push(&this.minHeap, &n)
		this.m[timestamp] = &n
	}

}

func (this *StockPrice) Current() int {
	return this.cur
}

func (this *StockPrice) Maximum() int {
	if this.maxHeap.Len() == 0 {
		return 0
	}
	return this.maxHeap.arr[0][0]
}

func (this *StockPrice) Minimum() int {
	if this.minHeap.Len() == 0 {
		return 0
	}
	return this.minHeap.arr[0][0]
}

