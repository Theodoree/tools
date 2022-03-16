package _00_399



/*
349. 两个数组的交集

给定两个数组，编写一个函数来计算它们的交集。

示例 1:

输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2]
示例 2:

输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [9,4]
*/

func intersection(nums1 []int, nums2 []int) []int {

    var result []int
    Map := make(map[int]int)

    if len(nums1) > len(nums2) {
        nums1, nums2 = nums2, nums1
    }

    for _, v := range nums1 {
        Map[v]++
    }

    for _, v := range nums2 {
        if _, ok := Map[v]; ok {
            result = append(result, v)
            delete(Map, v)
        }
    }

    return result
}