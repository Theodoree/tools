package _00_399

import "math"

/*
350. 两个数组的交集 II
输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
我们可以不考虑输出结果的顺序。
进阶:

如果给定的数组已经排好序呢？你将如何优化你的算法？
如果 nums1 的大小比 nums2 小很多，哪种方法更优？
如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？

*/
func intersect(nums1 []int, nums2 []int) []int {

    filterMap := make(map[int]int)
    filter1Map := make(map[int]int)

    for _, v := range nums1 {
        filterMap[v]++
    }

    for _, v := range nums2 {
        filter1Map[v]++
    }

    var result []int
    var cnt int
    for k, v := range filterMap {
        if v1, ok := filter1Map[k]; ok {
            if v1 == v {
                cnt = v
            } else {
                cnt = int(math.Min(float64(v), float64(v1)))
            }

            for i := 0; i < cnt; i++ {
                result = append(result, k)
            }

        }

    }

    return result
}