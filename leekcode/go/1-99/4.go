package __99

import (
	"fmt"
	"sort"
)

/*
4. 寻找两个有序数组的中位数

给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5
*/

func findMedianSortedArrays(nums []int, nums2 []int) float64 {
	nums = append(nums, nums2...)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	lennumber := len(nums)
	if lennumber%2 == 1 {
		return float64(nums[lennumber/2])
	} else {
		return float64(nums[lennumber/2-1]+nums[lennumber/2]) / 2.0
	}
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

}
