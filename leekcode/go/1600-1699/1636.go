package _600_1699

import "sort"

/*
1636. 按照频率将数组升序排序
给你一个整数数组 nums ，请你将数组按照每个值的频率 升序 排序。如果有多个值的频率相同，请你按照数值本身将它们 降序 排序。

请你返回排序后的数组。


-100 <= nums[i] <= 100
*/
func frequencySort(nums []int) []int {

	var counts [201]int
	for _, v := range nums {
		counts[v+100]++
	}

	sort.Slice(nums, func(i, j int) bool {
		total, total1 := counts[nums[i]+100], counts[nums[j]+100]

		if total != total1 {
			return total < total1
		}

		return nums[i] > nums[j]
	})

	return nums
}
