package _00_799

import "container/heap"

/*
703. 数据流中的第 K 大元素
设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。

请实现 KthLargest 类：

KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。


示例：

输入：
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
输出：
[null, 4, 5, 5, 8, 8]

解释：
KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
kthLargest.add(3);   // return 4
kthLargest.add(5);   // return 5
kthLargest.add(10);  // return 5
kthLargest.add(9);   // return 8
kthLargest.add(4);   // return 8


提示：
1 <= k <= 104
0 <= nums.length <= 104
-104 <= nums[i] <= 104
-104 <= val <= 104
最多调用 add 方法 104 次
题目数据保证，在查找第 k 大元素时，数组中至少有 k 个元素
*/
type maxHeap struct {
	arr []int
}

func (s *maxHeap) Push(x interface{}) {
	s.arr = append(s.arr, x.(int))
}

func (s *maxHeap) Pop() interface{} {
	l := s.Len() - 1
	v := s.arr[l]
	s.arr = s.arr[:l]
	return v
}

func (s *maxHeap) Len() int {
	return len(s.arr)
}
func (s *maxHeap) Less(i, j int) bool {
	return s.arr[i] < s.arr[j]
}
func (s *maxHeap) Swap(i, j int) {
	s.arr[i], s.arr[j] = s.arr[j], s.arr[i]
}

// 最大栈
type KthLargest struct {
	s maxHeap
	K int
}

func Constructor(k int, nums []int) KthLargest {
	var kl KthLargest
	kl.s.arr = nums
	kl.K = k
	heap.Init(&kl.s)

	for kl.s.Len() > kl.K {
		heap.Pop(&kl.s)
	}

	return kl
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.s, val)
	if this.s.Len() > this.K {
		heap.Pop(&this.s)
	}
	return this.s.arr[0]
}



