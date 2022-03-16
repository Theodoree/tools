package _00_299


/*
219. 存在重复元素 II
给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的绝对值最大为 k。
*/

func containsNearbyDuplicate(nums []int, k int) bool {
    m := map[int]int{}
    for i, num := range nums {
        if j, ok := m[nums[i]]; ok && i-j <= k {
            return true
        }
        m[num] = i
    }
    return false
}
