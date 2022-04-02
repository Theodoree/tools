package _00_799


type IntHeap struct {
	arr   [][2]float64
	IsMin bool
}

func (h IntHeap) Len() int { return len(h.arr) }
func (h IntHeap) Less(i, j int) bool {
	if h.IsMin {
		return h.arr[i][0]/h.arr[i][1] < h.arr[j][0]/h.arr[j][1]
	}
	return h.arr[i][0]/h.arr[i][1] > h.arr[j][0]/h.arr[j][1]
}
func (h IntHeap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]

}
func (h *IntHeap) Push(x interface{}) { h.arr = append(h.arr, x.([2]float64)) }
func (h *IntHeap) Pop() interface{} {
	n := len(h.arr) - 1
	x := (h.arr)[n]
	h.arr = (h.arr)[:n]
	return x
}

func (h *IntHeap) frontend() float64 {
	return h.arr[0][0] / h.arr[0][1]
}

/*
786. 第 K 个最小的素数分数
给你一个按递增顺序排序的数组 arr 和一个整数 k 。数组 arr 由 1 和若干 素数  组成，且其中所有整数互不相同。
对于每对满足 0 <= i < j < arr.length 的 i 和 j ，可以得到分数 arr[i] / arr[j] 。
那么第 k 个最小的分数是多少呢?  以长度为 2 的整数数组返回你的答案, 这里 answer[0] == arr[i] 且 answer[1] == arr[j] 。
*/

func kthSmallestPrimeFraction(arr []int, k int) []int {

	var h IntHeap

	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > i; j-- {
			if h.Len() == k && h.frontend() > float64(arr[i])/float64(arr[j]) {
				heap.Pop(&h)
				heap.Push(&h, [2]float64{float64(arr[i]), float64(arr[j])})
			} else if h.Len() < k {
				heap.Push(&h, [2]float64{float64(arr[i]), float64(arr[j])})
			}
		}
	}

	return []int{int(h.arr[0][0]), int(h.arr[0][1])}
}

