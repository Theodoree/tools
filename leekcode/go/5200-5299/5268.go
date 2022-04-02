package _200_5299


/*
5268. 找出两数组的不同
给你两个下标从 0 开始的整数数组 nums1 和 nums2 ，请你返回一个长度为 2 的列表 answer ，其中：
answer[0] 是 nums1 中所有 不 存在于 nums2 中的 不同 整数组成的列表。
answer[1] 是 nums2 中所有 不 存在于 nums1 中的 不同 整数组成的列表。
*/
func findDifference(nums1 []int, nums2 []int) [][]int {
	var cnt [2002]int8
	for i := 0; i < len(nums1); i++ {
		cnt[nums1[i]+1000] |= 1
	}
	for i := 0; i < len(nums2); i++ {
		cnt[nums2[i]+1000] |= 1 << 1
	}

	var result [][]int
	result = append(result, []int{})
	result = append(result, []int{})

	for i := 0; i < len(cnt); i++ {
		if cnt[i]-1|2-cnt[i] >= 0 {
			result[cnt[i]-1] = append(result[cnt[i]-1], i-1000)
		}
	}
	return result
}

