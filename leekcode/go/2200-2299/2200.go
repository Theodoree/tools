package _200_2299


/*
2200. 找出数组中的所有 K 近邻下标
给你一个下标从 0 开始的整数数组 nums 和两个整数 key 和 k 。K 近邻下标 是 nums 中的一个下标 i ，并满足至少存在一个下标 j 使得 |i - j| <= k 且 nums[j] == key 。
以列表形式返回按 递增顺序 排序的所有 K 近邻下标。
*/

func findKDistantIndices(nums []int, key int, k int) []int {
	var idx []int
	idx = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] == key {
			start := i - k
			end := i + k
			if start < 0 {
				start = 0
			}
			if end >= len(nums) {
				end = len(nums) - 1
			}
			for ; start <= end; start++ {
				idx[start] = 1
			}
		}
	}
	var result []int
	for k, val := range idx {
		if val == 1 {
			result = append(result, k)
		}
	}
	sort.Ints(result)

	return result
}

