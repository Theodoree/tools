package _800_1899

import "sort"

/*
1874. 两个数组的最小乘积和
给定两个长度相等的数组a和b，它们的乘积和为数组中所有的a[i] * b[i]之和，其中0 <= i < a.length。

比如a = [1,2,3,4]，b = [5,2,3,1]时，它们的乘积和为1*5 + 2*2 + 3*3 + 4*1 = 22
现有两个长度都为n的数组nums1和nums2，你可以以任意顺序排序nums1，请返回它们的最小乘积和。


*/
func minProductSum(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	var sum int
	for i := 0; i < len(nums1); i++ {
		sum += nums1[i] * nums2[len(nums2)-i-1]
	}
	return sum
}
