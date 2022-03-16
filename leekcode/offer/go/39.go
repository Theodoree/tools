package _go



/*
面试题39. 数组中出现次数超过一半的数字

数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1:

输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
输出: 2

限制：

1 <= 数组长度 <= 50000
*/

func majorityElement(nums []int) int {

    var m = make(map[int]int)

    for _, v := range nums {
        m[v]++

        if m[v] > len(nums)/2 {
            return v
        }
    }

    return -1
}
