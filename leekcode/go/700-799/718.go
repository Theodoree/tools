package _00_799



/*
718. 最长重复子数组
给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。
*/
func findLength(nums1 []int, nums2 []int) int {
	m, n, res := len(nums1), len(nums2), 0
	for i := 0; i < m; i++ {
		length := min(n, m-i)
		res = max(res, maxLength(nums1, nums2, i, 0, length))
	}
	for i := 0; i < n; i++ {
		length := min(m, n-i)
		res = max(res, maxLength(nums1, nums2, 0, i, length))
	}
	return res
}

func maxLength(nums1, nums2 []int, start1, start2, length int) int {
	res, tempRes := 0, 0
	for i := 0; i < length; i++ {
		if nums1[start1+i] == nums2[start2+i] {
			tempRes++
		} else {
			tempRes = 0
		}
		res = max(res, tempRes)
	}
	return res
}
