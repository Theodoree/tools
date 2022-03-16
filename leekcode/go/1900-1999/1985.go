package _900_1999

import (
	"container/heap"
	"sort"
)

/*
1985. 找出数组中的第 K 大整数
给你一个字符串数组 nums 和一个整数 k 。nums 中的每个字符串都表示一个不含前导零的整数。

返回 nums 中表示第 k 大整数的字符串。

注意：重复的数字在统计时会视为不同元素考虑。例如，如果 nums 是 ["1","2","2"]，那么 "2" 是最大的整数，"2" 是第二大的整数，"1" 是第三大的整数。



示例 1：

输入：nums = ["3","6","7","10"], k = 4
输出："3"
解释：
nums 中的数字按非递减顺序排列为 ["3","6","7","10"]
其中第 4 大整数是 "3"
示例 2：

输入：nums = ["2","21","12","1"], k = 3
输出："2"
解释：
nums 中的数字按非递减顺序排列为 ["1","2","12","21"]
其中第 3 大整数是 "2"
示例 3：

输入：nums = ["0","0"], k = 2
输出："0"
解释：
nums 中的数字按非递减顺序排列为 ["0","0"]
其中第 2 大整数是 "0"


提示：

1 <= k <= nums.length <= 104
1 <= nums[i].length <= 100
nums[i] 仅由数字组成
nums[i] 不含任何前导零
通过次数5,829提交次数14,160

*/

type StringHeap []string

func (h StringHeap) Len() int { return len(h) }
func (h StringHeap) Less(i, j int) bool {

	if len(h[i]) == len(h[j]) {
		for idx := 0; idx < len(h[i]); idx++ {
			if h[i][idx] < h[j][idx] {
				return true
			} else if h[i][idx] > h[j][idx] {
				return false
			}

		}
	}
	return len(h[i]) < len(h[j])
}
func (h StringHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *StringHeap) Push(x interface{}) { *h = append(*h, x.(string)) }
func (h *StringHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kthLargestNumber(nums []string, k int) string {
	var h StringHeap

	for _, v := range nums {
		if len(h) != k {
			heap.Push(&h, v)
			continue
		}

		if len(v) >= len(h[0]) {
			heap.Push(&h, v)
			heap.Pop(&h)
		}
	}

	return h[0]
}

func kthLargestNumber(nums []string, k int) string {
	sort.Slice(nums, func(i, j int) bool {
		if len(nums[i]) == len(nums[j]) {
			for idx := 0; idx < len(nums[i]); idx++ {
				if nums[i][idx] > nums[j][idx] {
					return true
				}else if nums[i][idx] < nums[j][idx]{
					return false
				}

			}

		}

		return len(nums[i]) > len(nums[j])
	})
	return nums[k-1]
}
