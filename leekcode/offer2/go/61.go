package _go


/*
剑指 Offer II 061. 和最小的 k 个数对
给定两个以升序排列的整数数组 nums1 和 nums2 , 以及一个整数 k 。

定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。

请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。
*/
type h struct {
	num1 []int
	num2 []int
	arr  []int
}

func (h *h) Len() int {
	return len(h.arr)
}

func (h *h) Less(i, j int) bool {
	iLeft := h.arr[i] & 0xffffffff
	iright := h.arr[i] >> 32

	jLeft := h.arr[j] & 0xffffffff
	jright := h.arr[j] >> 32

	return h.num1[iLeft]+h.num2[iright] < h.num1[jLeft]+h.num2[jright]
}
func (h *h) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *h) Push(x interface{}) {
	h.arr = append(h.arr, x.(int))
}
func (h *h) Pop() interface{} {
	l := h.Len()
	result := h.arr[l-1]
	h.arr = h.arr[:l-1]
	return result
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {

	_heap := h{num1: nums1, num2: nums2, arr: nil}
	for j := 0; j < len(nums2); j++ {
		heap.Push(&_heap, 0|j<<32)
	}

	var result [][]int
	for k > 0 && _heap.Len() > 0 {
		item := heap.Pop(&_heap).(int)

		left := item & 0xffffffff
		right := item >> 32

		result = append(result, []int{nums1[left], nums2[right]})
		k--
		if left+1 < len(nums1) {
			heap.Push(&_heap, left+1|right<<32)
		}
	}

	return result
}

