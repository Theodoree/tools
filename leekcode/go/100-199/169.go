package _00_199

/*
169. 求众数
给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在众数。
*/
func majorityElement(nums []int) int {
    maps := make(map[int]int, len(nums))

    limit := len(nums) / 2

    for _,v := range nums {
        if _,ok := maps[v]; !ok {
            maps[v] = 1
        } else {
            maps[v]++
        }

        if maps[v] > limit {
            return v
        }
    }

    return -1
}