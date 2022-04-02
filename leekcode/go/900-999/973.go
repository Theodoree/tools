package _00_999



/*
973. 最接近原点的 K 个点
给定一个数组 points ，其中 points[i] = [xi, yi] 表示 X-Y 平面上的一个点，并且是一个整数 k ，返回离原点 (0,0) 最近的 k 个点。

这里，平面上两点之间的距离是 欧几里德距离（ √(x1 - x2)2 + (y1 - y2)2 ）。

你可以按 任何顺序 返回答案。除了点坐标的顺序之外，答案 确保 是 唯一 的。
*/
type IntHeap struct {
	arr   [][2]int
	IsMin bool
}

func Sqrt(i int) float64 {
	return math.Sqrt(float64(i))
}

func (h IntHeap) Len() int { return len(h.arr) }
func (h IntHeap) Less(i, j int) bool {
	return Sqrt(h.arr[i][0]*h.arr[i][0]+h.arr[i][1]*h.arr[i][1]) > Sqrt(h.arr[j][0]*h.arr[j][0]+h.arr[j][1]*h.arr[j][1])
}
func (h IntHeap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]

}
func (h *IntHeap) Push(x interface{}) { h.arr = append(h.arr, x.([2]int)) }
func (h *IntHeap) Pop() interface{} {
	n := len(h.arr) - 1
	x := (h.arr)[n]
	h.arr = (h.arr)[:n]
	return x
}

func (h *IntHeap) fronend() float64 {
	return Sqrt(h.arr[0][0]*h.arr[0][0] + h.arr[0][1]*h.arr[0][1])
}

/*
973. 最接近原点的 K 个点
给定一个数组 points ，其中 points[i] = [xi, yi] 表示 X-Y 平面上的一个点，并且是一个整数 k ，返回离原点 (0,0) 最近的 k 个点。
这里，平面上两点之间的距离是 欧几里德距离（ √(x1 - x2)2 + (y1 - y2)2 ）。
你可以按 任何顺序 返回答案。除了点坐标的顺序之外，答案 确保 是 唯一 的。
*/

func kClosest(points [][]int, k int) [][]int {
	var h IntHeap
	for i := 0; i < len(points); i++ {
		n := Sqrt(points[i][0]*points[i][0] + points[i][1]*points[i][1])
		if h.Len() == k && h.fronend() > n {
			heap.Pop(&h)
			heap.Push(&h, [2]int{points[i][0], points[i][1]})
		} else if h.Len() < k {
			heap.Push(&h, [2]int{points[i][0], points[i][1]})
		}
	}

	var result [][]int
	for h.Len() > 0 {
		n := h.Pop().([2]int)
		result = append(result, n[:])
	}

	return result
}
