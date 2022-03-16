package _000_2099

/*
2032. 至少在两个数组中出现的值
给你三个整数数组 nums1、nums2 和 nums3 ，请你构造并返回一个 与这三个数组都不同的 数组，且由 至少 在 两个 数组中出现的所有值组成。数组中的元素可以按 任意 顺序排列。
*/
func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	sort.Ints(nums3)
	var buf [101]int

	fn := func(num []int) {
		var last = -1
		for _, v := range num {
			if v == last {
				continue
			}
			buf[v]++
			last = v
		}
	}

	fn(nums1)
	fn(nums2)
	fn(nums3)

	result := make([]int, 0, len(nums1))
	for num, v := range buf {
		if v >= 2 {
			result = append(result, num)
		}
	}
	return result
}
